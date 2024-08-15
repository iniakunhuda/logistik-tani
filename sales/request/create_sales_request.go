package request

import "time"

type CreateSalesRequest struct {
	IDPenjual uint                     `json:"id_penjual" validate:"required"` // ID pembibit
	IDPembeli uint                     `json:"id_pembeli" validate:"required"` // ID petani
	Status    string                   `json:"status" validate:"required"`
	Produk    []CreateSalesItemRequest `json:"produk" validate:"required"`
	Tanggal   time.Time                `json:"tanggal" validate:"required"`
}

type CreateSalesItemRequest struct {
	IDProduk uint      `json:"id_produk" validate:"required"`
	Jenis    string    `json:"jenis" validate:"required"`
	Harga    uint      `json:"harga" validate:"required"`
	Qty      uint      `json:"qty" validate:"required"`
	Tanggal  time.Time `json:"tanggal" validate:"required"`
	Catatan  string    `json:"catatan"`
}
