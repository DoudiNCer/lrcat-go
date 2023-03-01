package user

import "github.com/DodiNCer/lrcat-go/message"

// User 用于构建可以@的用户
type User struct {
	Tag string `json:"tag"`
	// 用户ID，使用Open ID
	UserId string `json:"user_id"`
	// 展示出来的用户名
	UserName string
}

// TextMsgUserJson 获取可插入到文本消息的代码
func (user User) TextMsgUserJson() string {
	return "<at user_id=\"" + user.UserId + "\">" + user.UserName + "</at>"
}

// RichTextMsgUserObject 获取可插入到富文本消息的结构体
func (user User) RichTextMsgUserObject() interface{} {
	user.Tag = message.TAG_AT
	return user
}

// GetAtAllUser 获取用于@全体的 User 对象
func GetAtAllUser() (user User) {
	var u = User{UserId: "all", UserName: "所有人"}
	return u
}

// GetAtUser 获取@某个用户的用户对象
func GetAtUser(userId, userName string) (user User) {
	var u = User{UserId: userId, UserName: userName}
	return u
}
