package model

import "time"

type StockTransaction struct {
	ID             uint         `gorm:"primaryKey;autoIncrement"`
	IDProductOwner int          `gorm:"not null;index"`                          // id_product_owner
	IDUser         int          `gorm:"not null;index"`                          // id_user
	StockMovement  int          `gorm:"not null"`                                // stock_movement
	Date           time.Time    `gorm:"not null"`                                // date
	Description    string       `gorm:"type:varchar(255);null"`                  // description
	ProductOwner   ProductOwner `gorm:"foreignKey:IDProductOwner;references:ID"` // Foreign key association
}

func (StockTransaction) TableName() string {
	return "stock_transaction"
}
