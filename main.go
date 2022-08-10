package main

import (
	"fmt"
	"github.com/go-co-op/gocron"
	"holidayRemind/holiday"
	"holidayRemind/rss"
	"time"
)

var timezone, _ = time.LoadLocation("Asia/Shanghai")
var calendar = map[string]holiday.DayProperty{}
var token = "afc3c084e0a0a7936196b6a686f9bd382dcb5859609ee58b7c234ff6d94ad929"

func main() {
	var err error
	rssCron := gocron.NewScheduler(timezone)
	_, err = rssCron.Every(1).Days().At("10:30;15:30").Do(rss.SspaiRssRequest, token)
	if err != nil {
		fmt.Printf("rss Error:%v\n", err.Error())
		return
	}
	rssCron.StartAsync()

	holiday.CreateCalendar(calendar)
	holidayCron := gocron.NewScheduler(timezone)
	_, err = holidayCron.Every(1).Days().At("10:30").Do(holiday.CheckTomorrowHoliday, &calendar, token)
	if err != nil {
		fmt.Printf("holidayRemind Error:%v\n", err.Error())
		return
	}
	holidayCron.StartBlocking()
}
