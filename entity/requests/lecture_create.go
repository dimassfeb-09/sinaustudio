package requests

type InsertLectureRequest struct {
	ID     int    `json:"id"`
	Name   string `binding:"required" json:"name"`
	UserID int    `binding:"required" json:"user_id"`
}
