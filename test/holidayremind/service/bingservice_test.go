package service

import (
	"holidayRemind/internal/holidayremind/service/bingservice"
	"testing"
)

func TestBingService(t *testing.T) {
	bingservice.Start()
	t.Log("success")
}
