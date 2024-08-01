package model

import (
	"time"
)

type Sales struct {
	ID               uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	NoInvoice        string    `gorm:"type:varchar(100);not null" json:"no_invoice" validate:"required"`
	IDPenjual        int       `gorm:"type:int;not null" json:"id_penjual" validate:"required"`
	IDPembeli        int       `gorm:"type:int;not null" json:"id_pembeli" validate:"required"`
	TotalHarga       int       `gorm:"type:int;not null" json:"total_harga" validate:"required"`
	Tanggal          time.Time `gorm:"type:date;not null" json:"tanggal" validate:"required"`
	Status           string    `gorm:"type:enum('open','pending','done');not null" json:"status" validate:"required"`
	IsPurchasedByIGM bool      `gorm:"type:boolean;not null" json:"is_purchased_by_igm" validate:"required"`
	InvPurchasedIGM  *string   `gorm:"type:varchar(100)" json:"inv_purchased_igm"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (Sales) TableName() string {
	return "sale"
}
