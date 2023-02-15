package helpers

import (
	"github.com/dimassfeb-09/sinaustudio.git/entity/response"
)

func ToErrorMsg(StatusCode int, Code string, Msg any) *response.ErrorMsg {
	return &response.ErrorMsg{
		Success:    false,
		StatusCode: StatusCode,
		ErrorKey:   Code,
		Msg:        Msg,
	}
}
