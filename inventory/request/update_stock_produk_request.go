package request

type UpdateStockProdukRequest struct {
	IDProduk    string `json:"id_produk"`
	StokTerbaru string `json:"stok_terbaru"`
}
