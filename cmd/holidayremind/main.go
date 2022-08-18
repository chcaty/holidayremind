package main

import "holidayRemind/internal/holidayremind/service"

func main() {
	service.BingService()
	service.RssService()
	service.HotTopService()
	service.HolidayService()
}
