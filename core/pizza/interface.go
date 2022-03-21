package pizza

import (
	"spiderman/dto"
)

type PizzaInterface interface {
	GetEmployeeInfo(code string) (dto.PizzaResponse, error)
	GetDeptInfo(code string) (string, error)
}

type Client struct {
	Conf dto.AppConf
}

func NewPizzaClient(conf dto.AppConf) PizzaInterface {
	var pc PizzaInterface
	pc = &Client{
		Conf: conf,
	}
	return pc
}
