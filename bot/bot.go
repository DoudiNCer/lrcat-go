package bot

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/DodiNCer/lrcat-go/message"
	"github.com/DodiNCer/lrcat-go/response"
	"net/http"
	"strings"
)

// Bot 机器人
type Bot struct {
	enabled bool
	webhook string
}

// NewBot 创建机器人
func NewBot(webhook string) (bot Bot) {
	var nb Bot
	nb.enabled = true
	nb.webhook = webhook
	return nb
}

// Disable 禁用机器人
func (bot *Bot) Disable() {
	bot.enabled = false
}

// ReEnable 重新启用机器人
func (bot *Bot) ReEnable() {
	bot.enabled = true
}

// Send 发送消息
func (bot Bot) Send(msg message.Message) (*response.Response, error) {
	// 处理全局开关
	if !bot.enabled {
		return nil, nil
	}
	if bot.webhook == "" {
		return nil, errors.New("bot not initialized")
	}
	url := bot.webhook
	msgJson, err := msg.GetMsgJson()
	if err != nil {
		return nil, err
	}
	fmt.Println(msgJson)
	post, err := http.Post(url, "application/json", strings.NewReader(msgJson))
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	if post.StatusCode != http.StatusOK {
		return nil, errors.New("POST status error: " + post.Status + "\n")
	}
	var resp response.Response
	err = json.NewDecoder(post.Body).Decode(&resp)
	defer post.Body.Close()
	if err != nil {
		return nil, err
	}
	return &resp, err
}
