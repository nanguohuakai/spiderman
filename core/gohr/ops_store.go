package gohr

import (
	"errors"
	"github.com/nanguohuakai/spiderman/dto"
	"github.com/nanguohuakai/spiderman/pkg/httpclient"
)

//GetJobChannel 职务通道
func (c *Client) GetJobChannel(input dto.JobChannelInput) (dto.JobChannelDataRes, error) {
	u := httpclient.HttpBuildQuery(c.GoHrConf.BaseUri, "/api/v1/ops/job/channel", input)
	var output dto.JobChannelDataRes
	err := httpclient.Get(u, c.AppConf, &output)
	if err != nil {
		return output, err
	}
	if output.Code != 0 {
		return output, errors.New(output.Msg)
	}
	return output, nil
}

//GetJobc 获取职位信息
func (c *Client) GetJobc(input dto.JobcInput) (dto.JobcDataRes, error) {
	u := httpclient.HttpBuildQuery(c.GoHrConf.BaseUri, "/api/v1/ops/job/jobc", input)
	var output dto.JobcDataRes
	err := httpclient.Get(u, c.AppConf, &output)
	if err != nil {
		return output, err
	}
	if output.Code != 0 {
		return output, errors.New(output.Msg)
	}
	return output, nil
}

//GetJobSeq 职务序列
func (c *Client) GetJobSeq(input dto.JobSeqInput) (dto.JobSeqDataRes, error) {
	u := httpclient.HttpBuildQuery(c.GoHrConf.BaseUri, "/api/v1/ops/job/seq", input)
	var output dto.JobSeqDataRes
	err := httpclient.Get(u, c.AppConf, &output)
	if err != nil {
		return output, err
	}
	if output.Code != 0 {
		return output, errors.New(output.Msg)
	}
	return output, nil
}

//GetJobSub 获取子序列
func (c *Client) GetJobSub(input dto.JobSubInput) (dto.JobSubDataRes, error) {
	u := httpclient.HttpBuildQuery(c.GoHrConf.BaseUri, "/api/v1/ops/job/sub", input)
	var output dto.JobSubDataRes
	err := httpclient.Get(u, c.AppConf, &output)
	if err != nil {
		return output, err
	}
	if output.Code != 0 {
		return output, errors.New(output.Msg)
	}
	return output, nil
}

//GetQuitList 是否被动离职
func (p *Client) GetQuitList(input dto.QuitOrTalInput) (dto.QuitDataRes, error) {
	u := p.GoHrConf.BaseUri + "/api/v1/ops/quit"
	var output dto.QuitDataRes
	err := httpclient.Post(u, p.AppConf, input, &output)
	if err != nil {
		return output, err
	}
	if output.Code != 0 {
		return output, errors.New(output.Msg)
	}
	return output, nil
}

//GetFullToTal 全转专
func (p *Client) GetFullToTal(input dto.QuitOrTalInput) (dto.TalDataRes, error) {
	u := p.GoHrConf.BaseUri + "/api/v1/ops/tal"
	var output dto.TalDataRes
	err := httpclient.Post(u, p.AppConf, input, &output)
	if err != nil {
		return output, err
	}
	if output.Code != 0 {
		return output, errors.New(output.Msg)
	}
	return output, nil
}

//ScheduleCreate 创建日程
func (p *Client) ScheduleCreate(input dto.ScheduleCreateOrUpdateInput) (dto.OpsScheduleGoHrRes, error) {
	u := p.GoHrConf.BaseUri + "/api/v1/ops/schedule/create"
	var output dto.OpsScheduleGoHrRes
	err := httpclient.Post(u, p.AppConf, input, &output)
	if err != nil {
		return output, err
	}
	if output.Code != 0 {
		return output, errors.New(output.Msg)
	}
	return output, nil
}

//ScheduleUpdate 更新日程
func (p *Client) ScheduleUpdate(input dto.ScheduleCreateOrUpdateInput) (dto.OpsScheduleGoHrRes, error) {
	u := p.GoHrConf.BaseUri + "/api/v1/ops/schedule/update"
	var output dto.OpsScheduleGoHrRes
	err := httpclient.Post(u, p.AppConf, input, &output)
	if err != nil {
		return output, err
	}
	if output.Code != 0 {
		return output, errors.New(output.Msg)
	}
	return output, nil
}

//ScheduleCancel 取消日程
func (p *Client) ScheduleCancel(input dto.ScheduleCancelInput) (dto.OpsScheduleGoHrRes, error) {
	u := p.GoHrConf.BaseUri + "/api/v1/ops/schedule/cancel"
	var output dto.OpsScheduleGoHrRes
	err := httpclient.Post(u, p.AppConf, input, &output)
	if err != nil {
		return output, err
	}
	if output.Code != 0 {
		return output, errors.New(output.Msg)
	}
	return output, nil
}