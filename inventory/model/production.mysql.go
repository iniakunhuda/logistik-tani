package model

import "time"

type Production struct {
	ID           uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	IDUser       int        `gorm:"not null" json:"id_user" validate:"required"`
	IDUserLand   int        `gorm:"not null" json:"id_user_land" validate:"required"`
	Title        string     `gorm:"type:varchar(255);not null" json:"title" validate:"required"`
	DateMonth    int        `gorm:"not null" json:"date_month" validate:"required"`
	DateYear     int        `gorm:"not null" json:"date_year" validate:"required"`
	DateStart    time.Time  `gorm:"type:date" json:"date_start"`
	DateEnd      *time.Time `gorm:"type:date" json:"date_end"`
	TotalPanenKg int        `gorm:"default:0" json:"total_panen_kg"`
	Status       string     `gorm:"type:enum('pending', 'done');not null" json:"status" validate:"required,oneof=pending done"`

	Histories []ProductionDetailDatum `gorm:"foreignKey:IDProduction;references:ID" json:"histories"`
}

type ProductionDetailDatum struct {
	ProductionDetail
	Production Production `gorm:"-" json:"production,omitempty"`
}

func (Production) TableName() string {
	return "production"
}
