package dto

type ScheduleInput struct {
	ScheduleId     string      `json:"schedule_id"`
	ScheduleName   string      `json:"schedule_name"`
	CallbackParams interface{} `json:"callback_params"`
	Value          string      `json:"value"`
}

type ScheduleRes struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Data *PostBody `json:"data"`
}

type PostBody struct {
	Tag            string `json:"schedule_id" binding:"required"`
	Name           string `json:"name"`
	CallbackParams string `json:"callback_params"`
}