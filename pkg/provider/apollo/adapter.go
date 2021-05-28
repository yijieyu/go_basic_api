package apollo

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"sync"
	"time"

	jsoniter "github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type ApolloClient struct {
	httpClient     *http.Client
	AppID          string
	Namespace      string
	ApolloServer   string
	Cluster        string
	NotificationId map[string]int
}

func NewApolloClient(appID, env string) *ApolloClient {

	apolloServer := defaultServer
	if s, ok := server[env]; ok {
		apolloServer = s
	}

	return &ApolloClient{
		httpClient: &http.Client{
			Timeout: 65 * time.Second,
			Transport: &http.Transport{
				DialContext: (&net.Dialer{
					Timeout:   65 * time.Second,
					KeepAlive: 65 * time.Second, // 保持长连接的时间
				}).DialContext, // 设置连接的参数
				ResponseHeaderTimeout: time.Minute * 65,
				IdleConnTimeout:       65 * time.Second, // 空闲连接的超时时间
				ExpectContinueTimeout: 65 * time.Second, // 等待服务第一个响应的超时时间
			},
		},
		AppID:          appID,
		ApolloServer:   apolloServer,
		Namespace:      "application",
		Cluster:        "default",
		NotificationId: map[string]int{},
	}
}

func (c *ApolloClient) Get(rp viper.RemoteProvider) (io.Reader, error) {

	b, err := c.GetNoCacheConfig()
	r := bytes.NewReader(b)
	return r, err
}
func (c *ApolloClient) Watch(rp viper.RemoteProvider) (io.Reader, error) {
	b, err := c.GetCacheConfig()
	r := bytes.NewReader(b)
	return r, err
}
func (c *ApolloClient) WatchChannel(rp viper.RemoteProvider) (<-chan *viper.RemoteResponse, chan bool) {

	quit := make(chan bool)
	viperResponseCh := make(chan *viper.RemoteResponse)
	go func(vc chan<- *viper.RemoteResponse, quit <-chan bool) {
		for {
			select {
			case <-quit:
				return
			default:
				// check 配置信息是否有变化
				isChange := c.V2Request()

				// 配置有变化，获取最新的配置
				if isChange {
					value, err := c.GetCacheConfig()
					vc <- &viper.RemoteResponse{Value: value, Error: err}
				}
			}
		}
	}(viperResponseCh, quit)
	return viperResponseCh, quit
}

func (c *ApolloClient) V2Request() (isChange bool) {
	reqData := []NotificationsReq{}

	for _, v := range strings.Split(c.Namespace, ",") {
		reqData = append(reqData, NotificationsReq{
			NamespaceName:  v,
			NotificationId: c.NotificationId[v],
		})
	}
	d, err := jsoniter.Marshal(reqData)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err":  err,
			"data": string(d),
		}).Warn("jsoniter.Marshal err")
		return false
	}

	v := url.Values{}
	v.Set("appId", c.AppID)
	v.Set("cluster", c.Cluster)
	v.Set("notifications", string(d))
	reqURL := fmt.Sprintf("%s/notifications/v2?", c.ApolloServer) + v.Encode()

	logrus.WithFields(logrus.Fields{
		"req_url": reqURL,
	}).Debug("notifications request")

	res, httpCode, err := c.getResp(reqURL)
	defer func() {
		if httpCode == http.StatusOK {
			isChange = true
		}
	}()

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"url":       reqURL,
			"http_code": httpCode,
			"err":       err,
		}).Debug("notifications getResp err")
		return
	}

	if res == nil {
		logrus.WithFields(logrus.Fields{
			"req_url":   reqURL,
			"http_code": httpCode,
			"err":       err,
		}).Debug("notifications getResp res is nil")
		return
	}

	logrus.WithFields(logrus.Fields{
		"req_url": reqURL,
		"res":     string(res),
	}).Debug("notifications resp")

	nResp := []NotificationsResp{}
	err = jsoniter.Unmarshal(res, &nResp)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"res": string(res),
			"err": err,
		}).Warn("jsoniter.Unmarshal err")
		return
	}

	for _, v := range nResp {
		c.NotificationId[v.NamespaceName] = v.NotificationId
	}

	return
}

func (c *ApolloClient) GetNoCacheConfig() ([]byte, error) {

	var wg sync.WaitGroup

	namespace := strings.Split(c.Namespace, ",")
	errs := map[int]error{}
	res := map[int]*AppConfigConfigData{}
	for i, v := range namespace {
		wg.Add(1)
		go func(n string, index int) {
			defer wg.Done()
			//reqURL = "http://test-apollo.weimiaocaishang.com/configs/wm-infoflow-api/default/application?ip=127.0.0.1&releaseKey=%s"
			reqURL := fmt.Sprintf("%s/configs/%s/%s/%s", c.ApolloServer, c.AppID, c.Cluster, n)

			result, err := c.getClusterConfig(reqURL)
			if err != nil {
				errs[index] = err
			} else {
				res[index] = result
			}
		}(v, i)
	}
	wg.Wait()

	if len(errs) > 0 {
		logrus.WithFields(logrus.Fields{
			"errs": errs,
		}).Error("cluster 获取失败")
		return nil, errors.New("cluster 获取失败")
	}

	if res[0] == nil {
		res[0] = &AppConfigConfigData{}
	}

	for i := 1; i < len(res); i++ {

		for k, v := range res[i].Configurations {
			res[0].Configurations[k] = v
		}
	}

	return c.getJson(*res[0])
}

func (c *ApolloClient) GetCacheConfig() ([]byte, error) {
	var wg sync.WaitGroup

	namespace := strings.Split(c.Namespace, ",")
	errs := map[int]error{}
	res := map[int]*AppConfigConfigData{}

	for i, v := range namespace {
		wg.Add(1)
		go func(n string, index int) {
			defer wg.Done()
			//reqURL := fmt.Sprintf("%s/configfiles/json/%s/%s/%s", c.ApolloServer, c.AppID, c.Cluster, n)
			reqURL := fmt.Sprintf("%s/configs/%s/%s/%s", c.ApolloServer, c.AppID, c.Cluster, n)
			result, err := c.getClusterConfig(reqURL)
			if err != nil {
				errs[index] = err
			} else {
				res[index] = result
			}
		}(v, i)
	}
	wg.Wait()

	if len(errs) > 0 {
		logrus.WithFields(logrus.Fields{
			"errs": errs,
		}).Error("cluster 获取失败")
		return nil, errors.New("cluster 获取失败")
	}

	if res[0] == nil {
		res[0] = &AppConfigConfigData{}
	}

	for i := 1; i < len(res); i++ {

		for k, v := range res[i].Configurations {
			res[0].Configurations[k] = v
		}
	}

	return c.getJson(*res[0])
}

func (c *ApolloClient) getClusterConfig(reqURL string) (*AppConfigConfigData, error) {

	logrus.WithFields(logrus.Fields{
		"req_url": reqURL,
	}).Debug("getConfig reqURL")

	res, httpCode, err := c.getResp(reqURL)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"req_url": reqURL,
			"err":     err,
		}).Warn("getConfig resp err")
		return nil, err
	}
	if httpCode != http.StatusOK {
		logrus.WithFields(logrus.Fields{
			"req_url": reqURL,
		}).Debug("getConfig resp http.Status != 200")
		return nil, errors.New("getConfig resp http.Status != 200")
	}

	if res == nil {
		logrus.WithFields(logrus.Fields{
			"req_url": reqURL,
		}).Warn("getConfig resp res is nil")
		return nil, errors.New("getConfig resp res is nil")
	}
	logrus.WithFields(logrus.Fields{
		"req_url":  reqURL,
		"resp_res": string(res),
	}).Debug("getConfig resp res")

	resData := &AppConfigConfigData{}
	err = jsoniter.Unmarshal(res, &resData)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"req_url": reqURL,
			"res":     string(res),
			"err":     err,
		}).Warn("getConfig resp res jsoniter.Unmarshal is err")
		return nil, err
	}

	return resData, nil
}

func (c *ApolloClient) getResp(reqURL string) ([]byte, int, error) {

	var resp *http.Response
	var err error

	for i := 0; i < 3; i++ {
		resp, err = c.httpClient.Get(reqURL)
		if err == nil {
			break
		}
	}

	if err != nil {
		return nil, 0, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotModified {
		return nil, resp.StatusCode, nil
	}

	if resp.StatusCode == http.StatusGatewayTimeout {
		return nil, resp.StatusCode, nil
	}

	if resp.StatusCode != http.StatusOK {
		return nil, resp.StatusCode, errors.New(resp.Status)
	}

	res, err := ioutil.ReadAll(resp.Body)
	return res, resp.StatusCode, err
}

func (c *ApolloClient) getClientIP() string {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		fmt.Println(err)
		return ""
	}

	for _, address := range addrs {

		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

func (c *ApolloClient) getJson(data AppConfigConfigData) ([]byte, error) {

	d := data.Configurations

	res := map[string]interface{}{}
	resSecond := map[string]map[string]interface{}{}
	thirdMap := map[string]map[string]map[string]map[string]interface{}{}
	arrMap := map[string][]interface{}{}

	reg := regexp.MustCompile(`\w+`)
	reg1 := regexp.MustCompile(`\d+`)

	for k, v := range d {
		arr := strings.Split(k, ".")
		if len(arr) == 1 {
			if strings.Contains(k, "[") {
				key := reg.FindString(k)
				if _, ok := arrMap[key]; !ok {
					arrMap[key] = []interface{}{}
				}
				arrMap[key] = append(arrMap[key], v)
			} else {
				res[k] = v
			}
		} else if len(arr) == 2 {
			if _, ok := resSecond[arr[0]]; !ok {
				resSecond[arr[0]] = map[string]interface{}{}
			}
			resSecond[arr[0]][arr[1]] = v
		} else if strings.Contains(k, "].") {
			if _, ok := thirdMap[arr[0]]; !ok {
				thirdMap[arr[0]] = map[string]map[string]map[string]interface{}{}
			}

			key := reg.FindString(arr[1])

			if _, ok := thirdMap[arr[0]][key]; !ok {
				thirdMap[arr[0]][key] = map[string]map[string]interface{}{}
			}
			index := reg1.FindString(arr[1])

			if _, ok := thirdMap[arr[0]][key][index]; !ok {
				thirdMap[arr[0]][key][index] = map[string]interface{}{}
			}

			thirdMap[arr[0]][key][index][arr[2]] = v

		}

	}
	for k, v := range resSecond {
		res[k] = v
	}
	for k, v := range arrMap {
		res[k] = v
	}

	thirdMap1 := map[string]map[string][]map[string]interface{}{}
	for k, v := range thirdMap {
		if _, ok := thirdMap1[k]; !ok {
			thirdMap1[k] = map[string][]map[string]interface{}{}
		}

		for kk, vv := range v {
			if _, ok := thirdMap1[k][kk]; !ok {
				thirdMap1[k][kk] = []map[string]interface{}{}
			}
			for _, vvv := range vv {
				thirdMap1[k][kk] = append(thirdMap1[k][kk], vvv)
			}

		}
	}

	for k, v := range thirdMap1 {
		res[k] = v
	}

	b, err := jsoniter.Marshal(res)

	return b, err
}
