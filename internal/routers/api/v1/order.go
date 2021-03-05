package v1

import (
	"log"
	"resource-pub/internal/service"
	"resource-pub/pkg/app"

	"github.com/gin-gonic/gin"
)

type Order struct{}

func NewOrder() Order {
	return Order{}
}

func (o Order) Lists(c *gin.Context) {
	params := service.OrderListRequest{}
	response := app.NewResource(c)
	err := c.ShouldBind(&params)
	if err != nil {
		log.Fatalf("params err:%v", err)
	}
	svc := service.New(c.Request.Context())
	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	orders, total, err := svc.GetOrderList(&params, &pager)
	if err != nil {
		response.ToErrorResponse("错误")
	}
	response.ToResponseList(orders, total)
}
