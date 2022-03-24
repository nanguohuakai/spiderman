package dto

type PizzaResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

type EmployeeInfoRes struct {
	Code int           `json:"code"`
	Msg  string        `json:"msg,omitempty"`
	Data *EmployeeInfo `json:"data,omitempty"`
}

type ProjectListRes struct {
	Code int             `json:"code"`
	Msg  string          `json:"msg,omitempty"`
	Data *[]OwnedProject `json:"data,omitempty"`
}

type AppointmentItemRes struct {
	Code int              `json:"code"`
	Msg  string           `json:"msg,omitempty"`
	Data *AppointmentItem `json:"data,omitempty"`
}

type ExperienceRecordRes struct {
	Code int                 `json:"code"`
	Msg  string              `json:"msg,omitempty"`
	Data *[]ExperienceRecord `json:"data,omitempty"`
}

type TalRecordsRes struct {
	Code int            `json:"code"`
	Msg  string         `json:"msg,omitempty"`
	Data *[]BasicRecord `json:"data,omitempty"`
}

type ExpItemRes struct {
	Code int             `json:"code"`
	Msg  string          `json:"msg,omitempty"`
	Data *[]BasicExpItem `json:"data,omitempty"`
}
