package entity

type Result struct {
	Msg    string      `json:"msg"`
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

func NewResult(status int, msg string, data interface{}) *Result {
	r := new(Result)
	r.Status = status
	r.Msg = msg
	r.Data = data
	return r
}
