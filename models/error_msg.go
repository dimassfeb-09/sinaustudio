package models

type ErrorMsg struct {
	Msg  error  `json:error`
	Code string `json:code`
}

func ToErrorMsg(Code string, Msg error) *ErrorMsg {
	return &ErrorMsg{
		Code: Code,
		Msg:  Msg,
	}
}
