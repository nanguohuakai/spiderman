package dto

import "time"

type PaginationOps struct {
	Page     int `form:"page" json:"page" binding:"required_with=PageSize"`
	PageSize int `form:"pageSize" json:"pageSize" binding:"required_with=Page"`
}

type DurationOps struct {
	StartDate time.Time `form:"startDate" json:"startDate" binding:"required_with=EndDate" time_format:"2006-01-02"`
	EndDate   time.Time `form:"endDate" json:"endDate"  binding:"required_with=StartDate" time_format:"2006-01-02"`
}

type WildSearchOps struct {
	WildVal    string   `form:"wildVal" json:"wildVal"`
	WildFields []string `form:"wildFields" json:"wildFields" binding:"required_with=WildVal"`
}

type OrderOps struct {
	OrderField string `form:"orderField" json:"orderField"`
	Direction  string `form:"direction" json:"direction"`
}

type TermFiltersOps struct {
	FilterField string `form:"filterField" json:"filterField"`
	FilterVal   string `form:"filterVal" json:"filterVal"`
}

type ScopeFiltersOps struct {
	FilterScopes []string `form:"filterScopes" json:"filterScopes"`
	FilterField  string   `form:"scopeFilterField" json:"scopeFilterField"`
}

type WildMatchOps struct {
	TargetField string   `form:"targetField" json:"targetField"`
	MatchValues []string `form:"matchValues" json:"matchValues"`
	MatchPos    string   `form:"matchPos" json:"matchPos"`
}
