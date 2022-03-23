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
