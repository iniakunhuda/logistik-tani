package request

type CreateProdukRequest struct {
	ID         uint   `json:"id"`
	IDUser     uint   `json:"id_user" validate:"required"`
	NamaProduk string `json:"nama_produk" validate:"required"`
	Hpp        uint   `json:"hpp"`
	HargaJual  uint   `json:"harga_jual" validate:"required"`
	Kategori   string `json:"kategori" validate:"required"`
	Jenis      string `json:"jenis" validate:"required"`
	StokAktif  uint   `json:"stok_aktif" validate:"required"`
	Varietas   string `json:"varietas" validate:"required"`
	Status     string `json:"status" validate:"required"`
}
