package sso

import (
	"errors"
	"spiderman/dto"
	"spiderman/pkg/httpclient"
)

func (s *Client) Verify(token string) ( dto.SsoResponse, error)  {
	u := "/api/v2/sso/verify?token=" + token + "&appid=" + s.Conf.AppId + "&appkey=" + s.Conf.AppKey
	var output dto.SsoResponse

	err := httpclient.Get(u, s.Conf, &output)

	if err != nil {
		return output, err
	}

	if output.Code != 0 {
		return output, errors.New(output.Msg)
	}

	return output, nil
}

func (s *Client) JsConfig(url string) (dto.SsoResponse, error)  {
	u := "/api/v1/yach/js_config?url=" + url + "&appid=" + s.Conf.AppId + "&appkey=" + s.Conf.AppKey
	var output dto.SsoResponse

	err := httpclient.Get(u, s.Conf, &output)

	if err != nil {
		return output, err
	}

	if output.Code != 0 {
		return output, errors.New(output.Msg)
	}

	return output, nil
}

func (s *Client) Ticket() (dto.SsoResponse, error)  {
	u := "/api/v1/sso/basic/get_ticket"
	var output dto.SsoResponse

	err := httpclient.Get(u, s.Conf, &output)

	if err != nil {
		return output, err
	}

	if output.Code != 0 {
		return output, errors.New(output.Msg)
	}

	return output, nil
}