package pizza

import (
	"errors"
	"github.com/nanguohuakai/spiderman/dto"
	"github.com/nanguohuakai/spiderman/pkg/httpclient"
)

//GetBuildList 获取楼宇信息
func (p *Client) GetBuildList(input dto.BuildListInput) (dto.BuildListRes, error) {
	u := p.PizzaConf.BaseUri + "/api/v1/build/list"
	var output dto.BuildListRes
	err := httpclient.Post(u, p.Conf, input, &output)
	if err != nil {
		return output, err
	}
	if output.Code != 0 {
		return output, errors.New(output.Msg)
	}
	return output, nil
}
