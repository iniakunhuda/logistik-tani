package purchaseigm

import "time"

type PurchaseIgm struct {
	ID           uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	NoInvoice    string    `gorm:"type:varchar(50)" json:"no_invoice"`
	PurchaseDate time.Time `gorm:"type:datetime;not null" json:"purchase_date" validate:"required"`
	Note         string    `gorm:"type:text" json:"note"`
	TotalTebu    float64   `gorm:"type:decimal(10,2);not null" json:"total_tebu" validate:"required"`
	TotalPrice   float64   `gorm:"type:decimal(10,2);not null" json:"total_price" validate:"required"`
	TotalFarmer  int       `gorm:"not null" json:"total_farmer" validate:"required"`
	Status       string    `gorm:"type:enum('open', 'pending', 'done');not null" json:"status" validate:"required,oneof=open pending done"`
}

func (PurchaseIgm) TableName() string {
	return "purchase_igm"
}
