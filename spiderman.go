// Package spiderman 为集成 Pizza, Sso, Schedule, Alert, Msg, Gohr, Comment等服务的客户端
package spiderman

import (
	"errors"
	"github.com/nanguohuakai/spiderman/core/alert"
	"github.com/nanguohuakai/spiderman/core/gohr"
	"github.com/nanguohuakai/spiderman/core/pizza"
	"github.com/nanguohuakai/spiderman/core/schedule"
	"github.com/nanguohuakai/spiderman/core/sso"
	"github.com/nanguohuakai/spiderman/dto"
)

//Spiderman support pipelining using the Pizza, Sso, Schedule and Alert methods
type Spiderman interface {
	Alert(conf dto.AlertConf) (alert.AlertInterface, error)             //Alert 消息服务
	Pizza(conf dto.PizzaConf) (pizza.PizzaInterface, error)             //Pizza pizza服务
	Schedule(conf dto.ScheduleConf) (schedule.ScheduleInterface, error) //Schedule 定时任务调度服务
	Sso(conf dto.SsoConf) (sso.SsoInterface, error)                     //Sso sso服务
	GoHr(conf dto.GoHrConf) (gohr.GoHrInterface, error)                 //GoHr gohr服务
}

type SpidermanClient struct {
	AppConf dto.AppConf
}

//The NewSpiderman interface is the primary interface for working with Spiderman
func NewSpiderman(conf dto.AppConf) (Spiderman, error) {
	//verify app conf
	if conf.Token == "" || conf.ServiceName == "" {
		return nil, errors.New("AppConf 配置信息缺失")
	}

	var uc Spiderman
	uc = &SpidermanClient{
		AppConf: conf,
	}
	return uc, nil

}

//Pizza store pizza service
func (receiver *SpidermanClient) Pizza(conf dto.PizzaConf) (pizza.PizzaInterface, error) {
	//verify pizza conf
	if conf.BaseUri == "" {
		return nil, errors.New("PizzaConf 配置信息缺失")
	}
	return pizza.NewPizzaClient(receiver.AppConf, conf), nil
}

//Sso store sso  service
func (receiver *SpidermanClient) Sso(conf dto.SsoConf) (sso.SsoInterface, error) {
	//verify sso conf
	if conf.BaseUri == "" || conf.AppId == "" || conf.AppKey == "" {
		return nil, errors.New("SsoConf 配置信息缺失")
	}
	return sso.NewSsoClient(receiver.AppConf, conf), nil
}

//Alert store alert service
func (receiver *SpidermanClient) Alert(conf dto.AlertConf) (alert.AlertInterface, error) {
	//verify pizza conf
	if conf.BaseUri == "" || conf.Level == "" || conf.Env == "" {
		return nil, errors.New("AlertConf 配置信息缺失")
	}

	var levelMap = []string{"warn", "debug", "info", "error"}
	f := false
	for _, v := range levelMap {
		if v == conf.Level {
			f = true
			break
		}
	}
	if !f {
		return nil, errors.New("AlertConf Level 信息错误")
	}
	return alert.NewAlertClient(receiver.AppConf, conf), nil
}

//Schedule store schedule service
func (receiver *SpidermanClient) Schedule(conf dto.ScheduleConf) (schedule.ScheduleInterface, error) {
	if conf.BaseUri == "" || conf.CallbackUri == "" {
		return nil, errors.New("ScheduleConf 配置信息缺失")
	}
	return schedule.NewClient(receiver.AppConf, conf), nil
}

//GoHr store GoHr service
func (receiver *SpidermanClient) GoHr(conf dto.GoHrConf) (gohr.GoHrInterface, error) {
	if conf.BaseUri == "" {
		return nil, errors.New("ScheduleConf 配置信息缺失")
	}
	return gohr.NewGoHrClient(receiver.AppConf, conf), nil
}
