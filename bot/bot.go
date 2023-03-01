package bot

import (
	"encoding/json"
	"errors"
	"github.com/DodiNCer/lrcat-go/message"
	"github.com/DodiNCer/lrcat-go/response"
	"net/http"
	"strconv"
	"strings"
)

// Bot 机器人
type Bot struct {
	webhook string
}

// NewBot 创建机器人
func (Bot) NewBot(webhook string) (bot Bot) {
	var nb Bot
	nb.webhook = webhook
	return nb
}

// Send 发送消息
func (bot Bot) Send(msg message.Message) (*response.Response, error) {
	if bot.webhook == "" {
		return nil, errors.New("bot not initialized")
	}
	url := bot.webhook
	post, err := http.Post(url, "application/json", strings.NewReader(msg.GetMsgJson()))
	if err != nil {
		return nil, err
	}
	var body []byte
	_, err = post.Body.Read(body)
	if err != nil {
		return nil, err
	}
	defer post.Body.Close()
	if post.Status != strconv.Itoa(http.StatusOK) {
		return nil, errors.New("POST status error: " + string(body))
	}
	var resp response.Response
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, err
}
