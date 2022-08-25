package bingservice

import (
	"fmt"
	"github.com/go-co-op/gocron"
	"holidayRemind/configs"
	"holidayRemind/internal/holidayremind/dingtalk"
	"holidayRemind/internal/holidayremind/scheduler"
	"holidayRemind/internal/holidayremind/smtp"
	"holidayRemind/internal/holidayremind/template"
	"holidayRemind/internal/pkg/uxnet"
	"log"
	"regexp"
	"strings"
)

func Start() {
	var err error
	imageScheduler := scheduler.GetSimpleScheduler()
	randomData := scheduler.RandomData{
		Lower: 90,
		Upper: 180,
		Unit:  scheduler.Minute,
	}
	err = imageScheduler.AddRandomJob(randomData, true, "BingImage", bingImageService)
	if err != nil {
		log.Printf("bingservice service Error:%v", err.Error())
		return
	}
	imageScheduler.StartAsync()
}

func bingImageService(job gocron.Job) error {
	var err error
	imageUrl := ""
	err = getBingImage(&imageUrl)
	if err != nil {
		return err
	}
	content := ""
	err = getSentences(&content)
	if err != nil {
		return err
	}
	err = sendDingTalkMessage(imageUrl, content)
	if err != nil {
		return err
	}
	err = sendEmail(imageUrl, content)
	if err != nil {
		return err
	}
	log.Printf("bingImage job's last run: %s this job's next run: %s", job.LastRun(), job.NextRun())
	return nil
}

func getBingImage(result *string) error {
	var resp []byte
	requestData := uxnet.BaseData{
		Url:     "https://tuapi.eees.cc/api.php?category=biying",
		Params:  nil,
		Headers: uxnet.DefaultHeader,
	}
	client := uxnet.GetHttpClient()
	err := client.Get(requestData, &resp)
	if err != nil {
		return err
	}
	body := string(resp)
	reg, _ := regexp.Compile("src=\".*\"")
	*result = reg.FindString(body)
	*result = strings.Replace(*result, "src=", "", -1)
	*result = strings.Replace(*result, "\"", "", -1)
	return nil
}

func getSentences(result *string) error {
	var resp []byte
	requestData := uxnet.BaseData{
		Url:     "https://v1.hitokoto.cn/?encode=text",
		Params:  nil,
		Headers: uxnet.DefaultHeader,
	}
	client := uxnet.GetHttpClient()
	err := client.Get(requestData, &resp)
	if err != nil {
		return err
	}
	*result = string(resp)
	return nil
}

func sendDingTalkMessage(imageUrl string, content string) error {
	var err error
	imageBody := ""
	err = template.GetTemplate(&imageBody, "ImageBody", template.MarkDown)
	if err != nil {
		return err
	}
	message := dingtalk.MarkdownMessageDTO{
		Title:       "今日推送",
		Text:        fmt.Sprintf(imageBody, "发图姬", content, imageUrl),
		Token:       configs.DingTalkToken,
		Tel:         "",
		IsRemind:    false,
		IsRemindAll: false,
	}
	err = dingtalk.SendMdMessage(message)
	if err != nil {
		return err
	}
	return nil
}

func sendEmail(imageUrl string, content string) error {
	var err error
	imageBody := ""
	err = template.GetTemplate(&imageBody, "ImageBody", template.Email)
	if err != nil {
		return err
	}
	bingImageEmail := smtp.SimpleEmail{
		Subject:    "今日推送",
		Html:       fmt.Sprintf(imageBody, "%;", "%;", content, imageUrl, "%\""),
		Attachment: nil,
		Receiver:   configs.Receiver,
	}
	err = smtp.SendEmail(bingImageEmail, configs.SmtpConfig)
	if err != nil {
		return err
	}
	return nil
}
