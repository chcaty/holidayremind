package rss

import (
	"fmt"
	"holidayRemind/dingtalk"
	"holidayRemind/smtp"
)

func sendRssMessage(channel *Channel, token string) {
	content := ""
	for _, item := range channel.Item {
		content += fmt.Sprintf(reqRssContent, item.Title, item.Link, item.Description)
	}
	text := fmt.Sprintf(reqRssMD, channel.Title, content)
	message := &dingtalk.Message{
		Title:    channel.Title + "今日推送",
		Text:     text,
		Token:    token,
		Tel:      "",
		IsRemind: false,
	}
	_, err := dingtalk.SendMessage(*message)
	if err != nil {
		fmt.Printf("sendMessage Error:%v\n", err.Error())
	}
}

func sendRssEmail(channel *Channel, receiver []string) {
	body := ""
	getEmailBody(channel, &body)
	err := smtp.SendEmail(channel.Title+"推送", body, nil, receiver)
	if err != nil {
		fmt.Printf("sendMessage Error:%v\n", err.Error())
	}
}

func getEmailBody(channel *Channel, body *string) {
	title := fmt.Sprintf(emailBodyTitle, channel.Title)
	content := ""
	for _, item := range channel.Item {
		content += fmt.Sprintf(emailBodyContent, item.PubDate, item.Link, item.Title, item.Description)
	}
	*body = fmt.Sprintf(emailBody, title, content)
}
