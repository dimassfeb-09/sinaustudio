package requests

type InsertClassRequest struct {
	Name string `binding:"required,min=4,max=12" json:"name"`
}
