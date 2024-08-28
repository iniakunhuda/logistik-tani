package model

import (
	"time"
)

type Sales struct {
	ID          uint               `gorm:"primaryKey;autoIncrement" json:"id"`
	NoInvoice   string             `gorm:"type:varchar(50);not null" json:"no_invoice" validate:"required"`
	IDSeller    int                `gorm:"not null" json:"id_seller" validate:"required"`
	IDBuyer     int                `gorm:"not null" json:"id_buyer" validate:"required"`
	TotalPrice  float64            `gorm:"type:decimal(10,2);not null" json:"total_price" validate:"required"`
	SalesDate   time.Time          `gorm:"type:datetime;not null" json:"sales_date" validate:"required"`
	Status      string             `gorm:"type:enum('open', 'pending', 'done');not null" json:"status" validate:"required,oneof=open pending done"`
	SalesDetail []SalesDetailDatum `gorm:"foreignKey:IDSales;references:ID" json:"products"`
}

type SalesDetailDatum struct {
	SalesDetail
	IDSales uint `gorm:"-" json:"id_sales,omitempty"`
}

func (Sales) TableName() string {
	return "sales"
}
