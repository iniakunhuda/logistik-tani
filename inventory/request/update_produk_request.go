package request

type UpdateProdukRequest struct {
	IDUser     uint   `json:"id_user"`
	NamaProduk string `json:"nama_produk"`
	Hpp        uint   `json:"hpp"`
	HargaJual  uint   `json:"harga_jual"`
	Kategori   string `json:"kategori"`
	Jenis      string `json:"jenis"`
	StokAktif  uint   `json:"stok_aktif"`
	Varietas   string `json:"varietas"`
	Status     string `json:"status"`
}
