package pizza

import (
	"errors"
	"github.com/nanguohuakai/spiderman/dto"
	"github.com/nanguohuakai/spiderman/pkg/httpclient"
)

//GetHrPsGr 个人行为数据
func (p *Client) GetHrPsGr(input dto.HrpsInput) (dto.HrpsGrRes, error) {
	path := "/api/v1/behavior/personal/data"
	u := httpclient.HttpBuildQuery(p.PizzaConf.BaseUri, path, input)

	var output dto.HrpsGrRes

	err := httpclient.Get(u, p.Conf, &output)
	if err != nil {
		return output, err
	}
	if output.Code != 0 {
		return output, errors.New(output.Msg)
	}
	return output, nil
}

//GetHrPsGrList 个人行为月度数据
func (p *Client) GetHrPsGrList(input dto.HrpsInput) (dto.HrpsGrListRes, error) {
	path := "/api/v1/behavior/personal/month/list"
	u := httpclient.HttpBuildQuery(p.PizzaConf.BaseUri, path, input)

	var output dto.HrpsGrListRes

	err := httpclient.Get(u, p.Conf, &output)
	if err != nil {
		return output, err
	}
	if output.Code != 0 {
		return output, errors.New(output.Msg)
	}
	return output, nil
}

//GetHrPs 组织健康、组织氛围
func (p *Client) GetHrPs(input dto.HrpsInput) (dto.HrpsRes, error) {
	path := "/api/v1/health/org/data"
	u := httpclient.HttpBuildQuery(p.PizzaConf.BaseUri, path, input)

	var output dto.HrpsRes

	err := httpclient.Get(u, p.Conf, &output)
	if err != nil {
		return output, err
	}
	if output.Code != 0 {
		return output, errors.New(output.Msg)
	}
	return output, nil
}


//GetHrPsList 组织健康、组织氛围 月度数据
func (p *Client) GetHrPsList(input dto.HrpsInput) (dto.HrpsListRes, error) {
	path := "/api/v1/health/org/month/list"
	u := httpclient.HttpBuildQuery(p.PizzaConf.BaseUri, path, input)

	var output dto.HrpsListRes

	err := httpclient.Get(u, p.Conf, &output)
	if err != nil {
		return output, err
	}
	if output.Code != 0 {
		return output, errors.New(output.Msg)
	}
	return output, nil
}