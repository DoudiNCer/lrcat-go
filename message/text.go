package message

import (
	"encoding/json"
	"github.com/DoudiNCer/lrcat-go/utils"
	"strconv"
	"time"
)

// TEXT_MESSAGE_CONTENT_TAG 普通消息消息体标签
const TEXT_MESSAGE_CONTENT_TAG = "text"

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
		TEXT_MESSAGE_CONTENT_TAG: msg,
	}
}

// GetMsgJson 获取用于发送的请求体 JSON
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

// SignMsg 签名消息
func (text *Text) SignMsg(key string) error {
	timestamp := time.Now().Unix()
	sign, err := utils.GetSign(key, timestamp)
	if err != nil {
		return err
	}
	text.Sign = sign
	text.Timestamp = strconv.FormatInt(timestamp, 10)
	return err
}
