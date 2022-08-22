package dingtalk

type MessageType string

const (
	Text       MessageType = "text"
	Link       MessageType = "link"
	Markdown   MessageType = "markdown"
	ActionCard MessageType = "actionCard"
)
