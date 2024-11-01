package models

type Response struct {
	Data       interface{} `json:"data"`
	TotalData  int64       `json:"totalData"`
	Page       int         `json:"page"`
	PageSize   int         `json:"pageSize"`
	TotalPages int         `json:"totalPages"`
}
