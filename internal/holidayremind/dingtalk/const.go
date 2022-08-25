package dingtalk

type messageType string

const (
	text       messageType = "text"
	link       messageType = "link"
	markdown   messageType = "markdown"
	actionCard messageType = "actionCard"
	feedCard   messageType = "feedCard"
)
