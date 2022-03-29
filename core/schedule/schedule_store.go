package schedule

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/nanguohuakai/spiderman/dto"
	"github.com/nanguohuakai/spiderman/pkg/httpclient"
	"net/http"
	"strings"
)
const(
	//RegisterTypeCron = 0, cron 表达式类型 例子 :"*/5 * * * *"
	RegisterTypeCron = 0
	//RegisterTypeTime = 1, 特定时间 例子: "2022-10-11 12:12" 精确到分
	RegisterTypeTime = 1
)

func (c *Client) RegisterCorn(input dto.ScheduleRegisterInput) error {
	return c.schedulePost(input, RegisterTypeCron)
}

func (c *Client) RegisterCornTime(input dto.ScheduleRegisterInput) error {
	return c.schedulePost(input, RegisterTypeTime)
}

//schedulePost 请求schedule
func (c *Client) schedulePost(input dto.ScheduleRegisterInput, registerType int) error {
	if !c.ScheduleConf.GetIsRegister() {
		return errors.New("未注册")
	}

	if err := input.Check(registerType); err != nil {
		return err
	}

	u := c.ScheduleConf.BaseUri + "/api/v1/schedule/register"
	postBody := dto.RegisterPostBody{
		ServiceName:  c.AppConf.ServiceName,
		ScheduleId:   input.ScheduleId,
		ScheduleName: input.ScheduleName,
		Value:        input.Value,
		Type:         registerType,
	}
	//input.CallbackParams 为nil， postBody.CallbackParams处理为空字符
	//input.CallbackParams 不为nil，postBody.CallbackParams处理为字符串
	postBody.CallbackParams = ""
	if input.CallbackParams != nil {
		params, err := json.Marshal(input.CallbackParams)
		postBody.CallbackParams = string(params)
		if err != nil {
			postBody.CallbackParams = ""
		}
	}

	postBody.CallbackUrl = c.ScheduleConf.CallbackUri + c.ScheduleConf.GetCallbackPath()
	//var uu, err = url.Parse(c.ScheduleConf.CallbackUri)
	//if err != nil {
	//	return err
	//}
	//
	//uu.Path = path.Join(uu.Path, c.ScheduleConf.GetCallbackPath())
	//postBody.CallbackUrl= uu.String()
	var output interface{}
	return httpclient.Post(u, c.AppConf, postBody, &output)
}

//RegisterHandler 注册回调路由
func (c *Client) RegisterHandler(r *gin.Engine, f gin.HandlerFunc, relativePath string) error {

	if c.ScheduleConf.GetIsRegister() {
		return errors.New("已注册")
	}
	path := relativePath
	if relativePath == "" {
		path = "/jobs"
	}
	r.POST(path, c.createScheduleAuth(c.AppConf.ServiceName, c.AppConf.Token), f)

	c.ScheduleConf.SetCallbackPath(path)
	c.ScheduleConf.SetIsRegister(true)

	return nil
}

//createScheduleAuth 权限校验
func (c *Client) createScheduleAuth(name, secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		clientToken := strings.Split(token, " ")

		if len(clientToken) != 2 || name != clientToken[0] || secret != clientToken[1] {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "没有api权限",
			})
			c.Abort()
		}
		c.Next()
	}
}

//UnRegister 取消定时任务
func (c *Client) UnRegister(scheduleId string) error {
	u := c.ScheduleConf.BaseUri + "/api/v1/schedule/unregister"
	postBody := dto.UnRegisterPostBody{
		ServiceName: c.AppConf.ServiceName,
		ScheduleId:  scheduleId,
	}

	var output interface{}
	return httpclient.Post(u, c.AppConf, postBody, &output)
}
