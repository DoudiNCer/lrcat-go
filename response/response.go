package response

const (
	SIGN_ERR = 19021
)

// 发送消息响应
type Response struct {
	code int64
	msg  string
	data interface{}
}
