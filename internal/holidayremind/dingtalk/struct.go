package dingtalk

// MarkdownMessage MarkDown消息
type MarkdownMessage struct {
	MsgType  messageType    `json:"msgtype"`  // 消息类型
	Markdown MarkdownParams `json:"markdown"` // Markdown消息参数
	At       AtParams       `json:"at"`       // 提醒参数
}

// MarkdownParams Markdown格式的消息参数
type MarkdownParams struct {
	Title string `json:"title"` //标题
	Text  string `json:"text"`  //内容
}

// AtParams 提醒参数
type AtParams struct {
	AtMobiles []string `json:"atMobiles,omitempty"` // 需要提醒的手机号数组
	IsAtAll   bool     `json:"isAtAll"`             // 是否提醒全部人
}

// MarkdownMessageDTO MarkDown消息DTO
type MarkdownMessageDTO struct {
	Title       string // 标题
	Text        string // 内容
	Token       string // 机器人Token
	Tel         string // 提醒人手机号
	IsRemind    bool   // 是否提醒
	IsRemindAll bool   // 是否全部提醒
}

type FeedCardMessage struct {
	MsgType  messageType    `json:"msgtype"` // 消息类型
	FeedCard FeedCardParams `json:"feedCard"`
}

type FeedCardParams struct {
	Links []FeedCardLink `json:"links"`
}

type FeedCardLink struct {
	Title      string `json:"title"`
	MessageUrl string `json:"messageURL"`
	PictureUrl string `json:"picURL"`
}

type FeedCardMessageDTO struct {
	Links []FeedCardLink
	Token string // 机器人Token
}
