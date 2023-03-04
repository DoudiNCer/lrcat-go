package message

// 各种基本消息的 tag
const (
	// At消息
	Tag_At = "at"
	// 文本消息
	Tag_Text = "text"
	// 超链接消息
	tag_Link = "a"
	// 表情信息
	//Tag_Emo = "emotion"
)

// 富文本消息文本消息单元
type RichUText struct {
	Tag  string `json:"tag"`
	Text string `json:"text"`
}

// 获取富文本消息文本消息单元
func GetRichUText(text string) RichUText {
	return RichUText{Tag: Tag_Text, Text: text}
}

// 富文本消息超链接单元
type RichULink struct {
	Tag  string `json:"tag"`
	Text string `json:"text"`
	Href string `json:"href"`
}

// 获取富文本消息超链接单元
func GetRichULink(text, link string) RichULink {
	return RichULink{Tag: tag_Link, Text: text, Href: link}
}

/*// 富文本消息表情单元
// 可用表情参考 https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message-reaction/emojis-introduce
type RichUEmoji struct {
	Tag       string `json:"tag"`
	EmojiType string `json:"emoji_Type"`
}

// 获取富文本消息表情单元
func GetRichUEmoji(emojiType string) RichUEmoji {
	return RichUEmoji{Tag: Tag_Emo, EmojiType: emojiType}
}*/
