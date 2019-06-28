package object

type ResponseError struct {
	ErrorCode float64        `json:"error_code"`
	Error     string         `json:"error_msg"`
	Params    []RequestParam `json:"request_params"`
}
