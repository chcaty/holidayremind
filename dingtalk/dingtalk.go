package dingtalk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"holidayRemind/common"
	"net/http"
)

func SendMessage(msg Message) (int, error) {
	err := sendMdMessage(msg)
	if err != nil {
		fmt.Printf("SendMdMessage to DingTalk error: %s", err.Error())
		return 500, err
	}
	fmt.Printf("SendMdMessage to DingTalk success: %s", msg.Text)
	return 200, nil
}

func sendMdMessage(msg Message) error {
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
	println(common.MapToJson(message))
	var payloadBytes, err = json.Marshal(message)
	if err != nil {
		return err
	}
	body := bytes.NewReader(payloadBytes)

	dingReq, err := http.NewRequest("POST",
		"https://oapi.dingtalk.com/robot/send", body)
	if err != nil {
		return err
	}
	dingReq.Header.Set("Content-Type", "application/json")

	params := dingReq.URL.Query()
	params.Add("access_token", msg.Token)
	dingReq.URL.RawQuery = params.Encode()

	dingResp, err := http.DefaultClient.Do(dingReq) // 发送请求到钉钉
	if err != nil {
		return err
	}
	if dingResp != nil {
		defer common.CloseBody(dingResp.Body)
	}
	return nil
}
