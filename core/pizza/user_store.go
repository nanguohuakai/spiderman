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

//GetEmployeeList 获取人员列表  （不支持分页）
func (p *Client) GetEmployeeList(input dto.EmployeeListInput) (dto.EmployeeListRes, error) {
	u := p.PizzaConf.BaseUri + "/api/v1/employee/list"
	var output dto.EmployeeListRes
	err := httpclient.Post(u, p.Conf, input, &output)
	if err != nil {
		return output, err
	}
	if output.Code != 0 {
		return output, errors.New(output.Msg)
	}
	return output, nil
}

//GetEmployeeListWithPage 获取人员列表  （支持分页）
func (p *Client) GetEmployeeListWithPage(input dto.EmployeeListInputWithPage) (dto.EmployeeListWithPageRes, error) {
	u := p.PizzaConf.BaseUri + "/api/v1/employee/list"
	var output dto.EmployeeListWithPageRes
	err := httpclient.Post(u, p.Conf, input, &output)
	if err != nil {
		return output, err
	}
	if output.Code != 0 {
		return output, errors.New(output.Msg)
	}
	return output, nil
}

//GetEmployeeListV2 获取人员列表  （不支持分页）
func (p *Client) GetEmployeeListV2(input dto.EmployeeListV2Input) (dto.EmployeeListRes, error) {
	u := p.PizzaConf.BaseUri + "/api/v2/employee/list"
	var output dto.EmployeeListRes
	err := httpclient.Post(u, p.Conf, input, &output)
	if err != nil {
		return output, err
	}
	if output.Code != 0 {
		return output, errors.New(output.Msg)
	}
	return output, nil
}

//GetEmployeeListWithPageV2 获取人员列表  （支持分页）
func (p *Client) GetEmployeeListWithPageV2(input dto.EmployeeListV2InputWithPage) (dto.EmployeeListWithPageRes, error) {
	u := p.PizzaConf.BaseUri + "/api/v2/employee/list"
	var output dto.EmployeeListWithPageRes
	err := httpclient.Post(u, p.Conf, input, &output)
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

//GetChangePromotionList 显示在人才手机端的 晋升足迹
func (p *Client) GetChangePromotionList(input dto.PizzaInput) (dto.PromotionTrackRecordRes, error) {
	var params = dto.PaginationParams{
		Page:     input.Page,
		PageSize: input.PageSize,
	}
	path := "/api/v1/changes/promotion/tracks/" + input.Workcode
	u := httpclient.HttpBuildQuery(p.PizzaConf.BaseUri, path, params)

	var output dto.PromotionTrackRecordRes

	err := httpclient.Get(u, p.Conf, &output)
	if err != nil {
		return output, err
	}
	if output.Code != 0 {
		return output, errors.New(output.Msg)
	}
	return output, nil
}

//GetEduList 员工的教育列表
func (p *Client) GetEduList(input dto.PizzaInput) (dto.EduItemRes, error) {
	var params = dto.PaginationParams{
		Page:     input.Page,
		PageSize: input.PageSize,
	}
	path := "/api/v1/edu/list/" + input.Workcode
	u := httpclient.HttpBuildQuery(p.PizzaConf.BaseUri, path, params)

	var output dto.EduItemRes

	err := httpclient.Get(u, p.Conf, &output)
	if err != nil {
		return output, err
	}
	if output.Code != 0 {
		return output, errors.New(output.Msg)
	}
	return output, nil
}

//GetKpiList 员工的绩效列表
func (p *Client) GetKpiList(workcode string) (dto.KpiListRes, error) {
	//var params interface{}
	u := p.PizzaConf.BaseUri + "/api/v1/kpi/list/" + workcode
	//u := httpclient.HttpBuildQuery(p.PizzaConf.BaseUri, path, params)
	var output dto.KpiListRes

	err := httpclient.Get(u, p.Conf, &output)
	if err != nil {
		return output, err
	}
	if output.Code != 0 {
		return output, errors.New(output.Msg)
	}
	return output, nil
}

//GetRewardList 获取员工奖惩信息
func (p *Client) GetRewardList(workcode string, input dto.RewardInput) (dto.RewardListRes, error) {
	path := "/api/v1/rewards/list/" + workcode
	u := httpclient.HttpBuildQuery(p.PizzaConf.BaseUri, path, input)

	var output dto.RewardListRes

	err := httpclient.Get(u, p.Conf, &output)
	if err != nil {
		return output, err
	}
	if output.Code != 0 {
		return output, errors.New(output.Msg)
	}
	return output, nil
}