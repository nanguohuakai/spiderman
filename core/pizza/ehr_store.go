package pizza

import (
	"errors"
	"github.com/nanguohuakai/spiderman/dto"
	"github.com/nanguohuakai/spiderman/pkg/httpclient"
)

//GetDimissionList 获取部门离职历史信息
func (p *Client) GetDimissionList(input dto.PaginationParams) (dto.DimissionListRes, error) {
	u := httpclient.HttpBuildQuery(p.PizzaConf.BaseUri, "/api/v1/dept/dimission/list", input)

	var output dto.DimissionListRes

	err := httpclient.Get(u, p.Conf, &output)
	if err != nil {
		return output, err
	}
	if output.Code != 0 {
		return output, errors.New(output.Msg)
	}
	return output, nil
}

//GetPromotionsList 获取部门晋升历史信息
func (p *Client) GetPromotionsList(input dto.PaginationParams) (dto.PromotionListRes, error) {
	u := httpclient.HttpBuildQuery(p.PizzaConf.BaseUri, "/api/v1/dept/dimission/list", input)

	var output dto.PromotionListRes

	err := httpclient.Get(u, p.Conf, &output)
	if err != nil {
		return output, err
	}
	if output.Code != 0 {
		return output, errors.New(output.Msg)
	}
	return output, nil
}

//GetLeaderList 获取部门领导历史信息
func (p *Client) GetLeaderList(input dto.PaginationParams) (dto.LeadershipsListRes, error) {
	u := httpclient.HttpBuildQuery(p.PizzaConf.BaseUri, "/api/v1/dept/dimission/list", input)

	var output dto.LeadershipsListRes

	err := httpclient.Get(u, p.Conf, &output)
	if err != nil {
		return output, err
	}
	if output.Code != 0 {
		return output, errors.New(output.Msg)
	}
	return output, nil
}

//GetEhrChanges 获取履历列表 (不支持分页)
func (p *Client) GetEhrChanges(input dto.EmployeeListInput) (dto.EhrChangesRes, error) {
	u := p.PizzaConf.BaseUri + "/api/v1/changes/list"
	var output dto.EhrChangesRes
	err := httpclient.Post(u, p.Conf, input, &output)
	if err != nil {
		return output, err
	}
	if output.Code != 0 {
		return output, errors.New(output.Msg)
	}
	return output, nil
}

//GetEhrChangesWithPage 获取履历列表  (支持分页)
func (p *Client) GetEhrChangesWithPage(input dto.EmployeeListInputWithPage) (dto.EhrChangeWithPageRes, error) {
	u := p.PizzaConf.BaseUri + "/api/v1/changes/list"
	var output dto.EhrChangeWithPageRes
	err := httpclient.Post(u, p.Conf, input, &output)
	if err != nil {
		return output, err
	}
	if output.Code != 0 {
		return output, errors.New(output.Msg)
	}
	return output, nil
}