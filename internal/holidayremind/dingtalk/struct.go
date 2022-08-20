package dingtalk

// AtParams 提醒参数
type AtParams struct {
	// 需要提醒的手机号数组
	AtMobiles []string `json:"atMobiles,omitempty"`
	// 是否提醒全部人
	IsAtAll bool `json:"isAtAll"`
}

// MarkdownMessage MarkDown消息
type MarkdownMessage struct {
	// 消息类型
	MsgType string `json:"msgtype"`
	// Markdown消息参数
	Markdown MarkdownParams `json:"markdown"`
	// 提醒参数
	At AtParams `json:"at"`
}

type MarkdownParams struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

type Message struct {
	Title       string
	Text        string
	Token       string
	Tel         string
	IsRemind    bool
	IsRemindAll bool
}
