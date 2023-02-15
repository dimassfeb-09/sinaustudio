package domain

type Users struct {
	ID      int    `json:"id"`
	UUID    string `json:"uuid"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Role    string `json:"role"`
	NPM     string `json:"npm"`
	ClassID int    `json:"class_id"`
}
