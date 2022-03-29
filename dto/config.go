package dto

import "strings"

type AppConf struct {
	ServiceName string `json:"service_name"`
	Token       string `json:"token"`
}

type SsoConf struct {
	BaseUri string `json:"base_uri"` //base url
	AppId   string `json:"appid"`
	AppKey  string `json:"appkey"`
}

type PizzaConf struct {
	BaseUri string `json:"base_uri"`
}

type AlertConf struct {
	BaseUri string `json:"base_uri"` //base url
	Level   string `json:"level"`    //日志级别限制 >= level
	Env     string `json:"env"`      //项目环境
}

type ScheduleConf struct {
	BaseUri      string `json:"base_uri"`     //base url talent-scheduler.chengjiukehu.com
	CallbackUri  string `json:"callback_uri"` //callback url
	callbackPath string `json:"callback_path"`
	isRegister   bool   `json:"is_register"` //是否注册
}

func (s *ScheduleConf) SetCallbackPath(path string) {
	s.callbackPath =  "/" + strings.Trim(path, "/")
}

func (s *ScheduleConf) SetIsRegister(isRegister bool) {
	s.isRegister = isRegister
}

func (s ScheduleConf) GetIsRegister() bool {
	return s.isRegister
}

func (s ScheduleConf) GetCallbackPath() string {
	return s.callbackPath
}