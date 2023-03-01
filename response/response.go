package response

const (
	// 一切正常
	SUCCESS = 0
	// 请求体格式错误
	BAD_REQUEST = 9499
	// 签名校验失败
	SIGN_ERR = 19021
	//关键词校验错误
	KEY_WORD_ERR = 19024
	// IP白名单限制
	IP_ADDR_ERR = 19022
)

// 发送消息响应
type Response struct {
	Code int64       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
