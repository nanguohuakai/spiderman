package sso

import (
	"errors"
	"github.com/nanguohuakai/spiderman/dto"
	"github.com/nanguohuakai/spiderman/pkg/httpclient"
)

//Verify -
func (s *Client) Verify(token string) (dto.SsoVerifyResponse, error) {
	u := s.SsoConf.BaseUri + "/api/v2/sso/verify?token=" + token + withQuery(s)
	var output dto.SsoVerifyResponse

	e := checkSsoConf(s.SsoConf)
	if e != nil {
		return output, e
	}

	err := httpclient.Get(u, s.Conf, &output)
	if err != nil {
		return output, err
	}
	if output.Code != 0 {
		return output, errors.New(output.Msg)
	}

	return output, nil
}

//JsConfig -
func (s *Client) JsConfig(url string) (dto.JsConfigResponse, error) {
	u := s.SsoConf.BaseUri + "/api/v1/yach/js_config?url=" + url + withQuery(s)
	var output dto.JsConfigResponse

	e := checkSsoConf(s.SsoConf)
	if e != nil {
		return output, e
	}

	err := httpclient.Get(u, s.Conf, &output)
	if err != nil {
		return output, err
	}

	if output.Code != 0 {
		return output, errors.New(output.Msg)
	}
	return output, nil
}

//Ticket -
func (s *Client) Ticket() (dto.TicketResponse, error) {
	u := s.SsoConf.BaseUri + "/api/v1/sso/basic/get_ticket" + withQuery(s)
	var output dto.TicketResponse
	//check sso conf appid & appkey
	e := checkSsoConf(s.SsoConf)
	if e != nil {
		return output, e
	}

	err := httpclient.Get(u, s.Conf, &output)
	if err != nil {
		return output, err
	}

	if output.Code != 0 {
		return output, errors.New(output.Msg)
	}
	return output, nil
}

//with appid & appkey
func withQuery(s *Client) string {
	return "&appid=" + s.SsoConf.AppId + "&appkey=" + s.SsoConf.AppKey
}

//check sso conf
func checkSsoConf(conf dto.SsoConf) error {
	var err error
	if conf.AppId == "" || conf.AppKey == "" {
		err = errors.New("ssoConf empty")
	}
	return err
}
