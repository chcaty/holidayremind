package moyuduckservice

import (
	"fmt"
	"holidayRemind/configs"
	"holidayRemind/internal/holidayremind/moyudock"
	"holidayRemind/internal/holidayremind/scheduler"
	"holidayRemind/internal/holidayremind/smtp"
	"holidayRemind/internal/holidayremind/template"
	"log"
)

func Start() {
	hotTopScheduler := scheduler.GetSimpleScheduler()
	err := hotTopScheduler.AddCornJob("30 10,15 * * *", false, "hotTop", weiboHotTopService)
	if err != nil {
		log.Printf("hotTop Error:%v\n", err.Error())
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
	err = sendHotTopEmail("微博热搜榜", &hotTops.Data.Weibo, email)
	if err != nil {
		return err
	}
	return nil
}

func sendHotTopEmail(title string, hotTopInfos *moyudock.HotTopInfo, email smtp.SimpleEmail) error {
	body := ""
	err := getHotTopEmailBody(title, *hotTopInfos, &body)
	if err != nil {
		return err
	}
	email.Subject = title + "推送"
	email.Html = body
	err = smtp.SendEmail(email, configs.SmtpConfig)
	if err != nil {
		return err
	}
	return nil
}

func getHotTopEmailBody(title string, hotTopInfos moyudock.HotTopInfo, body *string) error {
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
	for _, info := range (hotTopInfos).HotTops {
		content += fmt.Sprintf(templateMap["HotTopDetail"], info.Url, info.Title, "热度:"+info.HotValue)
	}
	*body = fmt.Sprintf(templateMap["HotTopBody"], "%;", title, content)
	return nil
}
