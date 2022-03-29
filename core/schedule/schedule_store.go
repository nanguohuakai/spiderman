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

func (c *Client) RegisterCorn(input dto.ScheduleInput) error {
	if ! c.ScheduleConf.GetIsRegister() {
		return errors.New("未注册")
	}
	u := c.ScheduleConf.BaseUri + "/api/v1/schedule/register"
	var postBody = postBody{
		ServiceName: c.AppConf.ServiceName,
		ScheduleId: input.ScheduleId,
		ScheduleName: input.ScheduleName,
		Value: input.Value,
		Type: 0,
	}

	postBody.CallbackParams = ""
	if input.CallbackParams != nil {
		params , err := json.Marshal(input.CallbackParams)
		postBody.CallbackParams = string(params)
		if err != nil {
			postBody.CallbackParams = ""
		}
	}

	postBody.CallbackUrl = c.ScheduleConf.CallbackUri + c.ScheduleConf.GetCallbackPath()

	var output interface{}
	return httpclient.Post(u, c.AppConf ,postBody, &output)
}


func (c *Client)RegisterHandler(r *gin.Engine, f gin.HandlerFunc, relativePath string) error {

	if  c.ScheduleConf.GetIsRegister() {
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

func  (c *Client)createScheduleAuth(name, secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		clientToken := strings.Split(token, " ")

		if len(clientToken) != 2 || name != clientToken[0] || secret != clientToken[1] {
			//log.Error("ScheduleAuth", log.String("params", fmt.Sprintf("input: token - %#v", token)))
			//alert.Send(fmt.Sprintf("没有api权限: %s, token:%s", c.Request.RequestURI, token), "error")
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "没有api权限",
			})
			c.Abort()
		}
		c.Next()

	}
}