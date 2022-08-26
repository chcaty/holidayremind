package main

import (
	"holidayRemind/internal/holidayremind/service/bingservice"
	"holidayRemind/internal/holidayremind/service/doubanservice"
	"holidayRemind/internal/holidayremind/service/holidayservice"
	"holidayRemind/internal/holidayremind/service/moyuduckservice"
	"holidayRemind/internal/holidayremind/service/vvhanservice"
)

func main() {
	bingservice.Start()
	moyuduckservice.Start()
	holidayservice.Start()
	doubanservice.Start()
	vvhanservice.Start()
	select {}
}
