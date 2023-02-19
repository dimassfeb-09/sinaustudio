package domain

type AuthRegisterUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
	NPM      string `json:"npm"`
	ClassID  int    `json:"class_id"`
}

type AuthLoginUser struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
