package dingtalk

import (
	"fmt"
	"holidayRemind/internal/pkg"
	"holidayRemind/internal/pkg/net"
)

// SendMdMessage 发送钉钉机器人Markdown消息
func SendMdMessage(msg Message) error {
	var title = msg.Title
	var At AtParams
	if msg.IsRemind {
		title = title + "@" + msg.Tel
		At = AtParams{
			IsAtAll:   msg.IsRemindAll,
			AtMobiles: []string{msg.Tel},
		}
	}
	message := MarkdownMessage{ // 构建 post 消息体
		MsgType: MsgTypeMarkdown,
		Markdown: MarkdownParams{
			Title: title,
			Text:  msg.Text,
		},
		At: At,
	}

	// 输出拼接好的字符串
	println(pkg.MapToJson(message))
	paramMap := map[string]string{
		"access_token": msg.Token,
	}
	result := ""
	net.Post("https://oapi.dingtalk.com/robot/send", message, paramMap, net.Json, &result)
	fmt.Printf("DingTalk Response Result:%s", result)
	return nil
}
