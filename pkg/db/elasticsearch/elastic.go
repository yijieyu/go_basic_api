package elasticsearch

import (
	"log"
	"net"
	"net/http"
	"os"
	"time"

	elastic "github.com/olivere/elastic/v7"
	"github.com/yijieyu/go_basic_api/pkg/configuration"
)

type Client struct {
	*elastic.Client
}

func New(conf configuration.Elastic) (*Client, error) {
	opts := []elastic.ClientOptionFunc{
		elastic.SetHttpClient(&http.Client{
			Timeout: 30 * time.Second,
			Transport: &http.Transport{
				DialContext: (&net.Dialer{
					Timeout:   30 * time.Second,
					KeepAlive: 30 * time.Second,
				}).DialContext,
				ForceAttemptHTTP2:     true,
				MaxIdleConnsPerHost:   conf.MaxIdleConns,
				MaxConnsPerHost:       conf.MaxOpenConns,
				IdleConnTimeout:       90 * time.Second,
				ExpectContinueTimeout: 1 * time.Second,
				ResponseHeaderTimeout: 10 * time.Second,
			},
		}),
		elastic.SetBasicAuth(conf.User, conf.Password),
		elastic.SetURL(conf.URL),
		elastic.SetSniff(false),
	}

	if conf.Debug {
		opts = append(opts, elastic.SetTraceLog(log.New(os.Stdout, "[elastic] ", log.LstdFlags)))
	}

	c, err := elastic.NewClient(opts...)
	return &Client{Client: c}, err
}
