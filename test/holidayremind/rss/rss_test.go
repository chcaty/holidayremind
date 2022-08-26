package rss

import (
	"holidayRemind/internal/pkg/rss"
	"testing"
)

func TestRssRequest(t *testing.T) {
	rssInfo := rss.Rss{}
	rssData := rss.RequestData{
		Url:         "https://sspai.com/feed",
		ChannelType: rss.Sspai,
	}
	err := rss.Request(rssData, &rssInfo)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(rssInfo)
}
