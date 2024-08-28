package model

import (
	"time"
)

type SalesIgm struct {
	ID            uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	NoInvoice     string    `gorm:"type:varchar(100);not null" json:"no_invoice" validate:"required"`
	IDSeller      int       `gorm:"type:int;not null" json:"id_seller" validate:"required"`
	NamaPembeli   string    `gorm:"type:varchar(255);not null" json:"nama_pembeli" validate:"required"`
	AlamatPembeli string    `gorm:"type:varchar(255);not null" json:"alamat_pembeli" validate:"required"`
	TelpPembeli   string    `gorm:"type:varchar(20);not null" json:"telp_pembeli" validate:"required"`
	TotalHarga    int       `gorm:"type:int;not null" json:"total_harga" validate:"required"`
	Catatan       string    `gorm:"type:text" json:"catatan"`
	Tanggal       time.Time `gorm:"type:date;not null" json:"tanggal" validate:"required"`
	Status        string    `gorm:"type:enum('open','progress','done');not null" json:"status" validate:"required"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (SalesIgm) TableName() string {
	return "sale"
}
