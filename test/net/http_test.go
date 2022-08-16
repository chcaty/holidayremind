package net

import (
	"holidayRemind/internal/pkg/net"
	"testing"
)

func TestGetUrl(t *testing.T) {
	t.Log(net.Get("https://api.moyuduck.com/random/xiezhen"))
}

func TestPostUrl(t *testing.T) {
	t.Log(net.Post("https://oapi.dingtalk.com/robot/send", nil, nil, net.Json))
}
