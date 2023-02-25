package response

type RoomResponse struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	URL       string `json:"url"`
	LectureID int    `json:"lecture_id"`
	StartRoom string `json:"start_room"`
	EndRoom   string `json:"end_room"`
}
