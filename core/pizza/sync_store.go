package pizza

import (
	"errors"
	"github.com/nanguohuakai/spiderman/dto"
	"github.com/nanguohuakai/spiderman/pkg/httpclient"
)

//syncResponse 获取同步任务返回结果
func syncResponse(p *Client, u string) (dto.PizzaResponse, error) {
	var output dto.PizzaResponse
	err := httpclient.Get(u, p.Conf, &output)
	if err != nil {
		return output, err
	}
	if output.Code != 0 {
		return output, errors.New(output.Msg)
	}
	return output, nil
}

//getSyncUrl 获取同步任务url
func getSyncUrl(p *Client, path string) string {
	return p.PizzaConf.BaseUri + path
}

//SyncEsChangeData 同步工作数据到es
func (p *Client) SyncEsChangeData() (dto.PizzaResponse, error) {
	return syncResponse(p, getSyncUrl(p, "/api/v2/changes/job/es"))
}

//SyncEsExpData 同步 experience_list 表数据到 es
func (p *Client) SyncEsExpData() (dto.PizzaResponse, error) {
	return syncResponse(p, getSyncUrl(p, "/api/v1/exp/job/es"))
}

//SyncEsEduData 同步教育信息到es
func (p *Client) SyncEsEduData() (dto.PizzaResponse, error) {
	return syncResponse(p, getSyncUrl(p, "/api/v1/edu/job/es"))
}

//SyncEsEmployeeData 同步用户信息到es
func (p *Client) SyncEsEmployeeData() (dto.PizzaResponse, error) {
	return syncResponse(p, getSyncUrl(p, "/api/v1/employee/job/es"))
}

//SyncEsRewardsData 同步用户奖惩信息到es
func (p *Client) SyncEsRewardsData() (dto.PizzaResponse, error) {
	return syncResponse(p, getSyncUrl(p, "/api/v1/rewards/job/es"))
}

//SyncEsKpiData 同步kpi信息到es
func (p *Client) SyncEsKpiData() (dto.PizzaResponse, error) {
	return syncResponse(p, getSyncUrl(p, "/api/v1/kpi/job/es"))
}