package purchaseigmmodel

type PurchaseReportsToBankDetail struct {
	ID                      uint `gorm:"primaryKey;autoIncrement" json:"id"`
	IDPurchaseReportsToBank int  `gorm:"not null" json:"id_purchase_reports_to_bank" validate:"required"`
	IDPurchaseIgm           int  `gorm:"not null" json:"id_purchase_igm" validate:"required"`

	PurchaseReportsToBank PurchaseReportsToBank `gorm:"foreignKey:IDPurchaseReportsToBank;references:ID" json:"purchase_reports_to_bank"`
	PurchaseIGM           PurchaseIgm           `gorm:"foreignKey:IDPurchaseIgm;references:ID" json:"purchase_igm"`
}

func (PurchaseReportsToBankDetail) TableName() string {
	return "purchase_reports_to_bank_detail"
}
