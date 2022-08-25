package dingtalk

import (
	"holidayRemind/configs"
	"holidayRemind/internal/holidayremind/dingtalk"
	"testing"
)

func TestSendMdMessage(t *testing.T) {
	message := dingtalk.MarkdownMessageDTO{
		Title:       "test",
		Text:        "test",
		Token:       configs.DingTalkToken,
		Tel:         "",
		IsRemind:    false,
		IsRemindAll: false,
	}
	err := dingtalk.SendMdMessage(message)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log("success")
}

func TestSendFeedCardMessage(t *testing.T) {
	message := dingtalk.FeedCardMessageDTO{
		Links: nil,
		Token: configs.DingTalkToken,
	}
	err := dingtalk.SendFeedCardMessage(message)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log("success")
}
