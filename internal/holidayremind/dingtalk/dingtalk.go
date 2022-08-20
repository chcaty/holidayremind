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
	if msg.IsRemind || msg.IsRemindAll {
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
	messageJson := ""
	fmt.Printf("dingtalk send message:%s\n", pkg.ToJson(&messageJson, message))
	var result []byte
	requestData := net.RequestBaseData{
		Url: "https://oapi.dingtalk.com/robot/send",
		Params: map[string]string{
			"access_token": msg.Token,
		},
		Headers: nil,
	}
	client := net.GetSimpleHttpClient()
	err := client.PostByJson(&result, requestData, message)
	if err != nil {
		return err
	}
	fmt.Printf("DingTalk Response Result:%s\n", result)
	return nil
}
