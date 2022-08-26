package service

import (
	"holidayRemind/internal/holidayremind/service/vvhanservice"
	"testing"
)

func TestVvhanService(t *testing.T) {
	vvhanservice.Start()
	t.Log("success")
}
