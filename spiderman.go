package spiderman

import (
	"spiderman/core/pizza"
	"spiderman/core/sso"
	"spiderman/dto"
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
