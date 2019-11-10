package defs

type Err struct {
	Error     string `json:"error"`
	ErrorCode string `json:"error_code"`
}

type ErroResponse struct {
	HttpSC int
	Error  Err
}

var (
	ErrorRequestBodyParseFailed = ErroResponse{HttpSC: 400, Error: Err{Error: "Request body is not correct", ErrorCode: "001"}}
	ErrorNotAuthUser            = ErroResponse{HttpSC: 401, Error: Err{Error: "Uer authentication failed", ErrorCode: "002"}}
	ErrorDBError                = ErroResponse{HttpSC: 500, Error: Err{Error: "db ops failed", ErrorCode: "003"}}
	ErrorInternalFaults         = ErroResponse{HttpSC: 500, Error: Err{Error: "internal service error", ErrorCode: "004"}}
)
