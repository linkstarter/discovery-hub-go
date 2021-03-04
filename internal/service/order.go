package service

import (
	"resource-pub/pkg/app"
	"resource-pub/pkg/convert"
	"strings"
)


type OrderListRequest struct {
	Type    string `form:"type"`
	Keyword string `form:"keyword"`
	Sort    string `form:"sort"`
}

type Order struct {
	ID uint32 `json:"id"`
}

func (svc *Service) GetOrderList(param *OrderListRequest, pager *app.Pager) ([]*Order, int, error) {
	strTypes := strings.Split(param.Type, ",")
	intTypes := []int{}
	for _, v := range strTypes {
		intTypes = append(intTypes, convert.StrTo(v).MustInt())
	}

	orders, err := svc.dao.GetOrderList(intTypes, param.Keyword, pager.Page, pager.PageSize)
	if err != nil {
		return nil, 0, err
	}

	var orderList []*Order
	for _, v := range orders {
		orderList = append(orderList, &Order{
			ID: v.ID,
		})
	}

	return orderList, len(orderList), nil
}