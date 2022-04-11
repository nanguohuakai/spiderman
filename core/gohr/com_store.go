package gohr

import (
	"errors"
	"github.com/nanguohuakai/spiderman/dto"
	"github.com/nanguohuakai/spiderman/pkg/httpclient"
)

//GetTalentOkrUserList 获取人才 okr 列表
func (c *Client) GetTalentOkrUserList(appid string, targetWorkcode string) (dto.OkrUserListRes, error) {
	var input = dto.OkrUserListInput{
		AppId:          appid,
		TargetWorkcode: targetWorkcode,
	}
	u := httpclient.HttpBuildQuery(c.GoHrConf.BaseUri, "/api/v1/okr/talent/user/list", input)
	var output dto.OkrUserListRes
	err := httpclient.Get(u, c.AppConf, &output)
	if err != nil {
		return output, err
	}
	if output.Code != 0 {
		return output, errors.New(output.Msg)
	}
	return output, nil
}

//GetRecordJournal 一线 6 小时
func (c *Client) GetRecordJournal(input dto.RecordJournalUserInput) (dto.RecordJournalYachRes, error) {
	u := httpclient.HttpBuildQuery(c.GoHrConf.BaseUri, "/api/v1/record/journal", input)
	var output dto.RecordJournalYachRes
	err := httpclient.Get(u, c.AppConf, &output)
	if err != nil {
		return output, err
	}
	if output.Code != 0 {
		return output, errors.New(output.Msg)
	}
	return output, nil
}

//GetHonorUserInfo 荣耀：获取用户详情
func (c *Client) GetHonorUserInfo(input dto.HonorUserInfoInput) (dto.HonorUserInfoRes, error) {
	u := httpclient.HttpBuildQuery(c.GoHrConf.BaseUri, "/api/v1/honor/user-info", input)
	var output dto.HonorUserInfoRes
	err := httpclient.Get(u, c.AppConf, &output)
	if err != nil {
		return output, err
	}
	if output.Code != 0 {
		return output, errors.New(output.Msg)
	}
	return output, nil
}

//GetYachWeekReport 获取 yach 周报
func (c *Client) GetYachWeekReport(input dto.WeekReportUserInput) (dto.WeekReportYachRes, error) {
	u := httpclient.HttpBuildQuery(c.GoHrConf.BaseUri, "/api/v1/yach/week-report", input)
	var output dto.WeekReportYachRes
	err := httpclient.Get(u, c.AppConf, &output)
	if err != nil {
		return output, err
	}
	if output.Code != 0 {
		return output, errors.New(output.Msg)
	}
	return output, nil
}

//GetReportListFrom360 获取汇报线 360
func (c *Client) GetReportListFrom360(workcode string, page int, pageSize int) (dto.Report360DataRes, error) {
	var input = dto.GoHrInput{
		Workcode: workcode,
		Page:     page,
		PageSize: pageSize,
	}
	u := httpclient.HttpBuildQuery(c.GoHrConf.BaseUri, "/api/v1/360/report", input)
	var output dto.Report360DataRes
	err := httpclient.Get(u, c.AppConf, &output)
	if err != nil {
		return output, err
	}
	if output.Code != 0 {
		return output, errors.New(output.Msg)
	}
	return output, nil
}
