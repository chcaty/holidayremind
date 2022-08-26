package dingtalk

import (
	"holidayRemind/internal/pkg/uxjson"
	"holidayRemind/internal/pkg/uxnet"
	"log"
)

// SendMdMessage 发送钉钉机器人Markdown消息
func SendMdMessage(dto MarkdownMessageDTO) error {
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
		MsgType: markdown,
		Markdown: MarkdownParams{
			Title: title,
			Text:  dto.Text,
		},
		At: At,
	}

	// 输出拼接好的字符串
	messageJson := ""
	err = uxjson.ToJson(body, &messageJson)
	if err != nil {
		return err
	}
	log.Printf("DingTalk Send MarkDownMessage Json:%s", messageJson)
	err = postMessage(dto.Token, body)
	if err != nil {
		return err
	}
	return nil
}

func SendFeedCardMessage(dto FeedCardMessageDTO) error {
	var err error
	var links []FeedCardLink
	for _, linkInfo := range dto.Links {
		links = append(links, FeedCardLink{
			Title:      linkInfo.Title,
			MessageUrl: linkInfo.MessageUrl,
			PictureUrl: linkInfo.PictureUrl,
		})
	}
	body := FeedCardMessage{
		MsgType: feedCard,
		FeedCard: FeedCardParams{
			Links: links,
		},
	}
	// 输出拼接好的字符串
	messageJson := ""
	err = uxjson.ToJson(body, &messageJson)
	if err != nil {
		return err
	}
	log.Printf("DingTalk Send FeedCardMessage Json:%s", messageJson)
	err = postMessage(dto.Token, body)
	if err != nil {
		return err
	}
	return nil
}

func postMessage(token string, body any) error {
	var err error
	var result []byte
	requestData := uxnet.BaseData{
		Url: "https://oapi.dingtalk.com/robot/send",
		Params: map[string]string{
			"access_token": token,
		},
		Headers: nil,
	}
	client := uxnet.GetHttpClient()
	err = client.PostByJson(requestData, body, &result)
	if err != nil {
		return err
	}
	log.Printf("DingTalk Response Result:%s", string(result))
	return nil
}
