package model

import (
	"time"
)

type SalesDetail struct {
	ID         uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	IDSales    uint      `gorm:"type:bigint;not null" json:"id_sales"`
	IDProduk   uint       `gorm:"type:int;not null" json:"id_produk"`
	Jenis      string    `gorm:"type:enum('pupuk','bibit','obat');not null" json:"jenis" validate:"required"`
	Harga      int       `gorm:"type:int;not null" json:"harga" validate:"required"`
	Qty        int       `gorm:"type:int;not null" json:"qty" validate:"required"`
	TotalHarga int       `gorm:"type:int;not null" json:"total_harga" validate:"required"`
	Tanggal    time.Time `gorm:"type:date;not null" json:"tanggal" validate:"required"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (SalesDetail) TableName() string {
	return "sales_detail"
}
