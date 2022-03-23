package alert

import "github.com/nanguohuakai/spiderman/dto"

type AlertInterface interface {
	//Send()
}

type Client struct {
	Conf      dto.AppConf
	AlertConf dto.AlertConf
}

func NewAlertClient(conf dto.AppConf, alertConf dto.AlertConf) AlertInterface {
	var pc AlertInterface
	pc = &Client{
		Conf:      conf,
		AlertConf: alertConf,
	}
	return pc
}
