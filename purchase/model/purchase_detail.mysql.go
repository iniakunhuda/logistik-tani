package model

type PurchaseDetail struct {
	ID             uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	IDPurchase     int     `gorm:"type:bigint;not null" json:"id_purchase"`
	IDProductOwner int     `gorm:"type:int;not null" json:"id_product"`
	Price          float64 `gorm:"type:decimal(10,2);not null" json:"price" validate:"required"`
	Qty            int     `gorm:"type:int;not null" json:"qty" validate:"required"`
	Subtotal       float64 `gorm:"type:decimal(10,2);not null" json:"subtotal" validate:"required"`
	Name           string  `gorm:"-" json:"name"`
	Description    string  `gorm:"-" json:"description"`
}

func (PurchaseDetail) TableName() string {
	return "purchase_detail"
}
