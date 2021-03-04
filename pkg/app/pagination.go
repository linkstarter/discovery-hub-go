package app

import (
	"resource-pub/global"
	"resource-pub/pkg/convert"

	"github.com/gin-gonic/gin"
)


func GetPage(c *gin.Context) int {
	page := convert.StrTo(c.Query("page")).MustInt()
	if page <= 0 {
		return 1
	}
	return page
}

func GetPageSize(c *gin.Context) int {
	pageSize := convert.StrTo(c.Query("pageSize")).MustInt()
	if pageSize <= 0 {
		return global.AppSetting.DefaultPageSize
	}
	return pageSize
}

func GetPageOffset(page, pageSize int) int {
	result := 0
	if page > 0 {
		result = (page - 1) * pageSize
	}

	return result
}