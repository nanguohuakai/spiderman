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

//GetLanYunProjectUser 获取蓝云项目
func (p *Client) GetLanYunProjectUser(input dto.UserProjectInput) (dto.ProjectUsersRes, error) {
	u := httpclient.HttpBuildQuery(p.PizzaConf.BaseUri, "/api/v1/project/user/general", input)
	var output dto.ProjectUsersRes

	err := httpclient.Get(u, p.Conf, &output)
	if err != nil {
		return output, err
	}
	if output.Code != 0 {
		return output, errors.New(output.Msg)
	}

	return output, nil
}

func (p *Client) GetLanYunProjectList(input dto.ProjectListInput) (dto.ProjectListWithPageRes, error) {
	u := httpclient.HttpBuildQuery(p.PizzaConf.BaseUri, "/api/v1/project/user/project/list", input)
	var output dto.ProjectListWithPageRes

	err := httpclient.Get(u, p.Conf, &output)
	if err != nil {
		return output, err
	}
	if output.Code != 0 {
		return output, errors.New(output.Msg)
	}

	return output, nil
}


//GetProjectAward 获取奖项列表（人才同步使用，其他系统暂不要使用）
func (p *Client) GetProjectAward(page int, pageSize int) (dto.ProjectAwardRes, error) {
	var params = dto.PaginationParams{
		Page:     page,
		PageSize: pageSize,
	}
	path := "/api/v1/project/awards-list"
	u := httpclient.HttpBuildQuery(p.PizzaConf.BaseUri, path, params)

	var output dto.ProjectAwardRes

	err := httpclient.Get(u, p.Conf, &output)
	if err != nil {
		return output, err
	}
	if output.Code != 0 {
		return output, errors.New(output.Msg)
	}
	return output, nil
}