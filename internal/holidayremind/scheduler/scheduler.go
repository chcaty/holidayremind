package scheduler

import (
	"fmt"
	"github.com/go-co-op/gocron"
	"github.com/golang-module/carbon"
	"time"
)

var timezone, _ = time.LoadLocation("Asia/Shanghai")

func SetScheduler(jobCron *gocron.Scheduler, corn string, jobFun interface{}, params ...interface{}) error {
	var err error
	jobCron = gocron.NewScheduler(timezone)
	_, err = jobCron.Cron(corn).Do(jobFun, params...)
	if err != nil {
		return err
	}
	return nil
}

func SetOnceCorn(corn *string, date time.Time) {
	carbonDate := carbon.Time2Carbon(date)
	*corn = fmt.Sprintf("%d %d %d %d %d", carbonDate.Minute(), carbonDate.Hour(), carbonDate.Day(), carbonDate.Month(), carbonDate.Year())
}

func SetRandomScheduler(jobCron *gocron.Scheduler, data RandomData, jobFun interface{}, params ...interface{}) error {
	var err error
	jobCron = gocron.NewScheduler(timezone)
	switch data.Unit {
	case Day:
		_, err = jobCron.EveryRandom(data.Lower, data.Upper).Days().Do(jobFun, params...)
	case Hour:
		_, err = jobCron.EveryRandom(data.Lower, data.Upper).Hours().Do(jobFun, params...)
	case Minute:
		_, err = jobCron.EveryRandom(data.Lower, data.Upper).Minutes().Do(jobFun, params...)
	case Second:
		_, err = jobCron.EveryRandom(data.Lower, data.Upper).Seconds().Do(jobFun, params...)
	}
	if err != nil {
		return err
	}
	return nil
}
