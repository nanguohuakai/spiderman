package sso

import (
	"git.100tal.com/jituan_xinxi_be/talent-spiderman-go/dto"
)

type SsoInterface interface {
	Verify(token string) ( dto.SsoResponse, error)
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
