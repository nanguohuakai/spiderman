package schedule

import (
	"github.com/gin-gonic/gin"
	"github.com/nanguohuakai/spiderman/dto"
	"strings"
)

type ScheduleInterface interface {
	RegisterCorn(input dto.ScheduleRegisterInput) error
	RegisterHandler(r *gin.Engine, f gin.HandlerFunc, relativePath string) error //RegisterHandler 注册回调路由
	RegisterCornTime(input dto.ScheduleRegisterInput) error
	UnRegister(scheduleId string) error //UnRegister 取消定时任务
}

type Client struct {
	AppConf      dto.AppConf
	ScheduleConf dto.ScheduleConf
}

func NewClient(conf dto.AppConf, scheduleConf dto.ScheduleConf) ScheduleInterface {
	scheduleConf.BaseUri = strings.TrimRight(scheduleConf.BaseUri, "/")
	scheduleConf.CallbackUri = strings.TrimRight(scheduleConf.CallbackUri, "/")
	var uc ScheduleInterface
	uc = &Client{
		AppConf:      conf,
		ScheduleConf: scheduleConf,
	}
	return uc
}
