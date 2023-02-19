package requests

type UpdateLectureRequest struct {
	ID   int    `json:"id"`
	Name string `binding:"required" json:"name"`
}
