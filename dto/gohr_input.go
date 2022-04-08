package dto

type OkrUserListInput struct {
	AppId          string `json:"appid" url:"appid" `
	TargetWorkcode string `json:"target_workcode" url:"target_workcode"`
}

type RecordJournalUserInput struct {
	AppID          string `json:"app_id" url:"app_id" binding:"required"`
	ViewerWorkcode string `json:"viewer_workcode" url:"viewer_workcode" binding:"required"`
	TargetWorkcode string `json:"target_workcode" url:"target_workcode" binding:"required"`
	Type           string `json:"type,omitempty" url:"type"`
	Size           int    `json:"size,omitempty" url:"size"`
}

type HonorUserInfoInput struct {
	AppID    string `json:"app_id" url:"app_id" binding:"required"`
	Workcode string `json:"workcode" url:"workcode" binding:"required"`
	GroupId  string `json:"group_id" url:"group_id" binding:"required"`
}

type WeekReportUserInput struct {
	AppID          string `json:"app_id" url:"app_id" binding:"required"`
	ViewerWorkcode string `json:"viewer_workcode" url:"viewer_workcode" binding:"required"`
	TargetWorkcode string `json:"target_workcode" url:"target_workcode" binding:"required"`
	Size           int    `json:"size,omitempty" url:"size"`
}

type JobChannelInput struct {
	EffStatus  string `url:"eff_status" json:"eff_status,omitempty"`
	JobChannel string `url:"job_channel" json:"job_channel,omitempty"`
	Page       int    `url:"page" json:"page,omitempty"`
	PerPage    int    `url:"per_page" json:"per_page,omitempty"`
}

type JobcInput struct {
	JobCode    string `url:"job_code" json:"job_code,omitempty"`
	EffStatus  string `url:"eff_status" json:"eff_status,omitempty"`
	JobFamily  string `url:"job_family" json:"job_family,omitempty"`
	JobSeq     string `url:"job_seq" json:"job_seq,omitempty"`
	SubSeq     string `url:"sub_seq" json:"sub_seq,omitempty"`
	JobChannel string `url:"job_channel" json:"job_channel,omitempty"`
	Page       int    `url:"page" json:"page,omitempty"`
	PerPage    int    `url:"per_page" json:"per_page,omitempty"`
}

type JobSeqInput struct {
	EffStatus string `url:"eff_status" json:"eff_status,omitempty"`
	JobFamily string `url:"job_family" json:"job_family,omitempty"`
	JobSeq    string `url:"job_seq" json:"job_seq,omitempty"`
	Page      int    `url:"page" json:"page,omitempty"`
	PerPage   int    `url:"per_page" json:"per_page,omitempty"`
}

type JobSubInput struct {
	EffStatus string `url:"eff_status" json:"eff_status,omitempty"`
	JobFamily string `url:"job_family" json:"job_family,omitempty"`
	JobSeq    string `url:"job_seq" json:"job_seq,omitempty"`
	JobSubSeq string `url:"job_sub_seq" json:"job_sub_seq,omitempty"`
	Page      int    `url:"page" json:"page,omitempty"`
	PerPage   int    `url:"per_page" json:"per_page,omitempty"`
}

type QuitOrTalInput struct {
	Workcodes []string `json:"workcodes,omitempty"`
	StartTime string   `json:"time_start,omitempty"`
	EndTime   string   `json:"time_end,omitempty"`
}

type ScheduleCreateOrUpdateInput struct {
	AppId       string `json:"app_id"`
	BizId       string `json:"biz_id"`
	Title       string `json:"title"`
	StartTime   string `json:"start_time"`
	EndTime     string `json:"end_time"`
	Participant string `json:"participant"`
	Timestamp   string `json:"timestamp"`
	Workcode    string `json:"workcode,omitempty"`
	MeetingType string `json:"meeting_type,omitempty"`
	Address     string `json:"address"`
	RemindTime  string `json:"remind_time,omitempty"`
	Visibility  string `json:"visibility,omitempty"`
	Icon_link   string `json:"icon_link,omitempty"`
	Remark      string `json:"remark,omitempty"`
	TplCode     string `json:"tpl_code,omitempty"`
}

type ScheduleCancelInput struct {
	AppId     string `json:"app_id"`
	BizId     string `json:"biz_id"`
	Timestamp string `json:"timestamp"`
	Workcode  string `json:"workcode,omitempty"`
	Reason    string `json:"reason,omitempty"`
	Ticket    string `json:"ticket"`
}
