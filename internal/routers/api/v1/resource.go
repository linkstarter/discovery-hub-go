package v1

import (
	"log"
	"resource-pub/internal/routers/service"
	"resource-pub/pkg/app"

	"github.com/gin-gonic/gin"
)


type Resource struct{}

func NewResource() Resource {
	return Resource{}
}

func (r Resource) Lists(c *gin.Context) {
	params := service.ResourceListRequest{}
	response := app.NewResource(c)
	err := c.ShouldBind(&params)
	if err != nil {
		log.Fatalf("params err:%v", err)
	}
	
}