package sso

import (
	"github.com/nanguohuakai/spiderman/dto"
)

type SsoInterface interface {
	Verify(token string) ( dto.SsoVerifyResponse, error)
	JsConfig(url string) ( dto.SsoResponse, error)
	Ticket() ( dto.SsoResponse, error)
}

type Client struct {
	Conf dto.AppConf
}

func NewSsoClient(conf dto.AppConf) SsoInterface {
	var pc SsoInterface
	pc = &Client{
		Conf: conf,
	}
	return pc
}
