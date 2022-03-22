package dto

type SsoResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

type SsoVerifyResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg,omitempty"`
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