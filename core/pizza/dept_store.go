package pizza

import (
	"errors"
	"github.com/nanguohuakai/spiderman/dto"
	"github.com/nanguohuakai/spiderman/pkg/httpclient"
)

//GetDeptList 获取部门信息 (不支持分页）
func (p *Client) GetDeptList(input dto.DeptListInput) (dto.DeptListRes, error) {
	u := p.PizzaConf.BaseUri + "/api/v1/dept/list"
	var output dto.DeptListRes
	err := httpclient.Post(u, p.Conf, input, &output)
	if err != nil {
		return output, err
	}
	if output.Code != 0 {
		return output, errors.New(output.Msg)
	}
	return output, nil
}

//GetDeptListWithPage 获取部门信息 (支持分页）
func (p *Client) GetDeptListWithPage(input dto.DeptListInputWithPage) (dto.DeptListWithPageRes, error) {
	u := p.PizzaConf.BaseUri + "/api/v1/dept/list"
	var output dto.DeptListWithPageRes
	err := httpclient.Post(u, p.Conf, input, &output)
	if err != nil {
		return output, err
	}
	if output.Code != 0 {
		return output, errors.New(output.Msg)
	}
	return output, nil
}