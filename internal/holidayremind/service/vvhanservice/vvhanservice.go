package vvhanservice

import (
	"errors"
	"fmt"
	"holidayRemind/configs"
	"holidayRemind/internal/holidayremind/template"
	"holidayRemind/internal/pkg/scheduler"
	smtp2 "holidayRemind/internal/pkg/smtp"
	"holidayRemind/internal/pkg/uxjson"
	"holidayRemind/internal/pkg/uxnet"
	"log"
)

func Start() {
	var err error
	vvhanScheduler := scheduler.GetScheduler()
	err = vvhanScheduler.AddCornJob("0 10 * * *", false, "vvhan", VvhanImageService)
	if err != nil {
		log.Printf("vvhan service Error:%v\n", err.Error())
		return
	}
	vvhanScheduler.StartAsync()
}

func VvhanImageService() error {
	var err error
	imageUrl := ""
	err = getCalendarImage(&imageUrl)
	if err != nil {
		return err
	}
	err = sendEmail(imageUrl)
	if err != nil {
		return err
	}
	return nil
}

func getCalendarImage(imageUrl *string) error {
	var err error
	var resp []byte
	requestData := uxnet.BaseData{
		Url:     "https://api.vvhan.com/api/moyu?type=json",
		Params:  nil,
		Headers: uxnet.DefaultHeader,
	}
	client := uxnet.GetHttpClient()
	err = client.Get(requestData, &resp)
	if err != nil {
		return err
	}
	moyu := MoyuResponse{}
	err = uxjson.ToObjectByBytes(resp, &moyu)
	if err != nil {
		return err
	}
	if moyu.Success {
		*imageUrl = moyu.Url
	} else {
		return fmt.Errorf("%w", errors.New("get moyu Image fail"))
	}
	return nil
}

func sendEmail(imageUrl string) error {
	var err error
	imageBody := ""
	err = template.GetTemplate(&imageBody, "CalendarBody", template.Email)
	if err != nil {
		return err
	}
	bingImageEmail := smtp2.SimpleEmail{
		Subject:    "摸鱼人推送",
		Html:       fmt.Sprintf(imageBody, "%;", "%;", imageUrl, "%\""),
		Attachment: nil,
		Receiver:   configs.Receiver,
	}
	err = smtp2.SendEmail(bingImageEmail, configs.SmtpConfig)
	if err != nil {
		return err
	}
	return nil
}

type MoyuResponse struct {
	Success bool   `json:"success"`
	Url     string `json:"url"`
}
