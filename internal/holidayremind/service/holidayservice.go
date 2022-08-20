package service

import (
	"fmt"
	"github.com/golang-module/carbon"
	"holidayRemind/configs"
	"holidayRemind/internal/holidayremind/dingtalk"
	"holidayRemind/internal/holidayremind/holiday"
	"holidayRemind/internal/holidayremind/scheduler"
	"holidayRemind/internal/holidayremind/template"
	"time"
)

func HolidayService() {
	var err error
	var calendar = map[string]holiday.DayProperty{}
	err = holiday.CreateCalendar(&calendar)
	if err != nil {
		fmt.Printf("create Calendar Error:%v\n", err.Error())
		return
	}
	holidayScheduler := scheduler.GetSimpleScheduler()
	err = holidayScheduler.AddCornJob("30 10 * * *", checkTomorrow, calendar)
	if err != nil {
		fmt.Printf("holidayRemind Error:%v\n", err.Error())
		return
	}
	holidayScheduler.StartBlocking()
}

func checkTomorrow(calender map[string]holiday.DayProperty) error {
	var err error
	today := carbon.Now().ToDateString()
	todayProp := holiday.DayProperty{}
	err = holiday.GetDayProp(&todayProp, today, calender)
	if err != nil {
		return err
	}
	tomorrow := carbon.Tomorrow().ToDateString()
	tomorrowProp := holiday.DayProperty{}
	err = holiday.GetDayProp(&tomorrowProp, tomorrow, calender)
	if err != nil {
		return err
	}
	if tomorrowProp.IsHoliday == tomorrowProp.IsHoliday {
		return nil
	}
	message := dingtalk.Message{}
	err = setHolidayMessage(&message, configs.DingTalkToken, tomorrowProp.IsHoliday, tomorrowProp)
	if err != nil {
		return err
	}
	err = sendDingTalkMessage(message, tomorrowProp.IsHoliday)
	if err != nil {
		return err
	}
	return nil
}

func setHolidayMessage(message *dingtalk.Message, token string, todayFlag bool, tomorrowProp holiday.DayProperty) error {
	if !todayFlag && tomorrowProp.IsHoliday {
		workBody := ""
		err := template.GetTemplate(&workBody, "WorkBody", template.MarkDown)
		if err != nil {
			return err
		}
		message.Title = "放假通知"
		message.Text = fmt.Sprintf(workBody, "假期提醒姬", tomorrowProp.Description)
	} else if todayFlag && !tomorrowProp.IsHoliday {
		holidayBody := ""
		err := template.GetTemplate(&holidayBody, "HolidayBody", template.MarkDown)
		if err != nil {
			return err
		}
		message.Title = "上班提醒"
		message.Text = fmt.Sprintf(holidayBody, "上班摸鱼姬")
	}
	message.Token = token
	return nil
}

func sendDingTalkMessage(message dingtalk.Message, tomorrowFlag bool) error {
	if tomorrowFlag {
		err := dingtalk.SendMdMessage(message)
		if err != nil {
			return err
		}
	} else {
		workRemindScheduler := scheduler.GetSimpleScheduler()
		corn := ""
		now := time.Now()
		sendTime := time.Date(now.Year(), now.Month(), now.Day(), 22, 30, 0, 0, time.Local)
		scheduler.SetOnceCorn(&corn, sendTime)
		err := workRemindScheduler.AddCornJob(corn, dingtalk.SendMdMessage, message)
		if err != nil {
			fmt.Printf("workRemind Error:%v\n", err.Error())
		}
		workRemindScheduler.StartAsync()
	}
	return nil
}
