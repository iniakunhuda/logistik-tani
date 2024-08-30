package purchaseigm

import "time"

type PurchaseReportsToBank struct {
	ID        uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	DateStart *time.Time `gorm:"type:date" json:"date_start"`
	DateEnd   *time.Time `gorm:"type:date" json:"date_end"`
	Note      string     `gorm:"type:text" json:"note"`
	Status    string     `gorm:"type:varchar(50);not null" json:"status" validate:"required"`
}

func (PurchaseReportsToBank) TableName() string {
	return "purchase_reports_to_bank"
}
