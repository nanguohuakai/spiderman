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

//GetDeptEmployeeCount 获取部门人数
func  (p *Client) GetDeptEmployeeCount(deptId string) (dto.DeptEmployeCountRes, error)  {
	path := "/api/v1/dept/employee/count/" + deptId
	var input interface{}
	u := httpclient.HttpBuildQuery(p.PizzaConf.BaseUri, path, input)

	var output dto.DeptEmployeCountRes

	err := httpclient.Get(u, p.Conf, &output)
	if err != nil {
		return output, err
	}
	if output.Code != 0 {
		return output, errors.New(output.Msg)
	}
	return output, nil
}