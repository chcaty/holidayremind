package service

import (
	"holidayRemind/internal/holidayremind/service/moyuduckservice"
	"testing"
)

func TestMoyuduckService(t *testing.T) {
	moyuduckservice.Start()
	t.Log("success")
}
