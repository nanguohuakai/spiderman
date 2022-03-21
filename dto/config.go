package dto

type AppConf struct {
	ServiceName string `json:"service_name"`
	Token       string `json:"token"`
	BaseUri     string `json:"base_uri"`
	Timeout     int    `json:"timeout"`
	AppId       string `json:"appid,omitempty"`
	AppKey      string `json:"appkey,omitempty"`
}
