package requests

type UserDeleteRequest struct {
	ConfirmPassword string `binding:"required" json:"confirmpassword"`
}
