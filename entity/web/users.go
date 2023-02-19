package web

type UserResponse struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	ClassID int    `json:"class_id"`
}
