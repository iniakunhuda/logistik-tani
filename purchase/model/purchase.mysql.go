package model

import (
	"time"
)

type Purchase struct {
	ID             uint                  `gorm:"primaryKey;autoIncrement" json:"id"`
	NoInvoice      string                `gorm:"type:varchar(255);not null" json:"no_invoice" validate:"required"`
	IDSeller       int                   `gorm:"type:int;" json:"id_seller"`
	SellerName     string                `gorm:"type:varchar(255);not null" json:"seller_name"`
	SellerAddress  string                `gorm:"type:varchar(255);not null" json:"seller_address"`
	SellerTelp     string                `gorm:"type:varchar(255);not null" json:"seller_telp"`
	IDBuyer        int                   `gorm:"type:int;not null" json:"id_buyer" validate:"required"`
	TotalPrice     float64                   `gorm:"type:decimal(10,2);not null" json:"total_price" validate:"required"`
	PurchaseDate   time.Time             `gorm:"type:date;not null" json:"purchase_date" validate:"required"`
	Status         string                `gorm:"type:enum('open','pending','done', 'cancel');not null" json:"status" validate:"required"`
	PurchaseDetail []PurchaseDetailDatum `gorm:"foreignKey:IDPurchase;references:ID" json:"products"`
}

type PurchaseDetailDatum struct {
	PurchaseDetail
	IDPurchase uint `gorm:"-" json:"id_purchase,omitempty"`
}

func (Purchase) TableName() string {
	return "purchase"
}
