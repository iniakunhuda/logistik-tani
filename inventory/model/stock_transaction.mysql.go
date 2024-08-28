package model

type StockTransaction struct {
	ID          uint    `gorm:"primaryKey;autoIncrement"`
	UserID      int     `gorm:"not null"`
	ProductID   int     `gorm:"not null"`
	Stock       int     `gorm:"not null"`
	PriceBuy    int     `gorm:"not null;default:0"`
	PriceSell   int     `gorm:"not null;default:0"`
	Description string  `gorm:"type:text"`
	Product     Product `gorm:"foreignKey:ProductID;references:ID"`
}

func (StockTransaction) TableName() string {
	return "stock_transaction"
}
