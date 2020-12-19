package response

type Response struct {
	Code string      `json:"response_code"`
	Desc string      `json:"response_desc"`
	Data interface{} `json:"response_data"`
}

func New() *Response {
	return &Response{
		Code: "XX",
		Desc: "General Error",
		Data: new(struct{}),
	}
}
