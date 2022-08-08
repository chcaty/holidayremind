package dingtalk

const (
	MsgTypeText       = "text"
	MsgTypeLink       = "link"
	MsgTypeMarkdown   = "markdown"
	MsgTypeActionCard = "actionCard"
)

type AtParams struct {
	AtMobiles []string `json:"atMobiles,omitempty"`
	IsAtAll   bool     `json:"isAtAll"`
}

type MarkdownMessage struct {
	MsgType  string         `json:"msgtype"`
	Markdown MarkdownParams `json:"markdown"`
	At       AtParams       `json:"at"`
}

type MarkdownParams struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}
