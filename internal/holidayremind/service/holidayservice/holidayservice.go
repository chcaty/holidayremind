package holidayservice

import (
	"fmt"
	"github.com/golang-module/carbon"
	"holidayRemind/configs"
	"holidayRemind/internal/holidayremind/dingtalk"
	"holidayRemind/internal/holidayremind/holiday"
	"holidayRemind/internal/holidayremind/scheduler"
	"holidayRemind/internal/holidayremind/template"
	"holidayRemind/internal/pkg/uxmap"
	"log"
	"time"
)

func Start() {
	var err error
	var calendar = map[string]holiday.DayProperty{}
	err = holiday.CreateCalendar(&calendar)
	if err != nil {
		log.Printf("create Calendar Error:%v", err.Error())
		return
	}
	holidayScheduler := scheduler.GetSimpleScheduler()
	err = holidayScheduler.AddCornJob("30 10 * * *", false, "holidayRemind", checkTomorrow, calendar)
	if err != nil {
		log.Printf("holidayRemind Error:%v", err.Error())
		return
	}
	holidayScheduler.StartAsync()
}

func checkTomorrow(calender map[string]holiday.DayProperty) error {
	var err error
	today := carbon.Now().ToDateString()
	todayProp := holiday.DayProperty{}
	err = uxmap.GetMapValue(calender, today, &todayProp)
	if err != nil {
		return err
	}
	tomorrow := carbon.Tomorrow().ToDateString()
	tomorrowProp := holiday.DayProperty{}
	err = uxmap.GetMapValue(calender, tomorrow, &tomorrowProp)
	if err != nil {
		return err
	}
	if tomorrowProp.IsHoliday == tomorrowProp.IsHoliday {
		return nil
	}
	message := dingtalk.MarkdownMessageDTO{}
	err = setHolidayMessage(configs.DingTalkToken, tomorrowProp.IsHoliday, tomorrowProp, &message)
	if err != nil {
		return err
	}
	err = sendDingTalkMessage(message, tomorrowProp.IsHoliday)
	if err != nil {
		return err
	}
	return nil
}

func setHolidayMessage(token string, todayFlag bool, tomorrowProp holiday.DayProperty, message *dingtalk.MarkdownMessageDTO) error {
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

func sendDingTalkMessage(message dingtalk.MarkdownMessageDTO, tomorrowFlag bool) error {
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
		err := workRemindScheduler.AddCornJob(corn, false, "workRemind", dingtalk.SendMdMessage, message)
		if err != nil {
			log.Printf("workRemind Error:%v", err.Error())
		}
		workRemindScheduler.StartAsync()
	}
	return nil
}
