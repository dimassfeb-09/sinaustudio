package requests

type InsertLectureRequest struct {
	ID   int    `json:"id"`
	Name string `binding:"require" json:"name"`
}
