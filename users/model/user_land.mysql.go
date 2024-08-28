package model

type UserLand struct {
	ID          uint    `gorm:"primaryKey" json:"id"`
	IDUser      uint    `gorm:"type:uint;not null" json:"id_user"`
	LandName    string  `gorm:"type:varchar(255);not null" json:"land_name" validate:"required"`
	LandAddress string  `gorm:"type:text;not null" json:"land_address" validate:"required"`
	LandArea    float64 `gorm:"type:decimal(10,2);not null" json:"land_area" validate:"required"`
	TotalObat   float64 `gorm:"type:decimal(10,2);not null" json:"total_obat" validate:"required"`
	TotalPupuk  float64 `gorm:"type:decimal(10,2);not null" json:"total_pupuk" validate:"required"`
	TotalBibit  float64 `gorm:"type:decimal(10,2);not null" json:"total_bibit" validate:"required"`
	TotalTebu   float64 `gorm:"type:decimal(10,2);not null" json:"total_tebu" validate:"required"`

	User UserLandDatum `gorm:"foreignKey:IDUser;references:ID" json:"user"`
}

type UserLandDatum struct {
	User
	Lands        string `gorm:"-" json:"lands,omitempty"`
	Token        string `gorm:"-" json:"token,omitempty"`
	TokenExpired string `gorm:"-" json:"token_expired,omitempty"`
	Password     string `gorm:"-" json:"password,omitempty"`
}

func (UserLand) TableName() string {
	return "user_land"
}
