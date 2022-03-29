package pizza

import (
	"github.com/nanguohuakai/spiderman/dto"
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
	GetDeptInfo(code string) (string, error)                                                              //GetDeptInfo 获取部门信息
	GetProjectList(viewerWorkcode string, viewedWorkcode string) (dto.ProjectListRes, error)              //GetProjectList 获取项目列表

	SyncEsChangeData() (dto.PizzaResponse, error)   //SyncEsChangeData 同步工作数据到es
	SyncEsExpData() (dto.PizzaResponse, error)      //SyncEsExpData 同步 experience_list 表数据到 es
	SyncEsEduData() (dto.PizzaResponse, error)      //SyncEsEduData 同步教育信息到es
	SyncEsEmployeeData() (dto.PizzaResponse, error) //SyncEsEmployeeData 同步用户信息到es
	SyncEsRewardsData() (dto.PizzaResponse, error)  //SyncEsRewardsData 同步用户奖惩信息到es
	SyncEsKpiData() (dto.PizzaResponse, error)      //SyncEsKpiData 同步用户奖惩信息到es
}

type Client struct {
	Conf      dto.AppConf
	PizzaConf dto.PizzaConf
}

func NewPizzaClient(conf dto.AppConf, pizzaConf dto.PizzaConf) PizzaInterface {
	var pc PizzaInterface
	pc = &Client{
		Conf:      conf,
		PizzaConf: pizzaConf,
	}
	return pc
}
