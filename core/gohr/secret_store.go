package gohr

import (
	"errors"
	"github.com/nanguohuakai/spiderman/dto"
	"github.com/nanguohuakai/spiderman/pkg/httpclient"
)

//GetSecretDeptByUser 获取部门
func (c *Client) GetSecretDeptByUser(input dto.SecretDeptInput) (dto.SecretRes, error) {
	u := httpclient.HttpBuildQuery(c.GoHrConf.BaseUri, "/api/v1/secret/depts", input)
	var output dto.SecretRes
	err := httpclient.Get(u, c.AppConf, &output)
	if err != nil {
		return output, err
	}
	if output.Code != 0 {
		return output, errors.New(output.Msg)
	}
	return output, nil
}

//RecentSecret -
func (c *Client) RecentSecret(input dto.SecretInput) (dto.SecretRes, error) {
	u := httpclient.HttpBuildQuery(c.GoHrConf.BaseUri, "/api/v1/secret/recent", input)
	var output dto.SecretRes
	err := httpclient.Get(u, c.AppConf, &output)
	if err != nil {
		return output, err
	}
	if output.Code != 0 {
		return output, errors.New(output.Msg)
	}
	return output, nil
}

//BeforeSecret -
func (c *Client) BeforeSecret(input dto.SecretInput) (dto.SecretRes, error) {
	u := httpclient.HttpBuildQuery(c.GoHrConf.BaseUri, "/api/v1/secret/before", input)
	var output dto.SecretRes
	err := httpclient.Get(u, c.AppConf, &output)
	if err != nil {
		return output, err
	}
	if output.Code != 0 {
		return output, errors.New(output.Msg)
	}
	return output, nil
}

//DecodeSecret 解密
func (p *Client) DecodeSecret(input []map[string]string) (dto.SecretRes, error) {
	u := p.GoHrConf.BaseUri + "/api/v1/secret/decode"
	var output dto.SecretRes
	err := httpclient.Post(u, p.AppConf, input, &output)
	if err != nil {
		return output, err
	}
	if output.Code != 0 {
		return output, errors.New(output.Msg)
	}
	return output, nil
}
