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
