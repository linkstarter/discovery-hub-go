package routers

import (
	"resource-pub/global"
	v1 "resource-pub/internal/routers/api/v1"

	"github.com/gin-gonic/gin"
)


func NewRouter() *gin.Engine {
	r := gin.New()
	if global.ServerSetting.RunMode == "debug" {
		r.Use(gin.Logger())
		r.Use(gin.Recovery())
	}
	resource := v1.NewResource()
	order := v1.NewOrder()
	g := r.Group("pub/resource")
	g.Use()
	{
		g.GET("/lists", resource.Lists)
		g.GET("order-lists", order.Lists)
	}
	return r
}