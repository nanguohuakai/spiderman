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

type JobSubDataRes struct {
	Code int        `json:"code"`
	Msg  string     `json:"msg,omitempty"`
	Data JobSubData `json:"data,omitempty"`
}

type QuitDataRes struct {
	Code int      `json:"code"`
	Msg  string   `json:"msg,omitempty"`
	Data QuitData `json:"data,omitempty"`
}

type TalDataRes struct {
	Code int     `json:"code"`
	Msg  string  `json:"msg,omitempty"`
	Data TalData `json:"data,omitempty"`
}

type OpsScheduleGoHrRes struct {
	Code int         `json:"errcode"`
	Msg  string      `json:"errmsg"`
	Data interface{} `json:"data"`
}

type OpsScheduleLisRes struct {
	Code int                    `json:"errcode"`
	Msg  string                 `json:"errmsg"`
	Data map[string]interface{} `json:"data"`
}

type Report360DataRes struct {
	Code int           `json:"errcode"`
	Msg  string        `json:"errmsg"`
	Data Report360Data `json:"data"`
}
