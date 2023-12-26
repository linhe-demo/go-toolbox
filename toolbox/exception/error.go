package exception

const (
	ParamCode    = 1001
	ApiCode      = 2001
	RocketMqCode = 3001
)

var ErrorMsgMap = map[int]string{
	ParamCode:    "参数异常",
	ApiCode:      "功能不支持",
	RocketMqCode: "Rocket-Mq 异常",
}

type CodeError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type CodeErrorResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func NewCodeError(code int, msg string) error {
	return &CodeError{Code: code, Msg: msg}
}

func NewError(code int, msg string) error {
	return NewCodeError(code, msg)
}

func (e *CodeError) Error() string {
	return e.Msg
}

func (e *CodeError) Data() *CodeErrorResponse {
	return &CodeErrorResponse{
		Code: e.Code,
		Msg:  ErrorMsgMap[e.Code],
	}
}
