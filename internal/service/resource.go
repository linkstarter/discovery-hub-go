package service


type ResourceListRequest struct {
	Type    string `form:"type"`
	Keyword string `form:"keyword"`
	Sort    string `form:"sort"`
}

type Resource struct {
	ID         uint32 `json:"id"`
	ResourceID string `json:"resourceId"`
	Title      string `json:"title"`
}

// 获取资源列表
// func (svc *Service) GetResourceList(params *ResourceListRequest, pager *app.Pager) ([]*Resource, int, error) {
	
// }