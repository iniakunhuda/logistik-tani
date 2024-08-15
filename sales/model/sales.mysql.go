package model

import (
	"time"
)

type Sales struct {
	ID               uint               `gorm:"primaryKey;autoIncrement" json:"id"`
	NoInvoice        string             `gorm:"type:varchar(100);not null" json:"no_invoice" validate:"required"`
	IDPenjual        uint               `gorm:"type:uint;not null" json:"id_penjual" validate:"required"`
	IDPembeli        uint               `gorm:"type:uint;not null" json:"id_pembeli" validate:"required"`
	TotalHarga       int                `gorm:"type:int;not null" json:"total_harga" validate:"required"`
	Tanggal          time.Time          `gorm:"type:date;not null" json:"tanggal" validate:"required"`
	Status           string             `gorm:"type:enum('open','pending','done');not null" json:"status" validate:"required"`
	IsPurchasedByIGM bool               `gorm:"type:boolean;not null" json:"is_purchased_by_igm" validate:"required"`
	SalesDetail      []SalesDetailDatum `gorm:"foreignKey:IDSales;references:ID" json:"products"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type SalesDetailDatum struct {
	SalesDetail
	IDSales uint `gorm:"-" json:"id_sales,omitempty"`
}

func (Sales) TableName() string {
	return "sales"
}
