# lrcat-go
A smart cat that helps you send messages with Feishu bot
## 安装
```bash
go get github.com/DoudiNCer/lrcat_go
```
## 使用
### 初始化机器人
```go
myBot := bot.NewBot("https://open.feishu.cn/open-apis/bot/v2/hook/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx")
```
### 全局开关
使用`Disable()`关闭机器人而不需要修改代码；使用`ReEnable()`重新开启：
```go
// 关闭机器人
bot.Disable()
// 重新打开机器人
bot.ReEnable()
```
> 为了提高效率，未对当前状态进行验证。
### 创建普通文本消息
```go
// 创建文本消息
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
### 创建富文本消息
```go
// 创建富文本消息
var msg message.RichText
// 创建“一篇”富文本消息
rtcp := message.InitRichTextContentPart("你好，猫猫")
// 段落段落切片
var plist []interface{}
// 添加@消息
plist = append(plist, user.GetAtUser("xxxxxxxxxxxxxxxxxxx", "猫猫学姐").RichTextMsgUserObject())
// 添加链接消息
plist = append(plist, message.GetRichULink("\n项目链接", "https://github.com/DoudiNCer/lrcat-go"))
// 将段落切片填充到消息中
rtcp.AddParam(plist)
// 指定语言版本
msg.AddContent("zh-cn", rtcp)
// 发送消息
send, err := myBot.Send(&msg)
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