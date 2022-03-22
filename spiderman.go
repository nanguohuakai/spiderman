package spiderman

import (
	"github.com/nanguohuakai/spiderman/core/pizza"
	"github.com/nanguohuakai/spiderman/core/sso"
	"github.com/nanguohuakai/spiderman/dto"
)

type SpidermanInterface interface {
	Pizza() pizza.PizzaInterface
	Sso()   sso.SsoInterface
}

type SpidermanClient struct {
	PizzaCore pizza.PizzaInterface
	SsoCore   sso.SsoInterface
	//GoHr  gohr.GoHr
}

func NewSpiderman(conf dto.AppConf) SpidermanInterface {
	var uc SpidermanInterface
	uc = &SpidermanClient{
		PizzaCore: pizza.NewPizzaClient(conf),
		SsoCore:   sso.NewSsoClient(conf),
	}
	return uc
}

func (receiver SpidermanClient) Pizza() pizza.PizzaInterface {
	return receiver.PizzaCore
}

func (receiver SpidermanClient) Sso() sso.SsoInterface {
	return receiver.SsoCore
}
