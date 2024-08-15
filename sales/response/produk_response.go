package response

import "time"

type ProdukResponse struct {
	ID         uint   `json:"id"`
	IDUser     uint   `json:"id_user"`
	NamaProduk string `json:"nama_produk" validate:"required"`
	Hpp        uint   `json:"hpp"`
	HargaJual  uint   `json:"harga_jual"`
	Kategori   string `json:"kategori" validate:"required"`
	Jenis      string `json:"jenis" validate:"required"`
	StokAktif  uint   `json:"stok_aktif"`
	Varietas   string `json:"varietas" validate:"required"`
	Status     string `json:"status" validate:"required"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
