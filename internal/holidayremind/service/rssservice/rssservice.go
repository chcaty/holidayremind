package rssservice

import (
	"fmt"
	"github.com/go-co-op/gocron"
	"holidayRemind/configs"
	"holidayRemind/internal/holidayremind/dingtalk"
	"holidayRemind/internal/holidayremind/rss"
	"holidayRemind/internal/holidayremind/smtp"
	"holidayRemind/internal/holidayremind/template"
	"holidayRemind/internal/pkg/scheduler"
	"log"
)

func Start() {
	var err error
	rssScheduler := scheduler.GetScheduler()
	err = rssScheduler.AddCornJob("30 10,15 * * *", true, "sspaiRss", sspaiRssServer)
	if err != nil {
		log.Printf("sspai rss Error:%v\n", err.Error())
		return
	}

	err = rssScheduler.AddCornJob("0 11,16 * * *", false, "appinnRss", appinnRssServer)
	if err != nil {
		fmt.Printf("appinn rss Error:%v\n", err.Error())
		return
	}
	rssScheduler.StartAsync()
}

func sspaiRssServer(job gocron.Job) error {
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
	fmt.Printf("sspai rss job's last run: %s\nthis job's next run: %s", job.LastRun(), job.NextRun())
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
		message := dingtalk.MarkdownMessageDTO{}
		err = sendRssMessage(channel, configs.DingTalkToken, "", message)
		if err != nil {
			return err
		}
	}
	if isEmail {
		// 发送邮件
		email := smtp.SimpleEmail{
			Receiver: configs.Receiver,
		}
		err = sendRssEmail(channel, email)
		if err != nil {
			return err
		}
	}
	return nil
}

func sendRssMessage(channel rss.Channel, token string, tel string, message dingtalk.MarkdownMessageDTO) error {
	var err error
	content := ""
	templateMap := map[string]string{}
	err = template.GetTemplateList(&templateMap, []string{
		"RssDetail", "RssBody",
	}, template.MarkDown)
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
	err = dingtalk.SendMdMessage(message)
	if err != nil {
		return err
	}
	return nil
}

func sendRssEmail(channel rss.Channel, email smtp.SimpleEmail) error {
	body := ""
	err := getEmailBody(channel, &body)
	if err != nil {
		return err
	}
	email.Subject = channel.Title + "推送"
	email.Html = body
	err = smtp.SendEmail(email, configs.SmtpConfig)
	if err != nil {
		return err
	}
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
