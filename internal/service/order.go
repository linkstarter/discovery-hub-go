package service

import (
	"resource-pub/pkg/app"
	"strings"
)

type OrderListRequest struct {
	Type    string `form:"type"`
	Keyword string `form:"keyword"`
	Sort    string `form:"sort"`
}

type Order struct {
	ID uint32 `json:"id"`
	MainID uint32 `json:"main_id"`
	ResourceId string `json:"resourceId"`
	Title string `json:"title"`
	Username string `json:"username"`
	Realname string `json:"realname"`
	TenantID string `json:"tenantId"`
	Type uint8 `json:"type"`
	CreatedAt uint32 `json:"createdAt"`
	Suffix string `json:"suffix"`
	FileSize string `json:"fileSize"`
	Status uint8	`json:"status"`
	Reason string `json:"reason"`
	AIStatus uint8 `json:"aistatus"`
}

func (svc *Service) GetOrderList(param *OrderListRequest, pager *app.Pager) ([]*Order, int, error) {
	var strTypes []string

	if param.Type != "" {
		strTypes = strings.Split(param.Type, ",")
	}
	orders, err := svc.dao.GetOrderList(strTypes, param.Keyword, pager.Page, pager.PageSize)
	if err != nil {
		return nil, 0, err
	}

	var orderList []*Order
	for _, v := range orders {
		orderList = append(orderList, &Order{
			ID: v.ID,
			MainID: v.MainID,
			ResourceId: v.ResourceId,
			Title: v.Title,
			Username: v.Username,
			Realname: v.Realname,
			TenantID: v.TenantID,
			Type: v.Type,
			FileSize: v.FileSize,
			Suffix: v.Suffix,
			Reason: v.Reason,
			Status: v.Status,
		})
	}

	count, err := svc.dao.GetOrderCount(strTypes, param.Keyword)
	if err != nil {
		return nil, 0, err
	}
	return orderList, count, nil
}
