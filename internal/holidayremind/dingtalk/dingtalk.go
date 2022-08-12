package dingtalk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"holidayRemind/internal/pkg"
	"net/http"
)

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
	var payloadBytes, err = json.Marshal(message)
	if err != nil {
		return fmt.Errorf("get message []btye fail. error: %w", err)
	}
	body := bytes.NewReader(payloadBytes)

	dingReq, err := http.NewRequest("POST", "https://oapi.dingtalk.com/robot/send", body)
	if err != nil {
		return fmt.Errorf("get request fail. error: %w", err)
	}
	dingReq.Header.Set("Content-Type", "application/json")

	params := dingReq.URL.Query()
	params.Add("access_token", msg.Token)
	dingReq.URL.RawQuery = params.Encode()

	dingResp, err := http.DefaultClient.Do(dingReq) // 发送请求到钉钉
	if err != nil {
		return fmt.Errorf("request fail. error: %w", err)
	}
	if dingResp != nil {
		defer pkg.CloseBody(dingResp.Body)
	}
	return nil
}
