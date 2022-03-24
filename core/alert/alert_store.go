package alert

import (
	"fmt"
	"github.com/nanguohuakai/spiderman/pkg/httpclient"
	"go.uber.org/zap/zapcore"
)

type alertParams struct {
	ServiceName string    `json:"service_name"`
	Body        alertBody `json:"body"`
}

type alertBody struct {
	Level   string `json:"level"`
	Env     string `json:"env"`
	Content string `json:"content"`
}

type alertRes struct {
	Code int    `json:"code"`
	Msg  string `json:"msg,omitempty"`
}

//getLogLevel 根据日志级别text 获取日志级别 zapcore.Level
//zapcore.Level  is a logging priority. Higher levels are more important.
func getLogLevel(level string) zapcore.Level {
	var lv zapcore.Level
	if err := lv.UnmarshalText([]byte(level)); err != nil {
		panic("get log lv error: " + fmt.Sprintf("%#v", err))
	}
	return lv
}

//do post sendAlert api
func (a *Client) sendAlert(msg string, level string) (alertRes, error) {
	var output alertRes
	//日志等级 当前输出日志等级level >= 限制等级级别 可发送
	if !getLogLevel(a.AlertConf.Level).Enabled(getLogLevel("warn")) {
		return output, nil
	}

	url := a.AlertConf.BaseUri + "/sendAlert"
	var params = alertParams{
		ServiceName: a.Conf.ServiceName,
		Body: alertBody{
			Level:   level,
			Env:     a.AlertConf.Env,
			Content: msg,
		},
	}

	err := httpclient.Post(url, a.Conf, params, &output)
	return output, err
}
//Warn -
//return bool true 发送成功； false 发送失败 ，error 失败原因
func (a *Client) Warn(msg string) (bool, error) {
	_, err := a.sendAlert(msg, "warn")
	if err != nil {
		return false, err
	}
	return true, nil
}

//Error -
//return bool true 发送成功； false 发送失败 ，error 失败原因
func (a *Client) Error(msg string) (bool, error) {
	_, err := a.sendAlert(msg, "error")
	if err != nil {
		return false, err
	}
	return true, nil
}

//Info -
//return bool true 发送成功； false 发送失败 ，error 失败原因
func (a *Client) Info(msg string) (bool, error) {
	_, err := a.sendAlert(msg, "info")
	if err != nil {
		return false, err
	}
	return true, nil
}

//Debug -
//return bool true 发送成功； false 发送失败 ，error 失败原因
func (a *Client) Debug(msg string) (bool, error) {
	_, err := a.sendAlert(msg, "debug")
	if err != nil {
		return false, err
	}
	return true, nil
}
