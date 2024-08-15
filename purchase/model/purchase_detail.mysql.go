package model

import (
	"time"
)

type PurchaseDetail struct {
	ID         uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	IDPurchase    uint   `gorm:"type:bigint;not null" json:"id_purchase"`
	IDProduk   uint   `gorm:"type:int;not null" json:"id_produk"`
	Jenis      string `gorm:"type:enum('pupuk','bibit','obat', 'alat');not null" json:"jenis" validate:"required"`
	Harga      int    `gorm:"type:int;not null" json:"harga" validate:"required"`
	Qty        int    `gorm:"type:int;not null" json:"qty" validate:"required"`
	TotalHarga int    `gorm:"type:int;not null" json:"total_harga" validate:"required"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (PurchaseDetail) TableName() string {
	return "purchase_detail"
}
