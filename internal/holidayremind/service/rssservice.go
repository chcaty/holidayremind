package service

import (
	"fmt"
	"github.com/go-co-op/gocron"
	"holidayRemind/configs"
	"holidayRemind/internal/holidayremind/dingtalk"
	"holidayRemind/internal/holidayremind/rss"
	"holidayRemind/internal/holidayremind/smtp"
)

func RssService() {
	var err error
	sspaiCron := gocron.NewScheduler(timezone)
	_, err = sspaiCron.Every(1).Days().At("10:30;15:30").Do(sspaiRssServer)
	if err != nil {
		fmt.Printf("sspai rss Error:%v\n", err.Error())
		return
	}
	sspaiCron.StartAsync()

	appinnCron := gocron.NewScheduler(timezone)
	_, err = appinnCron.Every(1).Days().At("10:30;15:30").Do(appinRssServer)
	if err != nil {
		fmt.Printf("appinn rss Error:%v\n", err.Error())
		return
	}
	appinnCron.StartAsync()
}

func sspaiRssServer() error {
	var err error
	// 获取Rss信息
	sspaiRss := rss.Rss{}
	sspai := rss.RequestData{
		Url:         "https://sspai.com/feed",
		ChannelType: rss.Sspai,
	}
	err = rss.Request(sspai, &sspaiRss)
	if err != nil {
		return err
	}
	err = rssNotion(sspaiRss.Channel, true, true)
	if err != nil {
		return err
	}
	return nil
}

func appinRssServer() error {
	var err error
	// 获取Rss信息
	appinnRss := rss.Rss{}
	appinn := rss.RequestData{
		Url:         "https://appin.com/feed",
		ChannelType: rss.Appinn,
	}
	err = rss.Request(appinn, &appinnRss)
	if err != nil {
		return err
	}
	err = rssNotion(appinnRss.Channel, true, true)
	if err != nil {
		return err
	}
	return nil
}

func rssNotion(channel rss.Channel, isDingTalk bool, isEmail bool) error {
	if isEmail {
		// 发送邮件
		sspaiEmail := smtp.EmailMessage{}
		setRssEmail(channel, &sspaiEmail, configs.Receiver)
		err := smtp.SendEmail(sspaiEmail, configs.SmtpConfig)
		if err != nil {
			return err
		}
	}
	if isDingTalk {
		// 推送到钉钉机器人
		sspaiMessage := dingtalk.Message{}
		setRssMessage(channel, &sspaiMessage, configs.DingTalkToken, "")
		err := dingtalk.SendMdMessage(sspaiMessage)
		if err != nil {
			return err
		}
	}
	return nil
}

func setRssMessage(channel rss.Channel, message *dingtalk.Message, token string, tel string) {
	content := ""
	for _, item := range channel.Item {
		content += fmt.Sprintf(reqRssContent, item.Title, item.Link, item.Description)
	}
	text := fmt.Sprintf(reqRssMD, channel.Title, content)
	message.Title = channel.Title + "今日推送"
	message.Text = text
	message.Token = token
	if len(tel) > 0 {
		message.IsRemind = true
		message.Tel = tel
	}
}

func setRssEmail(channel rss.Channel, email *smtp.EmailMessage, receiver []string) {
	body := ""
	getEmailBody(channel, &body)
	email.Subject = channel.Title + "推送"
	email.Html = body
	email.Attachment = nil
	email.Receiver = receiver
}

func getEmailBody(channel rss.Channel, body *string) {
	title := fmt.Sprintf(emailBodyTitle, channel.Title)
	content := ""
	for _, item := range channel.Item {
		content += fmt.Sprintf(emailBodyContent, item.PubDate, item.Link, item.Title, item.Description)
	}
	*body = fmt.Sprintf(emailBody, title, content)
}