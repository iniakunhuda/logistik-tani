package model

import "time"

type ProductionDetail struct {
	ID             uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	IDProduction   int       `gorm:"not null" json:"id_production" validate:"required"`
	IDProductOwner int       `gorm:"not null" json:"id_product_owner" validate:"required"`
	QtyUse         int       `gorm:"not null" json:"qty_use" validate:"required"`
	Note           string    `gorm:"type:text" json:"note"`
	Date           time.Time `gorm:"type:datetime;not null" json:"date" validate:"required"`

	Production Production `gorm:"foreignKey:IDProduction;references:ID" json:"production"`
}

func (ProductionDetail) TableName() string {
	return "production_detail"
}
