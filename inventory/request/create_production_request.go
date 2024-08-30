package request

import "time"

type CreateProductionRequest struct {
	ID           uint      `json:"id"`
	IDUser       int       `json:"id_user" validate:"required"`
	IDLand       int       `json:"id_land" validate:"required"`
	Title        string    `json:"title" validate:"required"`
	DateStart    time.Time `json:"date_start" validate:"required"`
	DateEnd      time.Time `json:"date_end"` // nullable field
	TotalPanenKg int       `json:"total_panen_kg"`
	Status       string    `json:"status" validate:"required,oneof=pending"`
}
