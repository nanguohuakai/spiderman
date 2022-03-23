package sso

import (
	"github.com/nanguohuakai/spiderman/dto"
)

type SsoInterface interface {
	Verify(token string) (dto.SsoVerifyResponse, error)
	JsConfig(url string) (dto.JsConfigResponse, error)
	Ticket() (dto.TicketResponse, error)
}

type Client struct {
	Conf    dto.AppConf
	SsoConf dto.SsoConf
}

func NewSsoClient(conf dto.AppConf, ssoConf dto.SsoConf) SsoInterface {
	var pc SsoInterface
	pc = &Client{
		Conf: conf,
		SsoConf: ssoConf,
	}
	return pc
}
