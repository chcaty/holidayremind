package biz

import (
	"fmt"
	"holidayRemind/configs"
	"holidayRemind/internal/holidayremind/template"
	"holidayRemind/internal/pkg/dingtalk"
	"holidayRemind/internal/pkg/rss"
	"holidayRemind/internal/pkg/smtp"
	"log"
	"regexp"
	"strings"
)

type RssUsecase struct {
}

func NewRssUsecase(logger log.Logger) *RssUsecase {
	return &RssUsecase{}
}

func (uc *RssUsecase) GetAppinnRss() error {
	var err error
	appinn := rss.Rss{}
	err = rssRequest("https://appinn.com/feed", &appinn)
	if err != nil {
		return err
	}
	err = rssNotion(appinn.Channel, true, true)
	if err != nil {
		return err
	}
	return nil
}

func (uc *RssUsecase) GetSspaiRss() error {
	var err error
	sspai := rss.Rss{}
	err = rssRequest("https://sspai.com/feed", &sspai)
	if err != nil {
		return err
	}
	cleanSspaiRss(&sspai.Channel)
	err = rssNotion(sspai.Channel, true, true)
	if err != nil {
		return err
	}
	return nil
}

func rssRequest(url string, responseRss *rss.Rss) error {
	err := rss.Request(url, responseRss)
	if err != nil {
		return err
	}
	return nil
}

func cleanZzttRss(channel *rss.Channel) {
	for i := 0; i < len(channel.Item); i++ {
		if channel.Item[i].Title == "黑料不打烊—精选专区" || channel.Item[i].Title == "." || channel.Item[i].Title == ".   " {
			channel.Item = append(channel.Item[:i], channel.Item[i+1:]...)
			// form the remove item index to start iterate next item
			i--
		} else {
			reg, _ := regexp.Compile("萝莉约啪.*偷窥视频")
			channel.Item[i].Description = reg.ReplaceAllString(channel.Item[i].Description, "")
		}
	}
}

func cleanSspaiRss(channel *rss.Channel) {
	for _, item := range channel.Item {
		reg := regexp.MustCompile(`.*<a`)
		results := reg.FindAllString(item.Description, -1)
		if results != nil {
			item.Description = reg.FindAllString(item.Description, -1)[0]
			item.Description = strings.Replace(item.Description, "<a", "", -1)
		}
	}
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
	err := setRssEmailBody(channel, &body)
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

func setRssEmailBody(channel rss.Channel, body *string) error {
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
