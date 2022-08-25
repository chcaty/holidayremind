package service

import (
	"holidayRemind/internal/holidayremind/service/doubanservice"
	"testing"
)

func TestDoubanService(t *testing.T) {
	doubanservice.Start()
	t.Log("success")
}
