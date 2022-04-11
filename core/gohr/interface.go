package gohr

import (
	"github.com/nanguohuakai/spiderman/dto"
	"strings"
)

type GoHrInterface interface {
	GetTalentOkrUserList(appid string, targetWorkcode string) (dto.OkrUserListRes, error)
	GetRecordJournal(input dto.RecordJournalUserInput) (dto.RecordJournalYachRes, error)
	GetHonorUserInfo(input dto.HonorUserInfoInput) (dto.HonorUserInfoRes, error)
	GetYachWeekReport(input dto.WeekReportUserInput) (dto.WeekReportYachRes, error)
	GetReportListFrom360(workcode string, page int, pageSize int) (dto.Report360DataRes, error)

	GetJobChannel(input dto.JobChannelInput) (dto.JobChannelDataRes, error)
	GetJobc(input dto.JobcInput) (dto.JobcDataRes, error)
	GetJobSub(input dto.JobSubInput) (dto.JobSubDataRes, error)
	GetQuitList(input dto.QuitOrTalInput) (dto.QuitDataRes, error)

	ScheduleCreate(input dto.ScheduleCreateOrUpdateInput) (dto.OpsScheduleGoHrRes, error)
	ScheduleUpdate(input dto.ScheduleCreateOrUpdateInput) (dto.OpsScheduleGoHrRes, error)
	ScheduleCancel(input dto.ScheduleCancelInput) (dto.OpsScheduleGoHrRes, error)
	ScheduleInvite(input dto.ScheduleInviteInput) (dto.OpsScheduleGoHrRes, error)
	ScheduleList24(input dto.ScheduleLisInput) (dto.OpsScheduleLisRes, error)
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
