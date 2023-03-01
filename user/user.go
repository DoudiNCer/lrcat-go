package user

// 用于构建可以@的用户
type User struct {
	Tag string `json:"tag"`
	// 用户ID，使用Open ID
	UserId string `json:"user_id"`
	// 展示出来的用户名
	UserName string
}

// 获取可插入到文本消息的代码
func (user User) TextMsgUserJson() string {
	return "<at user_id=\"" + user.UserId + "\">" + user.UserName + "</at>"
}

// 获取可插入到富文本消息的结构体
func (user User) RichTextMsgUserObject() interface{} {
	user.Tag = "at"
	return user
}

func GetAtAllUser() (user User) {
	var u = User{UserId: "all", UserName: "所有人"}
	return u
}
