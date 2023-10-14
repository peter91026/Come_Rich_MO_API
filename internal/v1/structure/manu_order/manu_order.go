package manu_order

import (
	"time"
	rawMaterial "eirc.app/internal/v1/structure/raw_material"
	model "eirc.app/internal/v1/structure"
)

type Table struct {
		ManuOrderID    string    `gorm:"column:manu_order_id;type:text;primary_key" json:"manu_order_id"`
		GoodsName      string    `gorm:"column:goods_name;type:text" json:"goods_name"`
		Identifier     string    `gorm:"column:identifier;type:text" json:"identifier"`
		TotalQuantity  string    `gorm:"column:total_quantity;type:text" json:"total_quantity"`
		CompletionDate time.Time `gorm:"column:completion_date;type:datetime" json:"completion_date"`
		DateOfIssuance time.Time `gorm:"column:date_of_issuance;type:datetime" json:"date_of_issuance"`
		CreatedAt       time.Time `gorm:"column:created_at;type:timestamp" json:"create_at"`
		UpdatedAt      time.Time `gorm:"column:updated_at;type:timestamp" json:"updated_at"`
		IsDeleted      bool      `gorm:"column:is_deleted;type:bool" json:"is_deleted"`

		RawMaterial []*rawMaterial.Table `gorm:"foreignkey:manu_order_id;references:manu_order_id"` //reference:自己資料表的欄位名稱  //foreignkey:關聯表的對應欄位名稱
	}


// 結構基底(SHOW/COPY)
type Base struct {
	ManuOrderID    string    `json:"manu_order_id"`
	GoodsName      string    `json:"goods_name"`
	Identifier     string    `json:"identifier"`
	TotalQuantity  string    `json:"total_quantity"`
	CompletionDate time.Time `json:"completion_date"`
	DateOfIssuance time.Time `json:"date_of_issuance"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	IsDeleted      bool      `json:"is_deleted"`

	RawMaterial []*rawMaterial.Base
}

// 清單顯示&查詢用
type Field struct {
	ManuOrderID    *string    `json:"manu_order_id,omitempty"`
	GoodsName      string    `json:"goods_name,omitempty"`
	Identifier     string    `json:"identifier,omitempty"`
	TotalQuantity  string    `json:"total_quantity,omitempty"`
	CompletionDate time.Time `json:"completion_date,omitempty"`
	DateOfIssuance time.Time `json:"date_of_issuance,omitempty"`
	CreatedAt       time.Time `json:"created_at,omitempty"`
	UpdatedAt      time.Time `json:"updated_at,omitempty"`
	IsDeleted      *bool      `json:"is_deleted,omitempty"`

	RawMaterial []*rawMaterial.Field
}

type Fields struct {
	Field
	model.InPage
}

type List struct {
	ManuOrders []*struct {
		Field
	} `json:"manu_orders"`
	model.OutPage
}

// 新增版本
type Created struct {
	ManuOrderID    string    `json:"manu_order_id,omitempty"`
	GoodsName      string    `json:"goods_name,omitempty"`
	Identifier     string    `json:"identifier,omitempty"`
	TotalQuantity  string    `json:"total_quantity,omitempty"`
	CompletionDate time.Time `json:"completion_date,omitempty"`
	DateOfIssuance time.Time `json:"date_of_issuance,omitempty"`
	CreatedAt       time.Time `json:"created_at,omitempty"`
	UpdatedAt      time.Time `json:"updated_at,omitempty"`
	IsDeleted      bool      `json:"is_deleted,omitempty"`
}

type Updated struct {
	ManuOrderID    string    `json:"manu_order_id,omitempty"`
	GoodsName      string    `json:"goods_name,omitempty"`
	Identifier     string    `json:"identifier,omitempty"`
	TotalQuantity  string    `json:"total_quantity,omitempty"`
	CompletionDate time.Time `json:"completion_date,omitempty"`
	DateOfIssuance time.Time `json:"date_of_issuance,omitempty"`
	CreatedAt       time.Time `json:"created_at,omitempty"`
	UpdatedAt      time.Time `json:"updated_at,omitempty"`
	IsDeleted      bool      `json:"is_deleted,omitempty"`
}

func (a *Table) TableName() string {
	return "manu_order"
}
