package service

import (
	"fmt"
	"github.com/go-co-op/gocron"
	"github.com/golang-module/carbon"
	"holidayRemind/configs"
	"holidayRemind/internal/holidayremind/dingtalk"
	"holidayRemind/internal/holidayremind/holiday"
)

func HolidayService() {
	var err error
	var calendar = map[string]holiday.DayProperty{}
	err = holiday.CreateCalendar(&calendar)
	if err != nil {
		fmt.Printf("create Calendar Error:%v\n", err.Error())
		return
	}
	holidayCron := gocron.NewScheduler(timezone)
	_, err = holidayCron.Every(1).Days().At("10:30").Do(checkTomorrow, calendar)
	if err != nil {
		fmt.Printf("holidayRemind Error:%v\n", err.Error())
		return
	}
	holidayCron.StartBlocking()
}

func checkTomorrow(calender map[string]holiday.DayProperty) error {
	today := carbon.Now().ToDateString()
	todayProp := holiday.DayProperty{}
	err := holiday.GetDayProp(today, calender, &todayProp)
	if err != nil {
		return err
	}
	tomorrow := carbon.Tomorrow().ToDateString()
	tomorrowProp := holiday.DayProperty{}
	err = holiday.GetDayProp(tomorrow, calender, &tomorrowProp)
	if err != nil {
		return err
	}
	if tomorrowProp.IsHoliday == tomorrowProp.IsHoliday {
		return nil
	}
	message := dingtalk.Message{}
	setHolidayMessage(&message, configs.DingTalkToken, tomorrowProp.IsHoliday, tomorrowProp)
	err = sendDingTalkMessage(message, tomorrowProp.IsHoliday)
	if err != nil {
		fmt.Printf("sendMessage Error:%v\n", err.Error())
		return err
	}
	return nil
}

func setHolidayMessage(message *dingtalk.Message, token string, todayFlag bool, tomorrowProp holiday.DayProperty) {
	if !todayFlag && tomorrowProp.IsHoliday {
		message.Title = "放假通知"
		message.Tel = fmt.Sprintf(reqHolidayMD, tomorrowProp.Description)
	} else if todayFlag && !tomorrowProp.IsHoliday {
		message.Title = "上班提醒"
		message.Tel = reqWorkMD
	}
	message.Token = token
}

func sendDingTalkMessage(message dingtalk.Message, tomorrowFlag bool) error {
	if tomorrowFlag {
		err := dingtalk.SendMdMessage(message)
		if err != nil {
			return err
		}
	} else {
		workRemind := gocron.NewScheduler(timezone)
		_, err := workRemind.At("21:30").Do(dingtalk.SendMdMessage, message)
		if err != nil {
			fmt.Printf("workRemind Error:%v\n", err.Error())
		}
		workRemind.StartAsync()
	}
	return nil
}
