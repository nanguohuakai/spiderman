package dto

import "strings"

type AppConf struct {
	ServiceName string `json:"service_name"` //服务名称
	Token       string `json:"token"`
}

type SsoConf struct {
	BaseUri string `json:"base_uri"` //base url sso服务域名
	AppId   string `json:"appid"`    //appid
	AppKey  string `json:"appkey"`   //appkey
}

type PizzaConf struct {
	BaseUri string `json:"base_uri"` //pizza 服务域名
}

type GoHrConf struct {
	BaseUri string `json:"base_uri"` //gohr 服务域名
}

type AlertConf struct {
	BaseUri string `json:"base_uri"` //base url alert 服务域名
	Level   string `json:"level"`    //日志级别限制 >= level
	Env     string `json:"env"`      //项目环境
}

type ScheduleConf struct {
	BaseUri      string `json:"base_uri"`      //base url schedule服务域名
	CallbackUri  string `json:"callback_uri"`  //callback url 回调域名
	callbackPath string `json:"callback_path"` //回调路径
	isRegister   bool   `json:"is_register"`   //是否注册
}

//SetCallbackPath 设置回调路径
func (s *ScheduleConf) SetCallbackPath(path string) {
	//去掉路径path 前面"/",统一添加"/"
	s.callbackPath = "/" + strings.Trim(path, "/")
}

//GetCallbackPath 获取回调路径
func (s ScheduleConf) GetCallbackPath() string {
	return s.callbackPath
}

//SetIsRegister 设置是否注册
func (s *ScheduleConf) SetIsRegister(isRegister bool) {
	s.isRegister = isRegister
}

//GetIsRegister 获取是否注册
func (s ScheduleConf) GetIsRegister() bool {
	return s.isRegister
}
