package scheduler

import (
	"fmt"
	"github.com/go-co-op/gocron"
	"github.com/golang-module/carbon"
	"sync"
	"time"
)

var timezone, _ = time.LoadLocation("Asia/Shanghai")
var instance *gocron.Scheduler
var once sync.Once

type SimpleScheduler struct {
}

type SimpleSchedulerFunc interface {
	AddCornJob(corn string, jobFun interface{}, params ...interface{}) error
	AddRandomJob(data RandomData, jobFun interface{}, params ...interface{}) error
	StartAsync()
	StartBlocking()
}

// GetScheduler 获得单例任务调度器
func (s *SimpleScheduler) GetScheduler() *gocron.Scheduler {
	once.Do(func() {
		instance = gocron.NewScheduler(timezone)
	})
	return instance
}

// AddCornJob 增加一个由Corn表达式决定何时触发的定时任务
func (s *SimpleScheduler) AddCornJob(corn string, jobFun interface{}, params ...interface{}) error {
	var err error
	_, err = s.GetScheduler().Cron(corn).Do(jobFun, params...)
	if err != nil {
		return err
	}
	return nil
}

// AddRandomJob 增加一个一定范围内随机触发的定时任务
func (s *SimpleScheduler) AddRandomJob(data RandomData, jobFun interface{}, params ...interface{}) error {
	var err error
	switch data.Unit {
	case Day:
		_, err = s.GetScheduler().EveryRandom(data.Lower, data.Upper).Days().Do(jobFun, params...)
	case Hour:
		_, err = s.GetScheduler().EveryRandom(data.Lower, data.Upper).Hours().Do(jobFun, params...)
	case Minute:
		_, err = s.GetScheduler().EveryRandom(data.Lower, data.Upper).Minutes().Do(jobFun, params...)
	case Second:
		_, err = s.GetScheduler().EveryRandom(data.Lower, data.Upper).Seconds().Do(jobFun, params...)
	}
	if err != nil {
		return err
	}
	return nil
}

// StartAsync 异步启动
func (s *SimpleScheduler) StartAsync() {
	s.GetScheduler().StartAsync()
}

// StartBlocking 阻塞启动
func (s *SimpleScheduler) StartBlocking() {
	s.GetScheduler().StartBlocking()
}

// SetOnceCorn 获取对应具体日期时间的Corn表达式
func SetOnceCorn(corn *string, date time.Time) {
	carbonDate := carbon.Time2Carbon(date)
	*corn = fmt.Sprintf("%d %d %d %d %d", carbonDate.Minute(), carbonDate.Hour(), carbonDate.Day(), carbonDate.Month(), carbonDate.Year())
}

// GetSimpleScheduler 生成一个SimpleScheduler对象
func GetSimpleScheduler() *SimpleScheduler {
	return new(SimpleScheduler)
}
