package dao

import (
	"resource-pub/internal/model"
	"resource-pub/pkg/app"
)


type Order struct {
	ID uint32 `json:"id"`
}

func (d *Dao) GetOrderList(types []int, keyword string, page, pageSize int) ([]*model.OrderRow, error) {
	order := model.Order{}
	return order.List(d.engine, types, keyword, app.GetPageOffset(page, pageSize), pageSize)
}