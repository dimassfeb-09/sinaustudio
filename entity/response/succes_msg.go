package response

type SuccessResponse struct {
	Success    bool   `json:"success"`
	StatusCode int    `json:"status_code"`
	Msg        string `json:"message"`
	Data       any    `json:"data"`
}
