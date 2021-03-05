package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Ctx *gin.Context
}

type Pager struct {
	// 页码
	Page int `json:"page"`
	// 每页数量
	PageSize int `json:"perPage"`
	// 总行数
	Total int `json:"total"`
}

func NewResource(ctx *gin.Context) *Response {
	return &Response{
		Ctx: ctx,
	}
}

func (r *Response) ToResponseList(list interface{}, totalRows int) {
	r.Ctx.JSON(http.StatusOK, gin.H{
		"list": list,
		"pager": Pager{
			Page:     GetPage(r.Ctx),
			PageSize: GetPageSize(r.Ctx),
			Total:    totalRows,
		},
	})
}

func (r *Response) ToErrorResponse(msg string) {
	response := gin.H{"code": 001, "msg": msg}
	r.Ctx.JSON(500, response)
}
