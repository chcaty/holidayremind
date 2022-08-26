package moyuduckservice

import (
	"fmt"
	"holidayRemind/configs"
	"holidayRemind/internal/holidayremind/moyuduck"
	"holidayRemind/internal/holidayremind/template"
	"holidayRemind/internal/pkg/scheduler"
	smtp2 "holidayRemind/internal/pkg/smtp"
	"log"
)

func Start() {
	hotTopScheduler := scheduler.GetScheduler()
	err := hotTopScheduler.AddCornJob("30 10,15 * * *", false, "hotTop", weiboHotTopService)
	if err != nil {
		log.Printf("hotTop Error:%v\n", err.Error())
		return
	}
	hotTopScheduler.StartAsync()
}

func weiboHotTopService() error {
	var err error
	hotTops := moyuduck.Response[moyuduck.HotTopSite]{}
	err = moyuduck.GetHotTop(&hotTops)
	if err != nil {
		return err
	}
	email := smtp2.SimpleEmail{
		Receiver: configs.HotTopReceiver,
	}
	err = sendHotTopEmail("微博热搜榜", &hotTops.Data.Weibo, email)
	if err != nil {
		return err
	}
	return nil
}

func sendHotTopEmail(title string, hotTopInfos *moyuduck.HotTopInfo, email smtp2.SimpleEmail) error {
	body := ""
	err := getHotTopEmailBody(title, *hotTopInfos, &body)
	if err != nil {
		return err
	}
	email.Subject = title + "推送"
	email.Html = body
	err = smtp2.SendEmail(email, configs.SmtpConfig)
	if err != nil {
		return err
	}
	return nil
}

func getHotTopEmailBody(title string, hotTopInfos moyuduck.HotTopInfo, body *string) error {
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
