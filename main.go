package main

import (
	"fmt"
	"github.com/go-co-op/gocron"
	"github.com/golang-module/carbon"
	"holidayRemind/common"
	"holidayRemind/holiday"
	"time"
)

var calendar = map[string]holiday.DayProperty{}
var token = "afc3c084e0a0a7936196b6a686f9bd382dcb5859609ee58b7c234ff6d94ad929"
var timezone, _ = time.LoadLocation("Asia/Shanghai")

func main() {
	holiday.CreateCalendar(calendar)

	holidayRemind := gocron.NewScheduler(timezone)
	_, err := holidayRemind.Every(1).Days().At("10:30").Do(CheckTomorrowHoliday)
	if err != nil {
		fmt.Printf("holidayRemind Error:%v\n", err.Error())
		return
	}
	holidayRemind.StartBlocking()
}

func CheckTomorrowHoliday() {
	now := carbon.Now().ToDateString()
	if today, ok := calendar[now]; ok {
		tomorrow := carbon.Tomorrow().ToDateString()
		if value, ok := calendar[tomorrow]; ok {
			title, text := "", ""
			if !today.IsHoliday && value.IsHoliday {
				title = "放假通知"
				text = fmt.Sprintf(holiday.ReqHolidayMD, value.Description)
			} else if today.IsHoliday && !value.IsHoliday {
				title = "上班提醒"
				text = holiday.ReqWorkMD
			} else {
				return
			}
			message := &common.Message{
				Title:    title,
				Text:     text,
				Token:    token,
				Tel:      "",
				IsRemind: false,
			}
			if value.IsHoliday {
				_, err := holiday.CommonSendMessage(*message)
				if err != nil {
					fmt.Printf("sendMessage Error:%v\n", err.Error())
					return
				}
			} else {
				workRemind := gocron.NewScheduler(timezone)
				_, err := workRemind.At("18:00").Do(CheckTomorrowHoliday)
				if err != nil {
					fmt.Printf("workRemind Error:%v\n", err.Error())
				}
				workRemind.StartBlocking()
			}
		}
	}
}
