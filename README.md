# lrcat-go
A smart cat that helps you send messages with Feishu bot

## 使用
### 初始化机器人
```go
myBot := bot.NewBot("https://open.feishu.cn/open-apis/bot/v2/hook/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx")
```
### 创建普通文本消息
```go
// 创建变量
var msg message.Text
// 填充消息内容
msg.AddContent("你好，猫猫学姐")
// 发送消息
resp, err := myBot.Send(&msg)
if err != nil {
    log.Println(send)
    log.Println(err.Error())
}
```
## 高级功能
### At
首先构建用户对象：
```go
user := user.GetAtUser("xxxxxxxxxxxxxxxxxxx", "猫猫学姐")
```
若需要@所有人，构建方法如下：
```go
all := user.GetAtAllUser()
```
对于普通消息，使用`TextMsgUserJson()`方法获取插入到普通消息中的代码：
```go
	msg.AddContent("你好, " + user.TextMsgUserJson())
```