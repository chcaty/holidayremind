package service

import (
	"fmt"
	"github.com/go-co-op/gocron"
	"holidayRemind/configs"
	"holidayRemind/internal/holidayremind/moyudock"
	"holidayRemind/internal/holidayremind/smtp"
	"holidayRemind/internal/holidayremind/template"
)

func HotTopService() {
	var err error
	weiboCron := gocron.NewScheduler(timezone)
	_, err = weiboCron.Every(1).Days().At("10:00;15:30").Do(weiboHotTopService)
	if err != nil {
		fmt.Printf("sspai rss Error:%v\n", err.Error())
		return
	}
	weiboCron.StartAsync()
}

func weiboHotTopService() error {
	var err error
	hotTops := moyudock.Response[moyudock.HotTopSite]{}
	err = moyudock.GetHotTop(&hotTops)
	if err != nil {
		return err
	}
	email := smtp.EmailMessage{}
	err = setHotTopEmail(&hotTops.Data.Weibo, "微博热搜榜", &email, configs.HotTopReceiver)
	if err != nil {
		return err
	}
	err = smtp.SendEmail(email, configs.SmtpConfig)
	if err != nil {
		return err
	}
	return nil
}

func setHotTopEmail(hotTopInfos *moyudock.HotTopInfo, title string, email *smtp.EmailMessage, receiver []string) error {
	body := ""
	err := getHotTopEmailBody(hotTopInfos, title, &body)
	if err != nil {
		return err
	}
	email.Subject = title + "推送"
	email.Html = body
	email.Attachment = nil
	email.Receiver = receiver
	return nil
}

func getHotTopEmailBody(hotTopInfos *moyudock.HotTopInfo, title string, body *string) error {
	var err error
	templateMap := map[string]string{}
	err = template.GetTemplateList([]string{
		"HotTopTitle", "HotTopDetail", "HotTopBody",
	}, template.Email, &templateMap)
	if err != nil {
		return err
	}
	title = fmt.Sprintf(templateMap["HotTopTitle"], title)
	content := ""
	for _, info := range (*hotTopInfos).HotTops {
		content += fmt.Sprintf(templateMap["HotTopDetail"], info.Url, info.Title, "热度:"+info.HotValue)
	}
	*body = fmt.Sprintf(templateMap["HotTopBody"], "%;", title, content)
	return nil
}
