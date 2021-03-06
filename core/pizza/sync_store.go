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

//SyncEsDeptData 同步部门信息到es
func (p *Client) SyncEsDeptData() (dto.PizzaResponse, error) {
	return syncResponse(p, getSyncUrl(p, "/api/v1/dept/job/es"))
}

//SyncEsPmAbilityData 同步pm-ability信息到es
func (p *Client) SyncEsPmAbilityData() (dto.PizzaResponse, error) {
	return syncResponse(p, getSyncUrl(p, "/api/v1/project/pm-ability/job/es"))
}

//SyncEsCulturalData 同步文化评分信息到es
func (p *Client) SyncEsCulturalData() (dto.PizzaResponse, error) {
	return syncResponse(p, getSyncUrl(p, "/api/v1/cultural/job/es"))
}

//SyncEsFamilyData 同步家庭信息到es
func (p *Client) SyncEsFamilyData() (dto.PizzaResponse, error) {
	return syncResponse(p, getSyncUrl(p, "/api/v1/family/job/es"))
}

//SyncEsDimissionData 同步部门离职历史信息到es
func (p *Client) SyncEsDimissionData() (dto.PizzaResponse, error) {
	return syncResponse(p, getSyncUrl(p, "/api/v1/dept/dimission/job/es"))
}

//SyncEsPromotionsData 同步部门晋升历史信息到es
func (p *Client) SyncEsPromotionsData() (dto.PizzaResponse, error) {
	return syncResponse(p, getSyncUrl(p, "/api/v1/dept/promotions/job/es"))
}

//SyncEsLeadershipsData 同步部门领导历史信息到es
func (p *Client) SyncEsLeadershipsData() (dto.PizzaResponse, error) {
	return syncResponse(p, getSyncUrl(p, "/api/v1/dept/leaderships/job/es"))
}