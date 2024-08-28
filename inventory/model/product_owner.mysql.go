package model

type ProductOwner struct {
	ID          uint    `gorm:"primaryKey;autoIncrement"`
	IDUser      int     `gorm:"not null"`
	IDProduct   int     `gorm:"not null"`
	Stock       int     `gorm:"not null"`
	PriceBuy    int     `gorm:"not null;default:0"`
	PriceSell   int     `gorm:"not null;default:0"`
	Description string  `gorm:"type:text"`
	Product     Product `gorm:"foreignKey:IDProduct;references:ID"`
}

func (ProductOwner) TableName() string {
	return "product_owner"
}
