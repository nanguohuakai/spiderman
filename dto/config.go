package dto

type AppConf struct {
	ServiceName string `json:"service_name"`
	Token       string `json:"token"`
}

type SsoConf struct {
	BaseUri     string `json:"base_uri"`
	AppId       string `json:"appid"`
	AppKey      string `json:"appkey"`
}

type PizzaConf struct {
	BaseUri     string `json:"base_uri"`
}

type AlertConf struct {
	BaseUri     string `json:"base_uri"`
}
