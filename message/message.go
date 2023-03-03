package message

const (
	TEXT_MESSAGE_TYPE         = "text"
	RICH_TEXT_MESSAGE_TYPE    = "post"
	SHARE_CHAT_MESSAGE_TYPE   = "share_chat"
	IMAGE_MESSAGE_TYPE        = "image"
	MESSACE_CARD_MESSAGE_TYPE = "interactive"
)

// Message 消息体
type Message interface {
	// GetMsgJson 获取用于发送的消息体
	GetMsgJson() (string, error)
	// SignMsg 签名消息
	SignMsg(key string)
}
