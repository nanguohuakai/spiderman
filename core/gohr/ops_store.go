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