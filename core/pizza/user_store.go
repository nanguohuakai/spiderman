package pizza

import (
	"errors"
	"spiderman/dto"
	"spiderman/pkg/httpclient"
)

func (p *Client) GetEmployeeInfo(code string) ( dto.PizzaResponse, error) {
	u := "/api/v1/employee/info/" + code
	var output dto.PizzaResponse

	err := httpclient.Get(u, p.Conf, &output)

	if err != nil {
		return output, err
	}

	if output.Code != 0 {
		return output, errors.New(output.Msg)
	}

	return output, nil
}