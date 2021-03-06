package pizza

import (
	"github.com/nanguohuakai/spiderman/dto"
	"strings"
)

//PizzaInterface pizza
type PizzaInterface interface {
	GetEmployeeInfo(workcode string) (dto.EmployeeInfoRes, error)                                         //GetEmployeeInfo 获取个人信息
	GetEmployeeList(input dto.EmployeeListInput) (dto.EmployeeListRes, error)                             //GetEmployeeList 获取人员列表  （不支持分页）
	GetEmployeeListWithPage(input dto.EmployeeListInputWithPage) (dto.EmployeeListWithPageRes, error)     //GetEmployeeListWithPage  获取人员列表  （支持分页）
	GetEmployeeListV2(input dto.EmployeeListV2Input) (dto.EmployeeListRes, error)                         //GetEmployeeListV2 获取人员列表  （不支持分页）
	GetEmployeeListWithPageV2(input dto.EmployeeListV2InputWithPage) (dto.EmployeeListWithPageRes, error) //GetEmployeeListWithPageV2  获取人员列表  （支持分页）
	GetAppointment(workcode string) (dto.AppointmentItemRes, error)                                       //GetAppointment 任命信息
	GetTalExperienceRecords(workcode string) (dto.ExperienceRecordRes, error)                             //GetTalExperienceRecords 入司后履历  人才 PC端履历
	GetTalRecords(input dto.PizzaInput) (dto.TalRecordsRes, error)                                        //GetTalRecords 入司后履历
	GetExps(input dto.PizzaInput) (dto.ExpItemRes, error)                                                 //GetExps  入司前履历
	GetChangePromotionList(input dto.PizzaInput) (dto.PromotionTrackRecordRes, error)                     //GetChangePromotionList 显示在人才手机端的 晋升足迹
	GetEduList(input dto.PizzaInput) (dto.EduItemRes, error)                                              //GetEduList 员工的教育列表
	GetKpiList(workcode string) (dto.KpiListRes, error)                                                   //GetKpiList 员工的绩效列表
	GetRewardList(workcode string, input dto.RewardInput) (dto.RewardListRes, error)                      //GetRewardList 获取员工奖惩信息
	GetCulturalList(input dto.CulturalListInput) (dto.CulturalListRes, error)                             //GetCulturalList 获取用户文化评分
	GetFamilyList(workcode string) (dto.FamilyRes, error)                                                 //GetFamilyList 获取家庭信息
	GetDottedLineList(input dto.PizzaWorkcodesInput) (dto.DottedLineWithPageRes, error)                   //GetDottedLineList 获取虚线上级信息

	GetDeptList(input dto.DeptListInput) (dto.DeptListRes, error)                         //GetDeptList 获取部门信息 (不支持分页）
	GetDeptListWithPage(input dto.DeptListInputWithPage) (dto.DeptListWithPageRes, error) //GetDeptListWithPage 获取部门信息 (支持分页）
	GetDeptEmployeeCount(deptId string) (dto.DeptEmployeCountRes, error)                  //GetDeptEmployeeCount 获取部门人数

	GetProjectList(viewerWorkcode string, viewedWorkcode string) (dto.ProjectListRes, error) //GetProjectList 获取项目列表
	GetLanYunProjectUser(input dto.UserProjectInput) (dto.ProjectUsersRes, error)            //GetLanYunProjectUser 获取蓝云项目
	GetLanYunProjectList(input dto.ProjectListInput) (dto.ProjectListWithPageRes, error)     //GetLanYunProjectList 获取蓝云项目列表
	GetProjectAward(page int, pageSize int) (dto.ProjectAwardRes, error)                     //GetProjectAward 获取奖项列表（人才同步使用，其他系统暂不要使用）

	GetBuildList(input dto.BuildListInput) (dto.BuildListRes, error) //GetBuildList 获取楼宇信息

	GetHrPsGr(input dto.HrpsInput) (dto.HrpsGrRes, error)         //GetHrPsGr 个人行为数据
	GetHrPsGrList(input dto.HrpsInput) (dto.HrpsGrListRes, error) //GetHrPsGrList 个人行为月度数据
	GetHrPs(input dto.HrpsInput) (dto.HrpsRes, error)             //GetHrPs 组织健康、组织氛围
	GetHrPsList(input dto.HrpsInput) (dto.HrpsListRes, error)     //GetHrPsList 组织健康、组织氛围 月度数据

	GetDimissionList(input dto.PaginationParams) (dto.DimissionListRes, error)                   //GetDimissionList 获取部门离职历史信息
	GetPromotionsList(input dto.PaginationParams) (dto.PromotionListRes, error)                  //GetPromotionsList 获取部门晋升历史信息
	GetLeaderList(input dto.PaginationParams) (dto.LeadershipsListRes, error)                    //GetLeaderList 获取部门领导历史信息
	GetEhrChanges(input dto.EmployeeListInput) (dto.EhrChangesRes, error)                        //GetEhrChanges 获取履历列表 (不支持分页)
	GetEhrChangesWithPage(input dto.EmployeeListInputWithPage) (dto.EhrChangeWithPageRes, error) //GetEhrChangesWithPage 获取履历列表  (支持分页)

	SyncEsChangeData() (dto.PizzaResponse, error)      //SyncEsChangeData 同步工作数据到es
	SyncEsExpData() (dto.PizzaResponse, error)         //SyncEsExpData 同步 experience_list 表数据到 es
	SyncEsEduData() (dto.PizzaResponse, error)         //SyncEsEduData 同步教育信息到es
	SyncEsEmployeeData() (dto.PizzaResponse, error)    //SyncEsEmployeeData 同步用户信息到es
	SyncEsRewardsData() (dto.PizzaResponse, error)     //SyncEsRewardsData 同步用户奖惩信息到es
	SyncEsKpiData() (dto.PizzaResponse, error)         //SyncEsKpiData 同步用户奖惩信息到es
	SyncEsDeptData() (dto.PizzaResponse, error)        //SyncEsDeptData 同步部门信息到es
	SyncEsPmAbilityData() (dto.PizzaResponse, error)   //SyncEsPmAbilityData 同步pm-ability信息到es
	SyncEsCulturalData() (dto.PizzaResponse, error)    //SyncEsCulturalData 同步文化评分信息到es
	SyncEsFamilyData() (dto.PizzaResponse, error)      //SyncEsFamilyData 同步家庭信息到es
	SyncEsDimissionData() (dto.PizzaResponse, error)   //SyncEsDimissionData 同步部门离职历史信息到es
	SyncEsPromotionsData() (dto.PizzaResponse, error)  //SyncEsPromotionsData 同步部门晋升历史信息到es
	SyncEsLeadershipsData() (dto.PizzaResponse, error) //SyncEsLeadershipsData 同步部门领导历史信息到es

}

type Client struct {
	Conf      dto.AppConf
	PizzaConf dto.PizzaConf
}

func NewPizzaClient(conf dto.AppConf, pizzaConf dto.PizzaConf) PizzaInterface {
	var pc PizzaInterface
	pizzaConf.BaseUri = strings.TrimRight(pizzaConf.BaseUri, "/")
	pc = &Client{
		Conf:      conf,
		PizzaConf: pizzaConf,
	}
	return pc
}
