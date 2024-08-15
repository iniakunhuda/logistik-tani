package response

import "github.com/iniakunhuda/logistik-tani/purchase/model"

type PurchaseResponse struct {
	model.Purchase
	IDPenjual     uint   `json:"id_penjual"`
	NamaPenjual   string `json:"nama_penjual"`
	AlamatPenjual string `json:"alamat_penjual"`
	TelpPenjual   string `json:"telp_penjual"`
	// PenjualDetail UserResponse `json:"penjual_detail"`
	PembeliDetail UserResponse `json:"pembeli_detail"`
}
