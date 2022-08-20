package service

import (
	"fmt"
	"holidayRemind/configs"
	"holidayRemind/internal/holidayremind/moyudock"
	"holidayRemind/internal/holidayremind/scheduler"
	"holidayRemind/internal/holidayremind/smtp"
	"holidayRemind/internal/holidayremind/template"
)

func HotTopService() {
	hotTopScheduler := scheduler.GetSimpleScheduler()
	err := hotTopScheduler.AddCornJob("30 10,15 * * *", weiboHotTopService)
	if err != nil {
		fmt.Printf("hotTop Error:%v\n", err.Error())
		return
	}
	hotTopScheduler.StartAsync()
}

func weiboHotTopService() error {
	var err error
	hotTops := moyudock.Response[moyudock.HotTopSite]{}
	err = moyudock.GetHotTop(&hotTops)
	if err != nil {
		return err
	}
	email := smtp.SimpleEmail{
		Receiver: configs.HotTopReceiver,
	}
	err = setHotTopEmail(&hotTops.Data.Weibo, "微博热搜榜", &email)
	if err != nil {
		return err
	}
	err = smtp.SendEmail(email, configs.SmtpConfig)
	if err != nil {
		return err
	}
	return nil
}

func setHotTopEmail(hotTopInfos *moyudock.HotTopInfo, title string, email *smtp.SimpleEmail) error {
	body := ""
	err := getHotTopEmailBody(hotTopInfos, title, &body)
	if err != nil {
		return err
	}
	email.Subject = title + "推送"
	email.Html = body
	return nil
}

func getHotTopEmailBody(hotTopInfos *moyudock.HotTopInfo, title string, body *string) error {
	var err error
	templateMap := map[string]string{}
	templateName := []string{
		"HotTopTitle", "HotTopDetail", "HotTopBody",
	}
	err = template.GetTemplateList(&templateMap, templateName, template.Email)
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
