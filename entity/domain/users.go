package domain

type Users struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
	NPM      string `json:"npm"`
	ClassID  int    `json:"class_id"`
}
