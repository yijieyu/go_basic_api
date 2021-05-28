package model

type AdLaunch struct {
	ID            int          `db:"id"`
	Name          string       `db:"name"`
	City          string       `db:"city"`
	StartTime     int64        `db:"start_time"`
	EndTime       int64        `db:"end_time"`
	Status        int          `db:"status"`
	Priority      int          `db:"priority"`
	InventoryType int          `db:"inventory_type"`
	Target        string       `db:"target"`
	Action        int          `db:"action"`
	PlaceID       int          `db:"place_id"`
	Material      []AdMaterial `db:"-"`
}

type AdPosition struct {
	ID     string `db:"id"`
	Name   string `db:"name"`
	Type   string `db:"type"`
	Status string `db:"status"`
}

type AdMaterial struct {
	ID       int    `db:"id"`
	AID      int    `db:"aid"`
	Material string `db:"material"`
	Type     int    `db:"type"`
}

// API 接口 协议
type Request struct {
	Device    *RequestDevice `json:"device"`
	App       *RequestApp    `json:"app"`
	Ad        *RequestAd     `json:"ad"`
	User      *RequestUser   `json:"user"`
	RequestID string         `json:"request_id"`
}

type RequestUser struct {
	ID int `json:"id"`
}

type RequestDevice struct {
	Os        int         `json:"os"` // 操作系统0：安卓，1：IOS
	Osv       string      `json:"osv"`
	Model     string      `json:"model"`
	UA        string      `json:"ua"`
	Make      string      `json:"make"`
	Brand     string      `json:"brand"`
	Mac       string      `json:"mac"`
	Imei      string      `json:"imei"`
	Imsi      string      `json:"imsi"`
	AndroidID string      `json:"android_id"`
	Idfa      string      `json:"idfa"`
	Idfv      string      `json:"idfv"`
	OpenudID  string      `json:"openudid"`
	Geo       *RequestGeo `json:"geo"`
	IP        string      `json:"ip"`
	Province  string      `json:"province"`
	City      string      `json:"city"`
}

type RequestGeo struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type RequestApp struct {
	Ver     string `json:"ver"`
	Bundle  string `json:"bundle"`
	Channel string `json:"channel"`
}

type RequestAd struct {
	Type    int `json:"type"`
	PlaceID int `json:"place_id"`
}

type ResponseAd struct {
	Config *ResponseConfig `json:"config"`
	Ads    []ResponseData  `json:"ads"`
}

type ResponseConfig struct {
}

type ResponseData struct {
	Aid           int      `json:"aid"`
	InventoryType int      `json:"inventory_type"`
	Amd           string   `json:"amd"`
	W             int      `json:"w"`
	H             int      `json:"h"`
	Action        int      `json:"action"`
	Target        string   `json:"target"`
	ImpTrackers   []string `json:"imp_trackers"`
	ClickTrackers []string `json:"click_trackers"`
}

type ReportParam struct {
	RequestID string `form:"request_id"`
	IP        string `form:"ip"`
	Mac       string `form:"mac"`
	Aid       int    `form:"aid"`
	PlaceID   int    `form:"place_id"`
	Os        int    `form:"os"`
	Osv       string `form:"osv"`
	AndroidID string `form:"android_id"`
	Channel   string `form:"channel"`
	AppVer    string `form:"app_ver"`
	AppBundle string `form:"app_bundle"`
	Idfa      string `form:"idfa"`
	Idfv      string `form:"idfv"`
	OpenudID  string `form:"openudid"`
	Uid       int    `form:"uid"`
}

type ThirdAdSwitchResp struct {
	Banner     bool `json:"banner"`
	OpenScreen bool `json:"open_screen"`
}
