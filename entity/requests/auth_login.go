package requests

type AuthLoginRequest struct {
	Email    string `binding:"required,email,min=5" json:"email"`
	Password string `binding:"required,alphanum,min=6" json:"password"`
}
