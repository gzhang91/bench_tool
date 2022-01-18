package bench

type BaseService interface {
	DoGet(v interface{}) error
	DoPost(v interface{}) error
	GetServiceType() OperatorType
}

type RequestParam struct {
	Uri     string            `json:"uri"`
	Request string            `json:"request"`
	Headers map[string]string `json:"headers"`
	Body    string            `json:"body"`
}
