package alert

import "github.com/nanguohuakai/spiderman/dto"

//AlertInterface -
type AlertInterface interface {
	Warn(msg string) (bool, error)
	Error(msg string) (bool, error)
	Info(msg string) (bool, error)
	Debug(msg string) (bool, error)
	K8sSend(msg string) (bool, error)
	//On(serviceName string) *Client todo 暂时不开放此方法
	sendAlert(msg string, level string) (dto.AlertRes, error)
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

// On 提供切换service name方法
//func (a *Client) On(serviceName string) *Client {
//	a.Conf.ServiceName = serviceName
//	return a
//}
