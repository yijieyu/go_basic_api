package apollo

var defaultServer = "http://apollo.weimiaocaishang.com"
var server = map[string]string{
	"debug":   "http://test-apollo.weimiaocaishang.com",
	"testing": "http://test-apollo.weimiaocaishang.com",
	"stage":   "http://stage-apollo.weimiaocaishang.com",
	"prod":    "http://apollo.weimiaocaishang.com",
}

type AppConfigConfigData struct {
	AppId          string                 `json:"appId"`
	Cluster        string                 `json:"cluster"`
	NamespaceName  string                 `json:"namespaceName"`
	ReleaseKey     string                 `json:"releaseKey"`
	Configurations map[string]interface{} `json:"configurations"`
	//Configurations json.RawMessage `json:"configurations"`
}

type AppConfig struct {
	AppID          string              `json:"app_id"`
	Namespace      string              `json:"namespace"`
	ApolloServer   string              `json:"apollo_server"`
	EnvPath        string              `json:"env_path"`
	EnvFormat      string              `json:"env_format"`
	Cluster        string              `json:"cluster"`
	NotificationId map[string]int      `json:"-"`
	ConfigData     AppConfigConfigData `json:"-"`
}

type NotificationsReq struct {
	NamespaceName  string `json:"namespaceName"`
	NotificationId int    `json:"notificationId"`
}

type NotificationsResp struct {
	NamespaceName  string `json:"namespaceName"`
	NotificationId int    `json:"notificationId"`
}
