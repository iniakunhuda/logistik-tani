package model

import (
	"time"
)

type SalesIgm struct {
	ID             uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	IDProductOwner int       `gorm:"type:int;not null" json:"id_product"`
	IDSeller       int       `gorm:"type:int;not null" json:"id_seller" validate:"required"`
	NoInvoice      string    `gorm:"type:varchar(100);not null" json:"no_invoice" validate:"required"`
	BuyerName      string    `gorm:"type:varchar(255);not null" json:"buyer_name" validate:"required"`
	BuyerAddress   string    `gorm:"type:varchar(255);not null" json:"buyer_address" validate:"required"`
	BuyerTelp      string    `gorm:"type:varchar(20);not null" json:"buyer_telp" validate:"required"`
	Price          float64   `gorm:"type:ype:decimal(10,2);not null" json:"price" validate:"required"`
	Qty            int       `gorm:"type:ype:int;not null" json:"qty" validate:"required"`
	TotalPrice     float64   `gorm:"type:ype:decimal(10,2);not null" json:"total_price" validate:"required"`
	SalesDate      time.Time `gorm:"type:date;not null" json:"sales_date" validate:"required"`
	Status         string    `gorm:"type:enum('open','pending','done');not null" json:"status" validate:"required"`
	Note           string    `gorm:"type:text" json:"note"`
}

func (SalesIgm) TableName() string {
	return "sales_igm_sugar"
}
