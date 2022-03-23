package pizza

import (
	"errors"
	"github.com/nanguohuakai/spiderman/dto"
	"github.com/nanguohuakai/spiderman/pkg/httpclient"
)

func (p *Client) GetEmployeeInfo(code string) ( dto.EmployeeInfoRes, error) {
	u := p.PizzaConf.BaseUri + "/api/v1/employee/info/" + code
	var output dto.EmployeeInfoRes

	err := httpclient.Get(u, p.Conf, &output)

	if err != nil {
		return output, err
	}
	if output.Code != 0 {
		return output, errors.New(output.Msg)
	}

	return output, nil
}
