package pizza

import (
	"github.com/nanguohuakai/spiderman/dto"
)

//PizzaInterface pizza
type PizzaInterface interface {
	GetEmployeeInfo(code string) (dto.EmployeeInfoRes, error)
	GetDeptInfo(code string) (string, error)
}

type Client struct {
	Conf      dto.AppConf
	PizzaConf dto.PizzaConf
}

func NewPizzaClient(conf dto.AppConf, pizzaConf dto.PizzaConf) PizzaInterface {
	var pc PizzaInterface
	pc = &Client{
		Conf:      conf,
		PizzaConf: pizzaConf,
	}
	return pc
}
