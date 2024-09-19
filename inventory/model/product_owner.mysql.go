package model

type ProductOwner struct {
	ID          uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	IDUser      int     `gorm:"not null" json:"id_user"`
	IDProduct   int     `gorm:"not null" json:"id_product"`
	Stock       int     `gorm:"not null" json:"stock"`
	PriceBuy    int     `gorm:"not null;default:0" json:"price_buy"`
	PriceSell   int     `gorm:"not null;default:0" json:"price_sell"`
	Description string  `gorm:"type:text" json:"description"`
	Product     Product `gorm:"foreignKey:IDProduct;references:ID" json:"product"`
}

func (ProductOwner) TableName() string {
	return "product_owner"
}
