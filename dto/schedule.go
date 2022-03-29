package dto

import (
	"errors"
	"github.com/gorhill/cronexpr"
	"time"
)

type ScheduleRegisterInput struct {
	ScheduleId     string      `json:"schedule_id"`
	ScheduleName   string      `json:"schedule_name"`
	CallbackParams interface{} `json:"callback_params"`
	Value          string      `json:"value"`
}

//Check registerType=0 value为 cron 表达式类型;1 - value为具体特定时间
func (s *ScheduleRegisterInput) Check(registerType int) error {
	if s.ScheduleId == ""{
		return errors.New("ScheduleRegisterInput ScheduleId empty")
	}

	if registerType == 0{
		_, err := cronexpr.Parse(s.Value)
		return err
	}

	if registerType == 1{
		_, err := time.ParseInLocation("2006-01-02 15:04", s.Value, time.Local)
		return err
	}
	return nil
}

type ScheduleRes struct {
	Code int       `json:"code"`
	Msg  string    `json:"msg"`
	Data *PostBody `json:"data"`
}

type PostBody struct {
	Tag            string `json:"schedule_id" binding:"required"`
	Name           string `json:"name"`
	CallbackParams string `json:"callback_params"`
}

type RegisterPostBody struct {
	ServiceName    string `json:"service_name"`
	ScheduleId     string `json:"schedule_id"`
	ScheduleName   string `json:"schedule_name"`
	CallbackUrl    string `json:"callback_url"`
	CallbackParams string `json:"callback_params"`
	Value          string `json:"value"`
	Type           int    `json:"type"`
}

type UnRegisterPostBody struct {
	ServiceName string `json:"service_name"`
	ScheduleId  string `json:"schedule_id"`
}
