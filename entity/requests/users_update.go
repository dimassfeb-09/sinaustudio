package requests

type UserUpdateRequest struct {
	ID      int    `binding:"required,alphanum" json:"id"`
	Name    string `binding:"required,alpha,min=5" json:"name"`
	Email   string `binding:"required,email,min=5" json:"email"`
	Role    string `binding:"required,alpha,min=5" json:"role"`
	ClassID int    `binding:"required,numeric" json:"class_id"`
}
