package app

import "github.com/gin-gonic/gin"

type Response struct {
	Ctx *gin.Context
}

type Pager struct {
	// 页码
	Page int `json:"page"`
	// 每页数量
	PerPage int `json:"perPage"`
	// 总行数
	Total int `json:"total"`
}

func NewResource(ctx *gin.Context) *Response {
	return &Response{
		Ctx: ctx,
	}
}