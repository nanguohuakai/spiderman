package dto

type PizzaInput struct {
	Workcode string `json:"workcode" binding:"required"`
	Page     int    `form:"page" json:"page" binding:"required_with=PageSize"`
	PageSize int    `form:"pageSize" json:"pageSize" binding:"required_with=Page"`
}

type PaginationParams struct {
	Page     int `json:"page,omitempty"`
	PageSize int `json:"pageSize,omitempty"`
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
