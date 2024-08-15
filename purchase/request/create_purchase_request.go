package request

import "time"

type CreatePurchaseRequest struct {
	IDPenjual     uint                        `json:"id_penjual"`
	NamaPenjual   string                      `json:"nama_penjual"`
	AlamatPenjual string                      `json:"alamat_penjual"`
	TelpPenjual   string                      `json:"telp_penjual"`
	IDPembeli     uint                        `json:"id_pembeli" validate:"required"` // ID petani
	Status        string                      `json:"status" validate:"required"`
	Produk        []CreatePurchaseItemRequest `json:"produk" validate:"required"`
	Tanggal       time.Time                   `json:"tanggal" validate:"required"`
}

type CreatePurchaseItemRequest struct {
	IDProduk uint   `json:"id_produk" validate:"required"`
	Jenis    string `json:"jenis" validate:"required"`
	Harga    uint   `json:"harga" validate:"required"`
	Qty      uint   `json:"qty" validate:"required"`
	Catatan  string `json:"catatan"`
}
