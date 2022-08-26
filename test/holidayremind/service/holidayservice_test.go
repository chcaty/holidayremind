package service

import (
	"holidayRemind/internal/holidayremind/service/holidayservice"
	"testing"
)

func TestHolidayService(t *testing.T) {
	holidayservice.Start()
	t.Log("success")
}
