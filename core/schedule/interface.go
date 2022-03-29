package schedule

import (
	"github.com/gin-gonic/gin"
	"github.com/nanguohuakai/spiderman/dto"
)



type ScheduleInterface interface {
	RegisterCorn(input dto.ScheduleInput) error
	RegisterHandler(r *gin.Engine, f gin.HandlerFunc, relativePath string) error
	//RegisterTime()
	//UnRegister()
}

type Client struct {
	AppConf      dto.AppConf
	ScheduleConf dto.ScheduleConf
}

func NewClient(conf dto.AppConf, scheduleConf dto.ScheduleConf) ScheduleInterface {
	var uc ScheduleInterface
	uc = &Client{
		AppConf:      conf,
		ScheduleConf: scheduleConf,
	}
	return uc
}

type postBody struct {
	ScheduleId     string `json:"schedule_id"`
	ScheduleName   string `json:"schedule_name"`
	CallbackUrl    string `json:"callback_url"`
	CallbackParams string `json:"callback_params"`
	Value          string `json:"value"`
	Type           int    `json:"type"`
}
