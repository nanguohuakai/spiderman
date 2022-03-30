package schedule

import (
	"github.com/gin-gonic/gin"
	"github.com/nanguohuakai/spiderman/dto"
	"strings"
)

//ScheduleInterface support pipelining using the RegisterCorn, RegisterCornTime, RegisterHandler and UnRegister methods
type ScheduleInterface interface {
	RegisterCorn(input dto.ScheduleRegisterInput) error                          //RegisterCorn 表达式类型定时任务 例子 :"*/1 * * * *"
	RegisterCornTime(input dto.ScheduleRegisterInput) error                      //RegisterCornTime 特定时间定时任务
	RegisterHandler(r *gin.Engine, f gin.HandlerFunc, relativePath string) error //RegisterHandler 注册回调路由
	UnRegister(scheduleId string) error                                          //UnRegister 取消定时任务
}

type Client struct {
	AppConf      dto.AppConf
	ScheduleConf dto.ScheduleConf
}

//The NewClient interface is the primary interface for working with ScheduleInterface
func NewClient(conf dto.AppConf, scheduleConf dto.ScheduleConf) ScheduleInterface {
	//特殊去掉域名最后元素"/"
	scheduleConf.BaseUri = strings.TrimRight(scheduleConf.BaseUri, "/")
	scheduleConf.CallbackUri = strings.TrimRight(scheduleConf.CallbackUri, "/")
	var uc ScheduleInterface
	uc = &Client{
		AppConf:      conf,
		ScheduleConf: scheduleConf,
	}
	return uc
}
