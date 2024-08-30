package request

import "time"

type CreateProductionDetailRequest struct {
	IDProduction int                                    `json:"id_production" validate:"required"`
	Products     []CreateProductionDetailProductRequest `json:"products" validate:"required"`
	Date         time.Time                              `json:"date" validate:"required"`
}

type CreateProductionDetailProductRequest struct {
	IDProduct int    `json:"id_product" validate:"required"`
	Qty       int    `json:"qty" validate:"required"`
	Note      string `json:"note"`
}
