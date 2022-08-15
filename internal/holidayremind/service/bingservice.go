package service

import (
	"fmt"
	"github.com/go-co-op/gocron"
	"github.com/golang-module/carbon"
	"holidayRemind/configs"
	"holidayRemind/internal/holidayremind/dingtalk"
	"holidayRemind/internal/holidayremind/smtp"
)

func BingService() {
	var err error
	imageCron := gocron.NewScheduler(timezone)
	_, err = imageCron.Every(1).Days().At("10:00;16:00").Do(bingImageService)
	if err != nil {
		fmt.Printf("appinn rss Error:%v\n", err.Error())
		return
	}
	imageCron.StartAsync()
}

func bingImageService() error {
	var err error
	err = sendDingTalk()
	if err != nil {
		return err
	}
	fmt.Println("rss send dingtalk success at ", carbon.Now().ToDateTimeString())
	err = sendEmail()
	if err != nil {
		return err
	}
	fmt.Println("rss send email success at ", carbon.Now().ToDateTimeString())
	return nil
}

func sendDingTalk() error {
	message := dingtalk.Message{
		Title:       "今日美图推送",
		Text:        reqImageMD,
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

func sendEmail() error {
	bingImageEmail := smtp.EmailMessage{
		Subject:    "今日美图推送",
		Html:       reqImageHtml,
		Attachment: nil,
		Receiver:   configs.Receiver,
	}
	err := smtp.SendEmail(bingImageEmail, configs.SmtpConfig)
	if err != nil {
		return err
	}
	return nil
}
