package dto

type PizzaInput struct {
	Workcode string `json:"workcode" binding:"required"`
	Page     int `form:"page" json:"page" binding:"required_with=PageSize"`
	PageSize int `form:"pageSize" json:"pageSize" binding:"required_with=Page"`
}

type PaginationParams struct {
	Page     int `json:"page,omitempty"`
	PageSize int `json:"pageSize,omitempty"`
}

