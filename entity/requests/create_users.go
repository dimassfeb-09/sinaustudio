package requests

type UserInsertRequest struct {
	UUID    string `binding:"required,alphanum" json:"uuid"`
	Name    string `binding:"required,alpha,min=5" json:"name"`
	Email   string `binding:"required,email,min=5" json:"email"`
	Role    string `binding:"required,alpha,min=5" json:"role"`
	NPM     string `binding:"required_if=Role mahasiswa" json:"npm"`
	ClassID int    `binding:"required,numeric" json:"class_id"`
}
