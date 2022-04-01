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

type EmployeeListRes struct {
	Code int            `json:"code"`
	Msg  string         `json:"msg,omitempty"`
	Data []EmployeeInfo `json:"data,omitempty"`
}

type EmployeeListWithPageRes struct {
	Code int                  `json:"code"`
	Msg  string               `json:"msg,omitempty"`
	Data EmployeeListWithPage `json:"data,omitempty"`
}

type EmployeeListWithPage struct {
	Page     int            `json:"page"`
	PageSize int            `json:"pageSize"`
	Total    int64          `json:"total"`
	List     []EmployeeInfo `json:"list"`
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

type PromotionTrackRecordRes struct {
	Code int                     `json:"code"`
	Msg  string                  `json:"msg,omitempty"`
	Data *[]PromotionTrackRecord `json:"data,omitempty"`
}

type EduItemRes struct {
	Code int        `json:"code"`
	Msg  string     `json:"msg,omitempty"`
	Data *[]EduItem `json:"data,omitempty"`
}

type KpiListRes struct {
	Code int        `json:"code"`
	Msg  string     `json:"msg,omitempty"`
	Data *[]KpiItem `json:"data,omitempty"`
}

type RewardListRes struct {
	Code int           `json:"code"`
	Msg  string        `json:"msg,omitempty"`
	Data *[]RewardItem `json:"data,omitempty"`
}

type DeptListRes struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg,omitempty"`
	Data *[]DeptInfo `json:"data,omitempty"`
}

type DeptListWithPageRes struct {
	Code int              `json:"code"`
	Msg  string           `json:"msg,omitempty"`
	Data DeptListWithPage `json:"data,omitempty"`
}

type DeptListWithPage struct {
	Page     int        `json:"page"`
	PageSize int        `json:"pageSize"`
	Total    int64      `json:"total"`
	List     []DeptInfo `json:"list"`
}

type DeptEmployeCountRes struct {
	Code int           `json:"code"`
	Msg  string        `json:"msg,omitempty"`
	Data EmployeeCount `json:"data,omitempty"`
}

type HrpsGrRes struct {
	Code int           `json:"code"`
	Msg  string        `json:"msg,omitempty"`
	Data *BehaviorItem `json:"data,omitempty"`
}

type HrpsGrListRes struct {
	Code int            `json:"code"`
	Msg  string         `json:"msg,omitempty"`
	Data []BehaviorItem `json:"data,omitempty"`
}

type HrpsRes struct {
	Code int      `json:"code"`
	Msg  string   `json:"msg,omitempty"`
	Data *OrgItem `json:"data,omitempty"`
}

type HrpsListRes struct {
	Code int       `json:"code"`
	Msg  string    `json:"msg,omitempty"`
	Data []OrgItem `json:"data,omitempty"`
}

type CulturalListRes struct {
	Code int            `json:"code"`
	Msg  string         `json:"msg,omitempty"`
	Data []CulturalItem `json:"data,omitempty"`
}

type ProjectUsersRes struct {
	Code int              `json:"code"`
	Msg  string           `json:"msg,omitempty"`
	Data []ProjectGeneral `json:"data,omitempty"`
}

type ProjectListWithPageRes struct {
	Code int                   `json:"code"`
	Msg  string                `json:"msg,omitempty"`
	Data ProjectListWithPage `json:"data,omitempty"`
}
