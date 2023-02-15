package requests

type UserDeleteRequest struct {
	UUID string `binding:"required,alphanum" json:"uuid"`
}
