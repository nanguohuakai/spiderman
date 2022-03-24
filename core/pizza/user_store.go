package pizza

import (
	"errors"
	"github.com/nanguohuakai/spiderman/dto"
	"github.com/nanguohuakai/spiderman/pkg/httpclient"
)

//GetEmployeeInfo 获取个人信息
func (p *Client) GetEmployeeInfo(workcode string) (dto.EmployeeInfoRes, error) {
	u := p.PizzaConf.BaseUri + "/api/v1/employee/info/" + workcode
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

//GetAppointment 任命信息
func (p *Client) GetAppointment(workcode string) (dto.AppointmentItemRes, error) {
	u := p.PizzaConf.BaseUri + "/api/v1/appointment/" + workcode
	var output dto.AppointmentItemRes

	err := httpclient.Get(u, p.Conf, &output)
	if err != nil {
		return output, err
	}
	if output.Code != 0 {
		return output, errors.New(output.Msg)
	}
	return output, nil
}

//GetTalExperienceRecords 入司后履历  人才 PC端履历
func (p *Client) GetTalExperienceRecords(workcode string) (dto.ExperienceRecordRes, error) {
	u := p.PizzaConf.BaseUri + "/api/v1/changes/talent/experience-records/" + workcode
	var output dto.ExperienceRecordRes

	err := httpclient.Get(u, p.Conf, &output)
	if err != nil {
		return output, err
	}
	if output.Code != 0 {
		return output, errors.New(output.Msg)
	}
	return output, nil
}

//GetTalRecords 入司后履历
func (p *Client) GetTalRecords(input dto.PizzaInput) (dto.TalRecordsRes, error) {
	var params = dto.PaginationParams{
		Page:     input.Page,
		PageSize: input.PageSize,
	}
	path := "/api/v1/changes/talent/tal-records/" + input.Workcode
	u := httpclient.HttpBuildQuery(p.PizzaConf.BaseUri, path, params)

	var output dto.TalRecordsRes

	err := httpclient.Get(u, p.Conf, &output)
	if err != nil {
		return output, err
	}
	if output.Code != 0 {
		return output, errors.New(output.Msg)
	}
	return output, nil
}

//GetExps  入司前履历
func (p *Client) GetExps(input dto.PizzaInput) (dto.ExpItemRes, error) {
	var params = dto.PaginationParams{
		Page:     input.Page,
		PageSize: input.PageSize,
	}
	path := "/api/v1/exp/list/" + input.Workcode
	u := httpclient.HttpBuildQuery(p.PizzaConf.BaseUri, path, params)

	var output dto.ExpItemRes

	err := httpclient.Get(u, p.Conf, &output)
	if err != nil {
		return output, err
	}
	if output.Code != 0 {
		return output, errors.New(output.Msg)
	}
	return output, nil
}
