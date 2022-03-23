package dto

type SsoResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

type SsoVerifyResponse struct {
	Code int       `json:"code"`
	Msg  string    `json:"msg,omitempty"`
	Data *UserInfo `json:"data,omitempty"`
}

type UserInfo struct {
	AccountId string `json:"account_id,omitempty"`
	Account   string `json:"account,omitempty"`
	Name      string `json:"name,omitempty"`
	Workcode  string `json:"workcode,omitempty"`
	Email     string `json:"email,omitempty"`
	YachId    string `json:"yachid,omitempty"`
}

type JsConfigResponse struct {
	Code int       `json:"code"`
	Msg  string    `json:"msg,omitempty"`
	Data *JsConfig `json:"data,omitempty"`
}

type JsConfig struct {
	CorpId    int    `json:"corp_id"  binding:"required"`
	TimeStamp int    `json:"time_stamp"  binding:"required"`
	NonceStr  string `json:"nonce_str"  binding:"required"`
	Signature string `json:"signature"  binding:"required"`
	AppId     string `json:"appid"  binding:"required"`
	Url       string `json:"url"  binding:"required"`
	AgentId   string `json:"agentId"  binding:"required"`
}

type TicketResponse struct {
	Code int     `json:"code"`
	Msg  string  `json:"msg,omitempty"`
	Data *Ticket `json:"data,omitempty"`
}

type Ticket struct {
	Ticket string `json:"ticket"`
}
