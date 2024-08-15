package model

import (
	"time"
)

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Name     string `gorm:"type:varchar(100);not null" json:"name" validate:"required"`
	Username string `gorm:"type:varchar(50);unique;not null" json:"username"`
	Email    string `gorm:"type:varchar(100);unique;not null" json:"email" validate:"required"`
	Password string `gorm:"type:varchar(255);not null" json:"password" validate:"min=5,max=20"`
	Alamat   string `gorm:"type:varchar(255);null" json:"alamat" validate:"required"`
	Telp     string `gorm:"type:varchar(15);null" json:"telp" validate:"required"`
	Role     string `gorm:"type:varchar(15);not null" json:"role" validate:"required"`

	Saldo       uint   `gorm:"default:0" json:"saldo"`
	// LastLogin   string `gorm:"null" json:"last_login"`
	AlamatKebun string `gorm:"type:varchar(255);null" json:"alamat_kebun"`
	TotalObat   uint   `gorm:"default:0" json:"total_obat"`
	TotalPupuk  uint   `gorm:"default:0" json:"total_pupuk"`
	TotalBibit  uint   `gorm:"default:0" json:"total_bibit"`
	TotalTebu   uint   `gorm:"default:0" json:"total_tebu"`
	LuasLahan   uint   `gorm:"default:0" json:"luas_lahan"`

	Token        *string    `gorm:"type:text;null" json:"token"`
	TokenExpired *time.Time `gorm:"null" json:"token_expired"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
