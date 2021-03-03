package routers

import (
	"resource-pub/global"

	"github.com/gin-gonic/gin"
)


func NewRouter() *gin.Engine {
	r := gin.New()
	if global.ServerSetting.RunMode == "debug" {
		r.Use(gin.Logger())
		r.Use(gin.Recovery())
	}

	g := r.Group("pub/resource")
	g.Use()
	{
		g.GET("/lists",)
	}
}