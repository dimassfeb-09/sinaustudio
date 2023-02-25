package requests

type UpdateRoomRequest struct {
	ID        int    `json:"id"`
	Name      string `binding:"required" json:"name"`
	URL       string `binding:"required" json:"url"`
	LectureID int    `binding:"required" json:"lecture_id"`
	StartRoom string `binding:"required" json:"start_room"`
	EndRoom   string `binding:"required" json:"end_room"`
}
