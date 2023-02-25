package requests

type UpdateMatkulRequest struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	KodeMatkul string `json:"kode_matkul"`
}
