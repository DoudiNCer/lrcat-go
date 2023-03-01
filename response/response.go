package response

const (
	SIGN_ERR = 19021
)

// 发送消息响应
type Response struct {
	Code int64       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
