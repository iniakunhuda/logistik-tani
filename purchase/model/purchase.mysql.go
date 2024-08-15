package model

import (
	"time"
)

type Purchase struct {
	ID             uint                  `gorm:"primaryKey;autoIncrement" json:"id"`
	NoInvoice      string                `gorm:"type:varchar(255);not null" json:"no_invoice" validate:"required"`
	IDPenjual      uint                  `gorm:"type:uint;not null" json:"id_penjual"`
	NamaPenjual    string                `gorm:"type:varchar(255);not null" json:"nama_penjual"`
	AlamatPenjual  string                `gorm:"type:varchar(255);not null" json:"alamat_penjual"`
	TelpPenjual    string                `gorm:"type:varchar(255);not null" json:"telp_penjual"`
	IDPembeli      uint                  `gorm:"type:uint;not null" json:"id_pembeli" validate:"required"`
	TotalHarga     int                   `gorm:"type:int;not null" json:"total_harga" validate:"required"`
	Tanggal        time.Time             `gorm:"type:date;not null" json:"tanggal" validate:"required"`
	Status         string                `gorm:"type:enum('open','pending','done', 'cancel');not null" json:"status" validate:"required"`
	PurchaseDetail []PurchaseDetailDatum `gorm:"foreignKey:IDPurchase;references:ID" json:"products"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type PurchaseDetailDatum struct {
	PurchaseDetail
	IDPurchase uint `gorm:"-" json:"id_purchase,omitempty"`
}

func (Purchase) TableName() string {
	return "purchase"
}
