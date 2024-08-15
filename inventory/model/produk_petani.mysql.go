package model

import (
	"time"
)

type ProdukPetani struct {
	ID         uint   `gorm:"primaryKey" json:"id"`
	IDUser     uint   `gorm:"type:int" json:"id_user"`
	NamaProduk string `gorm:"type:varchar(100);not null" json:"nama_produk" validate:"required"`
	Hpp        uint   `gorm:"default:0" json:"hpp"`
	HargaJual  uint   `gorm:"default:0" json:"harga_jual"`
	Kategori   string `gorm:"type:varchar(100);not null" json:"kategori" validate:"required"`
	Jenis      string `gorm:"type:varchar(100);not null" json:"jenis" validate:"required"`
	StokAktif  uint   `gorm:"default:0" json:"stok_aktif"`
	Varietas   string `gorm:"type:varchar(100);not null" json:"varietas" validate:"required"`
	Status     string `gorm:"type:varchar(100);not null" json:"status" validate:"required"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (ProdukPetani) TableName() string {
	return "inventory_petani"
}
