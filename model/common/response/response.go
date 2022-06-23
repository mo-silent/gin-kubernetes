package response

type CommonResponse struct {
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
