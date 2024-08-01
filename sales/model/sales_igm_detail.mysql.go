package model

import (
	"time"
)

type SalesIgmDetail struct {
	ID          uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	IDSalesIGM  uint   `gorm:"type:bigint;not null" json:"id_sales_igm"`
	IDProdukIGM int    `gorm:"type:int;default:1;not null" json:"id_produk_igm" validate:"required"`
	HargaBeli   int    `gorm:"type:int;not null" json:"harga_beli" validate:"required"`
	HargaJual   int    `gorm:"type:int;not null" json:"harga_jual" validate:"required"`
	Qty         int    `gorm:"type:int;not null" json:"qty" validate:"required"`
	TotalHarga  int    `gorm:"type:int;not null" json:"total_harga" validate:"required"`
	Catatan     string `gorm:"type:text" json:"catatan"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (SalesIgmDetail) TableName() string {
	return "sale"
}
