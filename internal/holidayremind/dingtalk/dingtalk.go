package dingtalk

import (
	"holidayRemind/internal/pkg"
	"holidayRemind/internal/pkg/net"
	"log"
)

// SendMdMessage 发送钉钉机器人Markdown消息
func SendMdMessage(dto MessageDTO) error {
	var err error
	var title = dto.Title
	var At AtParams
	if dto.IsRemind || dto.IsRemindAll {
		title = title + "@" + dto.Tel
		At = AtParams{
			IsAtAll:   dto.IsRemindAll,
			AtMobiles: []string{dto.Tel},
		}
	}
	body := MarkdownMessage{ // 构建 post 消息体
		MsgType: Markdown,
		Markdown: MarkdownParams{
			Title: title,
			Text:  dto.Text,
		},
		At: At,
	}

	// 输出拼接好的字符串
	messageJson := ""
	err = pkg.ToJson(body, &messageJson)
	if err != nil {
		return err
	}
	log.Printf("DingTalk Send Message Json:%s", messageJson)
	var result []byte
	requestData := net.RequestBaseData{
		Url: "https://oapi.dingtalk.com/robot/send",
		Params: map[string]string{
			"access_token": dto.Token,
		},
		Headers: nil,
	}
	client := net.GetSimpleHttpClient()
	err = client.PostByJson(requestData, body, &result)
	if err != nil {
		return err
	}
	log.Printf("DingTalk Response Result:%s", string(result))
	return nil
}
