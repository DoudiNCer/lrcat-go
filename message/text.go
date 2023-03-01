package message

import (
	"encoding/json"
	"strconv"
	"time"
)

// 普通文本消息
type Text struct {
	Timestamp string      `json:"timestamp"`
	Sign      string      `json:"sign"`
	Msg_type  string      `json:"msg_Type"`
	Content   interface{} `json:"content"`
}

func (text Text) GetMsgJson() (string, error) {
	if text.Msg_type == "" {
		text.Msg_type = TEXT_MESSAGE_TYPE
	}
	contentJson, err := json.Marshal(text.Content)
	if err != nil {
		return "", err
	}
	text.Content = string(contentJson)
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
