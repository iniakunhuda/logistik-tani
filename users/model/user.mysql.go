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

	Saldo        uint       `gorm:"default:0" json:"saldo"`
	Token        *string    `gorm:"type:text;null" json:"token"`
	TokenExpired *time.Time `gorm:"null" json:"token_expired"`
}

func (User) TableName() string {
	return "users"
}
