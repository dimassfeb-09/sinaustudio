package response

type ErrorMsg struct {
	Success    bool   `json:"success"`
	StatusCode int    `json:"status_code"`
	ErrorKey   string `json:"error_key"`
	Msg        any    `json:"message"`
}
