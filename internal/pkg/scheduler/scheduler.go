package scheduler

import (
	"fmt"
	"github.com/go-co-op/gocron"
	"github.com/golang-module/carbon"
	"sync"
	"time"
)

var timezone, _ = time.LoadLocation("Asia/Shanghai")
var simple *SimpleScheduler
var once sync.Once

type SimpleScheduler struct {
	scheduler *gocron.Scheduler
}

type Jobber interface {
	AddCornJob(corn string, isWithJobDetail bool, tag string, jobFun interface{}, params ...interface{}) error
	AddRandomJob(data RandomData, isWithJobDetail bool, tag string, jobFun interface{}, params ...interface{}) error
}

type Starter interface {
	StartAsync()
	StartBlocking()
	RunByTag(tag string) error
}

// GetScheduler 生成一个SimpleScheduler对象
func GetScheduler() *SimpleScheduler {
	once.Do(func() {
		simple = &SimpleScheduler{
			scheduler: gocron.NewScheduler(timezone),
		}
	})
	return simple
}

// AddCornJob 增加一个由Corn表达式决定何时触发的定时任务
func (s *SimpleScheduler) AddCornJob(corn string, isWithJobDetail bool, tag string, jobFun interface{}, params ...interface{}) error {
	var err error
	if isWithJobDetail {
		_, err = s.scheduler.Cron(corn).Tag(tag).DoWithJobDetails(jobFun, params...)
	} else {
		_, err = s.scheduler.Cron(corn).Tag(tag).Do(jobFun, params...)
	}
	if err != nil {
		return err
	}
	return nil
}

// AddRandomJob 增加一个一定范围内随机触发的定时任务
func (s *SimpleScheduler) AddRandomJob(data RandomData, isWithJobDetail bool, tag string, jobFun interface{}, params ...interface{}) error {
	var err error
	switch data.Unit {
	case Day:
		if isWithJobDetail {
			_, err = s.scheduler.EveryRandom(data.Lower, data.Upper).Days().Tag(tag).DoWithJobDetails(jobFun, params...)
		} else {
			_, err = s.scheduler.EveryRandom(data.Lower, data.Upper).Days().Tag(tag).Do(jobFun, params...)
		}
	case Hour:
		if isWithJobDetail {
			_, err = s.scheduler.EveryRandom(data.Lower, data.Upper).Hours().Tag(tag).DoWithJobDetails(jobFun, params...)
		} else {
			_, err = s.scheduler.EveryRandom(data.Lower, data.Upper).Hours().Tag(tag).Do(jobFun, params...)
		}
	case Minute:
		if isWithJobDetail {
			_, err = s.scheduler.EveryRandom(data.Lower, data.Upper).Minutes().Tag(tag).DoWithJobDetails(jobFun, params...)
		} else {
			_, err = s.scheduler.EveryRandom(data.Lower, data.Upper).Minutes().Tag(tag).Do(jobFun, params...)
		}
	case Second:
		if isWithJobDetail {
			_, err = s.scheduler.EveryRandom(data.Lower, data.Upper).Seconds().Tag(tag).DoWithJobDetails(jobFun, params...)
		} else {
			_, err = s.scheduler.EveryRandom(data.Lower, data.Upper).Seconds().Tag(tag).Do(jobFun, params...)
		}
	}
	if err != nil {
		return err
	}
	return nil
}

// StartAsync 异步启动
func (s *SimpleScheduler) StartAsync() {
	s.scheduler.StartAsync()
}

// StartBlocking 阻塞启动
func (s *SimpleScheduler) StartBlocking() {
	s.scheduler.StartBlocking()
}

// RunByTag 根据标签名启动
func (s *SimpleScheduler) RunByTag(tag string) error {
	err := s.scheduler.RunByTag(tag)
	if err != nil {
		return err
	}
	return nil
}

// SetOnceCorn 获取对应具体日期时间的Corn表达式
func SetOnceCorn(corn *string, date time.Time) {
	carbonDate := carbon.Time2Carbon(date)
	*corn = fmt.Sprintf("%d %d %d %d %d", carbonDate.Minute(), carbonDate.Hour(), carbonDate.Day(), carbonDate.Month(), carbonDate.Year())
}
