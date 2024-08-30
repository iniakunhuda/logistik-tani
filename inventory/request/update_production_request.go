package request

import "time"

type UpdateProductionRequest struct {
	Title        string    `json:"title"`
	DateStart    time.Time `json:"date_start"`
	DateEnd      time.Time `json:"date_end"`
	TotalPanenKg int       `json:"total_panen_kg"`
	Status       string    `json:"status" validate:"oneof=pending, done"`
}
