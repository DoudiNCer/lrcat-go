package message

import (
	"encoding/json"
	"github.com/DoudiNCer/lrcat-go/utils"
	"strconv"
	"time"
)

// RICH_TEXT_MESSAGE_CONTENT_TAG 富文本消息消息体标签
const RICH_TEXT_MESSAGE_CONTENT_TAG = "post"

// RichText 富文本消息
type RichText struct {
	Timestamp string                 `json:"timestamp,omitempty"`
	Sign      string                 `json:"sign,omitempty"`
	MsgType   string                 `json:"msg_type"`
	Content   map[string]interface{} `json:"content"`
}

// RichTextContentModel 富文本消息消息体模块
type RichTextContentModel struct {
	Title   string        `json:"title"`
	Content []interface{} `json:"content"`
}

// 初始化富文本消息消息体模块
func InitRichTextContentPart(title string) RichTextContentModel {
	var rtcm RichTextContentModel
	rtcm.Title = title
	return rtcm
}

// 为富文本消息消息体模块添加段落
func (rtcm *RichTextContentModel) AddParam(param []interface{}) {
	rtcm.Content = append(rtcm.Content, param)
}

// AddContent 将富文本消息消息体模块添加到消息体中
func (rt *RichText) AddContent(lang string, content RichTextContentModel) {
	if len(rt.Content) == 0 {
		cont := make(map[string]interface{})
		rt.Content = map[string]interface{}{
			RICH_TEXT_MESSAGE_CONTENT_TAG: cont,
		}
	}
	i := rt.Content[RICH_TEXT_MESSAGE_CONTENT_TAG]
	m := i.(map[string]interface{})
	m[lang] = content
	rt.Content[RICH_TEXT_MESSAGE_CONTENT_TAG] = m
}

// GetMsgJson 获取用于发送的请求体 JSON
func (rt RichText) GetMsgJson() (string, error) {
	if rt.MsgType == "" {
		rt.MsgType = RICH_TEXT_MESSAGE_TYPE
	}
	textJson, err := json.Marshal(rt)
	if err != nil {
		return "", err
	}
	return string(textJson), nil
}

// SignMsg 签名消息
func (rt *RichText) SignMsg(key string) error {
	timestamp := time.Now().Unix()
	sign, err := utils.GetSign(key, timestamp)
	if err != nil {
		return err
	}
	rt.Sign = sign
	rt.Timestamp = strconv.FormatInt(timestamp, 10)
	return err
}
