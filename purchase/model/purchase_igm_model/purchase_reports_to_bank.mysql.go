package purchaseigmmodel

import "time"

type PurchaseReportsToBank struct {
	ID        uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	DateStart *time.Time `gorm:"type:date" json:"date_start"`
	DateEnd   *time.Time `gorm:"type:date" json:"date_end"`
	NoReport  string     `gorm:"type:varchar(255);not null" json:"no_report" validate:"required"`
	Note      string     `gorm:"type:text" json:"note"`
	Status    string     `gorm:"type:enum('open', 'pending', 'done');not null" json:"status" validate:"required,oneof=open pending done"`

	Details []PurchaseReportsToBankDetail `gorm:"foreignKey:IDPurchaseReportsToBank;references:ID" json:"details"`
}

func (PurchaseReportsToBank) TableName() string {
	return "purchase_reports_to_bank"
}
