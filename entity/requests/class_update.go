package requests

type UpdateClassRequest struct {
	ID   int    `json:"id"`
	Name string `binding:"required,min=4" json:"name"`
}
