package service

import (
	"fmt"
	"holidayRemind/configs"
	"holidayRemind/internal/holidayremind/dingtalk"
	"holidayRemind/internal/holidayremind/scheduler"
	"holidayRemind/internal/holidayremind/smtp"
	"holidayRemind/internal/holidayremind/template"
	"holidayRemind/internal/pkg/net"
	"regexp"
	"strings"
)

func BingService() {
	var err error
	imageScheduler := scheduler.GetSimpleScheduler()
	randomData := scheduler.RandomData{
		Lower: 90,
		Upper: 180,
		Unit:  scheduler.Minute,
	}
	err = imageScheduler.AddRandomJob(randomData, bingImageService)
	if err != nil {
		fmt.Printf("bing service Error:%v\n", err.Error())
		return
	}
	imageScheduler.StartAsync()
}

func bingImageService() error {
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
	err = sendDingTalk(imageUrl, content)
	if err != nil {
		return err
	}
	err = sendEmail(imageUrl, content)
	if err != nil {
		return err
	}
	return nil
}

func getBingImage(result *string) error {
	var resp []byte
	requestData := net.RequestBaseData{
		Url:     "https://tuapi.eees.cc/api.php?category=biying",
		Params:  nil,
		Headers: net.DefaultHeader,
	}
	client := net.GetSimpleHttpClient()
	err := client.Get(&resp, requestData)
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
	requestData := net.RequestBaseData{
		Url:     "https://v1.hitokoto.cn/?encode=text",
		Params:  nil,
		Headers: net.DefaultHeader,
	}
	client := net.GetSimpleHttpClient()
	err := client.Get(&resp, requestData)
	if err != nil {
		return err
	}
	*result = string(resp)
	return nil
}

func sendDingTalk(imageUrl string, content string) error {
	var err error
	imageBody := ""
	err = template.GetTemplate(&imageBody, "ImageBody", template.MarkDown)
	if err != nil {
		return err
	}
	message := dingtalk.Message{
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
