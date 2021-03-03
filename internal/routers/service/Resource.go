package service

import "resource-pub/pkg/app"

type ResourceListRequest struct {
	Type    string `form:"type"`
	Keyword string `form:"keyword"`
	Sort    string `form:"sort"`
	app.Pager
}
