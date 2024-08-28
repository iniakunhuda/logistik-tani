package model

type Product struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	Name        string `gorm:"size:255;not null"`
	Category    string `gorm:"type:enum('pupuk','obat','bibit','gula','alat');not null"`
	Description string `gorm:"type:text"`
	PriceBuy    int    `gorm:"not null;default:0"`
	PriceSell   int    `gorm:"not null;default:0"`
}

func (Product) TableName() string {
	return "product"
}
