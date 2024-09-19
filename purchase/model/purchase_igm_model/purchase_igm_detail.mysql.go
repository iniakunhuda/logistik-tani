package purchaseigmmodel

type PurchaseIgmDetail struct {
	ID            uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	IDPurchaseIgm int     `gorm:"not null" json:"id_purchase_igm" validate:"required"`
	IDUser        int     `gorm:"not null" json:"id_user" validate:"required"`
	IDUserLand    int     `gorm:"not null" json:"id_user_land" validate:"required"`
	IDProduction  int     `gorm:"not null" json:"id_production" validate:"required"`
	TotalKg       float64 `gorm:"type:decimal(10,2);not null" json:"total_kg" validate:"required"`
	HargaKg       float64 `gorm:"type:decimal(10,2);not null" json:"harga_kg" validate:"required"`
	Subtotal      float64 `gorm:"type:decimal(10,2);not null" json:"subtotal" validate:"required"`

	PurchaseIgm PurchaseIgm `gorm:"foreignKey:IDPurchaseIgm;references:ID" json:"purchase_igm"`
}

func (PurchaseIgmDetail) TableName() string {
	return "purchase_igm_detail"
}
