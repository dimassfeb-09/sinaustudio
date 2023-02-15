package web

type UserResponse struct {
	UUID    string `json:"uuid"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	ClassID int    `json:"class_id"`
}
