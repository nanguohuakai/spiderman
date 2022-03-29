package dto

type PizzaInput struct {
	Workcode string `json:"workcode" binding:"required"`
	Page     int    `form:"page" json:"page" binding:"required_with=PageSize"`
	PageSize int    `form:"pageSize" json:"pageSize" binding:"required_with=Page"`
}

type PaginationParams struct {
	Page     int `json:"page,omitempty" url:"page,omitempty"`
	PageSize int `json:"pageSize,omitempty" url:"pageSize,omitempty"`
}

type EmployeeListInput struct {
	WildSearchOps   *WildSearchOps    `json:"wild_search"`
	WildMatchOps    *WildMatchOps     `json:"wild_match"`
	OrderOps        []OrderOps        `json:"order_ops"`
	TermFiltersOps  []TermFiltersOps  `json:"term_filter_ops"`
	ScopeFiltersOps []ScopeFiltersOps `json:"scope_filter_ops"`
	TermExcludeOps  []TermFiltersOps  `json:"term_exclude_ops"`
}

type EmployeeListInputWithPage struct {
	Pagination      PaginationOps     `json:"pagination"`
	WildSearchOps   *WildSearchOps    `json:"wild_search"`
	WildMatchOps    *WildMatchOps     `json:"wild_match"`
	OrderOps        []OrderOps        `json:"order_ops"`
	TermFiltersOps  []TermFiltersOps  `json:"term_filter_ops"`
	ScopeFiltersOps []ScopeFiltersOps `json:"scope_filter_ops"`
	TermExcludeOps  []TermFiltersOps  `json:"term_exclude_ops"`
}

type EmployeeListV2Input struct {
	OrderOps   []OrderOps                 `json:"order_ops"`
	Conditions []EmployeeListFilterStruct `json:"conditions"`
}

type EmployeeListV2InputWithPage struct {
	Pagination PaginationOps              `json:"pagination"`
	OrderOps   []OrderOps                 `json:"order_ops"`
	Conditions []EmployeeListFilterStruct `json:"conditions"`
}

type EmployeeListFilterStruct struct {
	WildSearchOps   *WildSearchOps    `json:"wild_search"`
	WildMatchOps    *WildMatchOps     `json:"wild_match"`
	TermFiltersOps  []TermFiltersOps  `json:"term_filter_ops"`
	ScopeFiltersOps []ScopeFiltersOps `json:"scope_filter_ops"`
	TermExcludeOps  []TermFiltersOps  `json:"term_exclude_ops"`
}

type RewardInput struct {
	Types     []string `json:"types[],omitempty" url:"types[],omitempty"`
	StartDate *string  `json:"startDate,omitempty" url:"startDate,omitempty"`
	EndDate   *string  `json:"endDate,omitempty"  url:"endDate,omitempty"`
}

type DeptListInput struct {
	WildSearchOps   *WildSearchOps    `json:"wild_search"`
	WildMatchOps    *WildMatchOps     `json:"wild_match"`
	OrderOps        []OrderOps        `json:"order_ops"`
	TermFiltersOps  []TermFiltersOps  `json:"term_filter_ops"`
	ScopeFiltersOps []ScopeFiltersOps `json:"scope_filter_ops"`
}

type DeptListInputWithPage struct {
	Pagination      PaginationOps     `json:"pagination"`
	WildSearchOps   *WildSearchOps    `json:"wild_search"`
	WildMatchOps    *WildMatchOps     `json:"wild_match"`
	OrderOps        []OrderOps        `json:"order_ops"`
	TermFiltersOps  []TermFiltersOps  `json:"term_filter_ops"`
	ScopeFiltersOps []ScopeFiltersOps `json:"scope_filter_ops"`
}
