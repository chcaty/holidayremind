package service

import (
	"fmt"
	"github.com/go-co-op/gocron"
	"holidayRemind/configs"
	"holidayRemind/internal/holidayremind/dingtalk"
	"holidayRemind/internal/holidayremind/rss"
	"holidayRemind/internal/holidayremind/scheduler"
	"holidayRemind/internal/holidayremind/smtp"
	"holidayRemind/internal/holidayremind/template"
)

func RssService() {
	var err error
	sspaiScheduler := gocron.Scheduler{}
	err = scheduler.SetScheduler(&sspaiScheduler, "30 10,15 * * *", sspaiRssServer)
	if err != nil {
		fmt.Printf("sspai rss Error:%v\n", err.Error())
		return
	}
	sspaiScheduler.StartAsync()

	appinnScheduler := gocron.Scheduler{}
	err = scheduler.SetScheduler(&appinnScheduler, "0 11,16 * * *", appinnRssServer)
	if err != nil {
		fmt.Printf("appinn rss Error:%v\n", err.Error())
		return
	}
	appinnScheduler.StartAsync()
}

func sspaiRssServer() error {
	var err error
	sspaiRss := rss.Rss{}
	err = rssRequest("https://sspai.com/feed", rss.Sspai, &sspaiRss)
	if err != nil {
		return err
	}
	err = rssNotion(sspaiRss.Channel, true, true)
	if err != nil {
		return err
	}
	return nil
}

func appinnRssServer() error {
	var err error
	appinnRss := rss.Rss{}
	err = rssRequest("https://appinn.com/feed", rss.Appinn, &appinnRss)
	if err != nil {
		return err
	}
	err = rssNotion(appinnRss.Channel, true, true)
	if err != nil {
		return err
	}
	return nil
}

func rssRequest(url string, rssType rss.ChannelType, responseRss *rss.Rss) error {
	rssData := rss.RequestData{
		Url:         url,
		ChannelType: rssType,
	}
	err := rss.Request(rssData, responseRss)
	if err != nil {
		return err
	}
	return nil
}

func rssNotion(channel rss.Channel, isDingTalk bool, isEmail bool) error {
	var err error
	if isDingTalk {
		// 推送到钉钉机器人
		message := dingtalk.Message{}
		err = setRssMessage(channel, &message, configs.DingTalkToken, "")
		if err != nil {
			return err
		}
		err = dingtalk.SendMdMessage(message)
		if err != nil {
			return err
		}
	}
	if isEmail {
		// 发送邮件
		email := smtp.SimpleEmail{
			Receiver: configs.Receiver,
		}
		err = setRssEmail(channel, &email)
		if err != nil {
			return err
		}
		err = smtp.SendEmail(email, configs.SmtpConfig)
		if err != nil {
			return err
		}
	}
	return nil
}

func setRssMessage(channel rss.Channel, message *dingtalk.Message, token string, tel string) error {
	var err error
	content := ""
	templateMap := map[string]string{}
	err = template.GetTemplateList(&templateMap, []string{
		"RssDetail", "RssBody",
	}, template.Email)
	if err != nil {
		return err
	}
	for _, item := range channel.Item {
		content += fmt.Sprintf(templateMap["RssDetail"], item.Title, item.Link, item.Description)
	}
	text := fmt.Sprintf(templateMap["RssBody"], "Rss通知姬", channel.Title, content)
	message.Title = channel.Title + "今日推送"
	message.Text = text
	message.Token = token
	if len(tel) > 0 {
		message.IsRemind = true
		message.Tel = tel
	}
	return nil
}

func setRssEmail(channel rss.Channel, email *smtp.SimpleEmail) error {
	body := ""
	err := getEmailBody(channel, &body)
	if err != nil {
		return err
	}
	email.Subject = channel.Title + "推送"
	email.Html = body
	return nil
}

func getEmailBody(channel rss.Channel, body *string) error {
	var err error
	templateMap := map[string]string{}
	err = template.GetTemplateList(&templateMap, []string{
		"RssTitle", "RssDetail", "RssBody",
	}, template.Email)
	if err != nil {
		return err
	}
	title := fmt.Sprintf(templateMap["RssTitle"], channel.Title)
	content := ""
	for _, item := range channel.Item {
		content += fmt.Sprintf(templateMap["RssDetail"], item.PubDate, item.Link, item.Title, item.Description)
	}
	*body = fmt.Sprintf(templateMap["RssBody"], "%;", title, content)
	return nil
}
