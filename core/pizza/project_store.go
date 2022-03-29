package pizza

import (
	"errors"
	"github.com/nanguohuakai/spiderman/dto"
	"github.com/nanguohuakai/spiderman/pkg/httpclient"
)

type projectListInput struct {
	ViewerWorkcode string `json:"viewer_workcode" url:"viewer_workcode"`
	ViewedWorkcode string `json:"viewed_workcode" url:"viewed_workcode"`
}

//GetProjectList 获取项目列表
func (p *Client) GetProjectList(viewerWorkcode string, viewedWorkcode string) (dto.ProjectListRes, error) {
	var params = projectListInput{
		ViewerWorkcode: viewerWorkcode,
		ViewedWorkcode: viewedWorkcode,
	}
	u := httpclient.HttpBuildQuery(p.PizzaConf.BaseUri, "/api/v1/project/query-list", params)
	var output dto.ProjectListRes

	err := httpclient.Get(u, p.Conf, &output)
	if err != nil {
		return output, err
	}
	if output.Code != 0 {
		return output, errors.New(output.Msg)
	}

	return output, nil
}
