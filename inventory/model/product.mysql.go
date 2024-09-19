package model

type Product struct {
	ID          uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string `gorm:"size:255;not null" json:"name"`
	Category    string `gorm:"type:enum('pupuk','obat','bibit','gula','alat');not null" json:"category"`
	Description string `gorm:"type:text" json:"description"`
	PriceBuy    int    `gorm:"not null;default:0" json:"price_buy"`
	PriceSell   int    `gorm:"not null;default:0" json:"price_sell"`
}

func (Product) TableName() string {
	return "product"
}
