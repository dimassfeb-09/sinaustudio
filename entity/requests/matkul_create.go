package requests

type InsertMatkulRequest struct {
	KodeMatkul string `binding:"required" json:"kode_matkul"`
	Name       string `binding:"required" json:"name"`
}
