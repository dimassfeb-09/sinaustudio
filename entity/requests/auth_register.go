package requests

type AuthRegisterRequest struct {
	Name     string `binding:"required,min=5" json:"name"`
	Email    string `binding:"required,email,min=5" json:"email"`
	Password string `binding:"required,min=6" json:"password"`
	Role     string `binding:"required,min=5" json:"role"`
	NPM      string `binding:"required_if=Role mahasiswa" json:"npm"`
	ClassID  int    `binding:"required,numeric" json:"class_id"`
}
