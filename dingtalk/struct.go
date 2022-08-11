package dingtalk

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

type Message struct {
	Title       string
	Text        string
	Token       string
	Tel         string
	IsRemind    bool
	IsRemindAll bool
}
