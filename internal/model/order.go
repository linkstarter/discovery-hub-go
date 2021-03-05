package model

import (
	"github.com/jinzhu/gorm"
)

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

type OrderRow struct {
	ID uint32
	MainID uint32
	ResourceId string
	Title string
	Username string
	Realname string
	TenantID string
	Type uint8
	CreatedAt uint32
	Suffix string
	FileSize string
	Status uint8
	Reason string
	AIStatus uint8
}

func (o Order) TableName() string {
	return "public_order"
}

func (o Order) List(db *gorm.DB, types []string, keyword string, pageOffset, pageSize int) ([]*OrderRow, error) {
	cols := []string{"id", "main_id", "resource_id", "title", "username", "realname", "tenant_id", "type", "created_at"}
	cols = append(cols, []string{"status", "suffix", "filesize", "ai_status", "reason"}...)
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	db = db.Select(cols).Table(o.TableName() + " AS o")
	if keyword != "" {
		db = db.Where("o.`title` LIKE ?", keyword)
	}
	if len(types) > 0 {
		db = db.Where("o.`type` IN (?)", types)
	}

	rows, err := db.Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []*OrderRow
	for rows.Next() {
		r := &OrderRow{}
		if err := rows.Scan(&r.ID, &r.MainID, &r.ResourceId, &r.Title, 
			&r.Username, &r.Realname, &r.TenantID, &r.Type, &r.CreatedAt,
			&r.Status, &r.Suffix, &r.FileSize, &r.AIStatus, &r.Reason); err != nil {
			return nil, err
		}
		orders = append(orders, r)
	}
	return orders, nil
}

func (o Order) Count(db *gorm.DB, types []string, keyword string) (int, error) {
	var count int
	db = db.Table(o.TableName() + " AS o")
	if keyword != "" {
		db = db.Where("o.`title` LIKE ?", keyword)
	}
	if len(types) > 0 {
		db = db.Where("o.`type` IN (?)", types)
	}

	err := db.Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
