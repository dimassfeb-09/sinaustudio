package requests

type UserUpdateRequest struct {
	UUID    string `binding:"required,alphanum" json:"uuid"`
	Name    string `binding:"required,alpha,min=5" json:"name"`
	Email   string `binding:"required,email,min=5" json:"email"`
	Role    string `binding:"required,alpha,min=5" json:"role"`
	NPM     string `binding:"required_if=Role mahasiswa" json:"npm"`
	ClassID int    `binding:"required,numeric" json:"class_id"`
}

type UserUpdateEmailRequest struct {
	UUID  string `binding:"required,alphanum" json:"uuid"`
	Email string `binding:"required,email,min=5" json:"email"`
}

type UserUpdateNameRequest struct {
	UUID string `binding:"required,alphanum" json:"uuid"`
	Name string `binding:"required,email,min=5" json:"name"`
}

type UserUpdateNPMRequest struct {
	UUID string `binding:"required,alphanum" json:"uuid"`
	NPM  string `binding:"required,email,min=5" json:"npm"`
}
