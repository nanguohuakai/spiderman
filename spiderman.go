//spiderman client
package spiderman

import (
	"errors"
	"github.com/nanguohuakai/spiderman/core/alert"
	"github.com/nanguohuakai/spiderman/core/pizza"
	"github.com/nanguohuakai/spiderman/core/sso"
	"github.com/nanguohuakai/spiderman/dto"
)

type Spiderman interface {
	Pizza(conf dto.PizzaConf) (pizza.PizzaInterface, error)
	Sso(conf dto.SsoConf) (sso.SsoInterface, error)
	Alert(conf dto.AlertConf) (alert.AlertInterface, error)
}

type SpidermanClient struct {
	AppConf dto.AppConf
}

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

//Pizza store
func (receiver *SpidermanClient) Pizza(conf dto.PizzaConf) (pizza.PizzaInterface, error) {
	//verify pizza conf
	if conf.BaseUri == "" {
		return nil, errors.New("PizzaConf 配置信息缺失")
	}
	return pizza.NewPizzaClient(receiver.AppConf, conf), nil
}

//Sso store
func (receiver *SpidermanClient) Sso(conf dto.SsoConf) (sso.SsoInterface, error) {
	//verify sso conf
	if conf.BaseUri == "" || conf.AppId == "" || conf.AppKey == "" {
		return nil, errors.New("SsoConf 配置信息缺失")
	}
	return sso.NewSsoClient(receiver.AppConf, conf), nil
}

//Alert store
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
