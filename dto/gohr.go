package dto

type OkrUserListRes struct {
	Code int              `json:"code"`
	Msg  string           `json:"msg,omitempty"`
	Data UserListResponse `json:"data,omitempty"`
}

type RecordJournalYachRes struct {
	Code int                       `json:"code"`
	Msg  string                    `json:"msg,omitempty"`
	Data RecordJournalYachResponse `json:"data,omitempty"`
}

type HonorUserInfoRes struct {
	Code int              `json:"code"`
	Msg  string           `json:"msg,omitempty"`
	Data UserInfoResponse `json:"data,omitempty"`
}

type WeekReportYachRes struct {
	Code int                    `json:"code"`
	Msg  string                 `json:"msg,omitempty"`
	Data WeekReportYachResponse `json:"data,omitempty"`
}

type JobChannelDataRes struct {
	Code int            `json:"code"`
	Msg  string         `json:"msg,omitempty"`
	Data JobChannelData `json:"data,omitempty"`
}

type JobcDataRes struct {
	Code int      `json:"code"`
	Msg  string   `json:"msg,omitempty"`
	Data JobcData `json:"data,omitempty"`
}

type JobSeqDataRes struct {
	Code int        `json:"code"`
	Msg  string     `json:"msg,omitempty"`
	Data JobSeqData `json:"data,omitempty"`
}
