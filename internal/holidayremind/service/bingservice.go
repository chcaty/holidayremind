package service

import (
	"fmt"
	"github.com/go-co-op/gocron"
	"holidayRemind/configs"
	"holidayRemind/internal/holidayremind/dingtalk"
	"holidayRemind/internal/holidayremind/smtp"
	"holidayRemind/internal/pkg/net"
	"regexp"
	"strings"
)

func BingService() {
	var err error
	imageCron := gocron.NewScheduler(timezone)
	_, err = imageCron.EveryRandom(30, 90).Minutes().Do(bingImageService)
	if err != nil {
		fmt.Printf("bing service Error:%v\n", err.Error())
		return
	}
	imageCron.StartAsync()
}

func bingImageService() error {
	imageUrl := ""
	getBingImage(&imageUrl)
	content := ""
	getSentences(&content)
	var err error
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

func getBingImage(result *string) {
	body := ""
	net.Get("https://tuapi.eees.cc/api.php?category=biying", &body)
	reg, _ := regexp.Compile("src=\".*\"")
	*result = reg.FindString(body)
	*result = strings.Replace(*result, "src=", "", -1)
	*result = strings.Replace(*result, "\"", "", -1)
}

func getSentences(result *string) {
	net.Get("https://v1.hitokoto.cn/?encode=text", result)
}

func sendDingTalk(imageUrl string, content string) error {
	message := dingtalk.Message{
		Title:       "今日推送",
		Text:        fmt.Sprintf(reqImageMD, content, imageUrl),
		Token:       configs.DingTalkToken,
		Tel:         "",
		IsRemind:    false,
		IsRemindAll: false,
	}
	err := dingtalk.SendMdMessage(message)
	if err != nil {
		return err
	}
	return nil
}

func sendEmail(imageUrl string, content string) error {
	bingImageEmail := smtp.EmailMessage{
		Subject:    "今日推送",
		Html:       fmt.Sprintf(reqImageHtml, "%;", "%;", content, imageUrl, "%\""),
		Attachment: nil,
		Receiver:   configs.Receiver,
	}
	err := smtp.SendEmail(bingImageEmail, configs.SmtpConfig)
	if err != nil {
		return err
	}
	return nil
}
