package dto

type PizzaInput struct {
	Workcode string `json:"workcode" binding:"required"`
	Page     int    `form:"page" json:"page" binding:"required_with=PageSize"`
	PageSize int    `form:"pageSize" json:"pageSize" binding:"required_with=Page"`
}

type PizzaWorkcodesInput struct {
	Workcode []string `form:"workcode[],omitempty" json:"workcode,omitempty" url:"workcode[],omitempty"`
	Page     int      `form:"page" json:"page" url:"page" binding:"required_with=PageSize"`
	PageSize int      `form:"pageSize" json:"pageSize" url:"pageSize" binding:"required_with=Page"`
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

type HrpsInput struct {
	Workcode  string `json:"workcode"`
	StartDate string `json:"start_date"` //开始日期 2006-01-02
	EndDate   string `json:"end_date"`   //结束日期 2006-01-02
}

type CulturalListInput struct {
	Workcodes string `json:"workcodes"`
	StartYear string `json:"startYear"` //开始日期 2006
	EndYear   string `json:"endYear"`   //结束日期 2006
}

type UserProjectInput struct {
	WorkCode  string   `url:"workcode" json:"workcode" binding:"required"`
	Status    []string `url:"status[]" json:"status" binding:"required"`
	StartDate string   `url:"start_date" json:"start_date" binding:"required"`
	EndDate   string   `url:"end_date" json:"end_date"`
}

type ProjectListInput struct {
	WorkCode  string `url:"workcode" json:"workcode" binding:"required"`
	Status    string `url:"status" json:"status" binding:"required"`
	StartDate string `url:"start_date" json:"start_date" binding:"required"`
	EndDate   string `url:"end_date" json:"end_date"`
	Page      int    `url:"page" json:"page" binding:"required_with=PageSize"`
	PageSize  int    `url:"pageSize" json:"pageSize" binding:"required_with=Page"`
}

type BuildListInput struct {
	CountryCode  string `json:"country_code"`
	ProvinceCode string `json:"province_code"`
	CityCode     string `json:"city_code"`
	DistrictCode string `json:"district_code"`
	BuildingCode string `json:"building_code"`
	FloorCode    string `json:"floor_code"`
}
