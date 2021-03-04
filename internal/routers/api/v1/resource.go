package v1

import (
	"log"
	"resource-pub/internal/service"
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
	res := service.Resource{
		ID: 1,
		ResourceID: "T544",
		Title: "test",
	}
	// pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	response.ToResponseList(res, 1)
}