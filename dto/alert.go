package dto

type AlertParams struct {
	ServiceName string    `json:"service_name"`
	Body        AlertBody `json:"body"`
}

type AlertBody struct {
	Level   string `json:"level"`
	Env     string `json:"env"`
	Content string `json:"content"`
}

type AlertRes struct {
	Code int    `json:"code"`
	Msg  string `json:"msg,omitempty"`
}
