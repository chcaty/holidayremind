package service

import (
	"holidayRemind/internal/holidayremind/service/rssservice"
	"testing"
)

func TestRssService(t *testing.T) {
	rssservice.Start()
	t.Log("success")
}
