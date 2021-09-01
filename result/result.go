package result

const (
	CODE_OK = 200
)
const (
	CODE_CLIENT_ERROR = iota + 400
)
const (
	CODE_INTERNAL_ERROR = iota + 500
)

type Result struct {
	Code   int         `json:"code"`
	Msg    string      `json:"name"`
	Result interface{} `json:"result"`
}

func OK_RESULT(obj interface{}) *Result {
	return &Result{
		Code:   CODE_OK,
		Msg:    "成功",
		Result: obj,
	}
}

func OK() *Result {
	return &Result{
		Code:   CODE_OK,
		Msg:    "成功",
		Result: nil,
	}
}

func ERROR(msg string) *Result {
	return &Result{
		Code:   CODE_CLIENT_ERROR,
		Msg:    msg,
		Result: nil,
	}
}
