package message

import (
	"encoding/json"
	"strconv"
	"time"
)

// 群名片消息消息体标签
const SHARE_CHAT_MESSAGE_CONTENT_TAG = "share_chat_id"

// 群名片消息
type ShareChat struct {
	Timestamp string                 `json:"timestamp,omitempty"`
	Sign      string                 `json:"sign,omitempty"`
	MsgType   string                 `json:"msg_type"`
	Content   map[string]interface{} `json:"content"`
}

// 添加消息内容
func (sc *ShareChat) AddContent(msg string) {
	sc.Content = map[string]interface{}{
		SHARE_CHAT_MESSAGE_CONTENT_TAG: msg,
	}
}

func (sc ShareChat) GetMsgJson() (string, error) {
	if sc.MsgType == "" {
		sc.MsgType = SHARE_CHAT_MESSAGE_TYPE
	}
	textJson, err := json.Marshal(sc)
	if err != nil {
		return "", err
	}
	return string(textJson), nil
}

// 签名消息
func (sc *ShareChat) SignMsg(key string) {
	sc.Sign = key
	sc.Timestamp = strconv.FormatInt(time.Now().UnixNano(), 10)
}
