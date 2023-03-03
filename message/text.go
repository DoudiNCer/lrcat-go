package message

import (
	"encoding/json"
	"strconv"
	"time"
)

// Text 普通文本消息
type Text struct {
	Timestamp string                 `json:"timestamp,omitempty"`
	Sign      string                 `json:"sign,omitempty"`
	MsgType   string                 `json:"msg_type"`
	Content   map[string]interface{} `json:"content"`
}

// AddContent 添加消息内容
func (text *Text) AddContent(msg string) {
	text.Content = map[string]interface{}{
		TEXT_MESSAGE_TYPE: msg,
	}
}

func (text Text) GetMsgJson() (string, error) {
	if text.MsgType == "" {
		text.MsgType = TEXT_MESSAGE_TYPE
	}
	textJson, err := json.Marshal(text)
	if err != nil {
		return "", err
	}
	return string(textJson), nil
}

func (text *Text) SignMsg(key string) {
	text.Sign = key
	text.Timestamp = strconv.FormatInt(time.Now().UnixNano(), 10)
}
