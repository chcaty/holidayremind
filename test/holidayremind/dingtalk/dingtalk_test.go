package dingtalk

import (
	"holidayRemind/configs"
	dingtalk2 "holidayRemind/internal/pkg/dingtalk"
	"testing"
)

func TestSendMdMessage(t *testing.T) {
	message := dingtalk2.MarkdownMessageDTO{
		Title:       "test",
		Text:        "test",
		Token:       configs.DingTalkToken,
		Tel:         "",
		IsRemind:    false,
		IsRemindAll: false,
	}
	err := dingtalk2.SendMdMessage(message)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log("success")
}

func TestSendFeedCardMessage(t *testing.T) {
	message := dingtalk2.FeedCardMessageDTO{
		Links: nil,
		Token: configs.DingTalkToken,
	}
	err := dingtalk2.SendFeedCardMessage(message)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log("success")
}
