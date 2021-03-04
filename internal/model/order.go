package model

import "github.com/jinzhu/gorm"

type Order struct {
	ID uint32 `json:"id"`
}

type OrderRow struct {
	ID uint32
}

func (o Order) TableName() string {
	return "public_order"
}

func (o Order) List(db *gorm.DB, types []int, keyword string, pageOffset, pageSize int) ([]*OrderRow, error) {
	cols := []string{"id"}

	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	rows, err := db.Select(cols).Table(o.TableName() + " AS o").Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []*OrderRow
	for rows.Next() {
		r := &OrderRow{}
		if err := rows.Scan(&r.ID); err != nil {
			return nil, err
		}
		orders = append(orders, r)
	}
	return orders, nil
}