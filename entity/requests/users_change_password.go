package requests

type UserChangePassword struct {
	RecentPassword string `binding:"required,min=6" json:"recent_password"`
	NewPassword    string `binding:"required,min=6" json:"new_password"`
}
