package doubanservice

import (
	"github.com/go-co-op/gocron"
	"holidayRemind/configs"
	"holidayRemind/internal/holidayremind/dingtalk"
	"holidayRemind/internal/holidayremind/douban"
	"holidayRemind/internal/holidayremind/scheduler"
	"log"
)

func Start() {
	var err error
	doubanScheduler := scheduler.GetSimpleScheduler()
	err = doubanScheduler.AddCornJob("30 16 * * 5", true, "movieWeeklyBest", movieWeeklyBestService)
	if err != nil {
		log.Printf("douban movie weekly best service Error:%v\n", err.Error())
		return
	}
	err = doubanScheduler.AddCornJob("30 16 * * 5", true, "chineseTvWeeklyBest", chineseTvWeeklyBestService)
	if err != nil {
		log.Printf("douban chinese tv weekly best service Error:%v\n", err.Error())
		return
	}
	doubanScheduler.StartAsync()
}

func movieWeeklyBestService(job gocron.Job) error {
	var err error
	params := douban.CollectionParams{
		Start:     0,
		Count:     6,
		ItemsOnly: 1,
		ForMobile: 1,
	}
	collection := douban.CollectionResponse{}
	err = douban.GetWeeklyBestByType(douban.Movie, params, &collection)
	if err != nil {
		return err
	}
	err = sendDingTalkMessage(collection.Items)
	if err != nil {
		return err
	}
	log.Printf("movieWeeklyBest job's last run: %s this job's next run: %s", job.LastRun(), job.NextRun())
	return nil
}

func chineseTvWeeklyBestService(job gocron.Job) error {
	var err error
	params := douban.CollectionParams{
		Start:     0,
		Count:     6,
		ItemsOnly: 1,
		ForMobile: 1,
	}
	collection := douban.CollectionResponse{}
	err = douban.GetWeeklyBestByType(douban.ChineseTV, params, &collection)
	if err != nil {
		return err
	}
	err = sendDingTalkMessage(collection.Items)
	if err != nil {
		return err
	}
	log.Printf("chineseTvWeeklyBest job's last run: %s this job's next run: %s", job.LastRun(), job.NextRun())
	return nil
}

func sendDingTalkMessage(items []douban.CollectionItem) error {
	var links []dingtalk.FeedCardLink
	for _, item := range items {
		links = append(links, dingtalk.FeedCardLink{
			Title:      item.Title,
			MessageUrl: item.Url,
			PictureUrl: item.CoverUrl,
		})
	}
	dto := dingtalk.FeedCardMessageDTO{
		Token: configs.DingTalkToken,
		Links: links,
	}
	err := dingtalk.SendFeedCardMessage(dto)
	if err != nil {
		return err
	}
	return nil
}
