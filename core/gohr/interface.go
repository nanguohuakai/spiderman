package gohr

import (
	"github.com/nanguohuakai/spiderman/dto"
	"strings"
)

type GoHrInterface interface {
	GetTalentOkrUserList(appid string, targetWorkcode string) (dto.OkrUserListRes, error)       //GetTalentOkrUserList 获取人才 okr 列表
	GetRecordJournal(input dto.RecordJournalUserInput) (dto.RecordJournalYachRes, error)        //GetRecordJournal 一线 6 小时
	GetHonorUserInfo(input dto.HonorUserInfoInput) (dto.HonorUserInfoRes, error)                //GetHonorUserInfo 荣耀：获取用户详情
	GetYachWeekReport(input dto.WeekReportUserInput) (dto.WeekReportYachRes, error)             //GetYachWeekReport 获取 yach 周报
	GetReportListFrom360(workcode string, page int, pageSize int) (dto.Report360DataRes, error) //GetReportListFrom360 获取汇报线 360

	GetJobChannel(input dto.JobChannelInput) (dto.JobChannelDataRes, error) //GetJobChannel 职务通道
	GetJobc(input dto.JobcInput) (dto.JobcDataRes, error)                   //GetJobc 获取职位信息
	GetJobSeq(input dto.JobSeqInput) (dto.JobSeqDataRes, error)             //GetJobSeq 职务序列
	GetJobSub(input dto.JobSubInput) (dto.JobSubDataRes, error)             //GetJobSub 获取子序列
	GetQuitList(input dto.QuitOrTalInput) (dto.QuitDataRes, error)          //GetQuitList 是否被动离职
	GetFullToTal(input dto.QuitOrTalInput) (dto.TalDataRes, error)          //GetFullToTal 全转专

	ScheduleCreate(input dto.ScheduleCreateOrUpdateInput) (dto.OpsScheduleGoHrRes, error) //ScheduleCreate 创建日程
	ScheduleUpdate(input dto.ScheduleCreateOrUpdateInput) (dto.OpsScheduleGoHrRes, error) //ScheduleUpdate 更新日程
	ScheduleCancel(input dto.ScheduleCancelInput) (dto.OpsScheduleGoHrRes, error)         //ScheduleCancel 取消日程
	ScheduleInvite(input dto.ScheduleInviteInput) (dto.OpsScheduleGoHrRes, error)         //ScheduleInvite 加入日程
	ScheduleList24(input dto.ScheduleLisInput) (dto.OpsScheduleLisRes, error)             //ScheduleList24 忙闲

	GetSecretDeptByUser(input dto.SecretDeptInput) (dto.SecretRes, error) //GetSecretDeptByUser 获取部门
	RecentSecret(input dto.SecretInput) (dto.SecretRes, error)
	BeforeSecret(input dto.SecretInput) (dto.SecretRes, error)
	DecodeSecret(input []map[string]string) (dto.SecretRes, error) //解密
}

type Client struct {
	AppConf  dto.AppConf
	GoHrConf dto.GoHrConf
}

func NewGoHrClient(conf dto.AppConf, GoHrConf dto.GoHrConf) GoHrInterface {
	//特殊去掉域名最后元素"/"
	GoHrConf.BaseUri = strings.TrimRight(GoHrConf.BaseUri, "/")
	var uc GoHrInterface
	uc = &Client{
		AppConf:  conf,
		GoHrConf: GoHrConf,
	}
	return uc
}
