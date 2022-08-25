package net

import (
	"holidayRemind/internal/pkg/uxnet"
	"testing"
)

func TestGetUrl(t *testing.T) {
	var resp []byte
	t.Log(uxnet.GetHttpClient().Get(uxnet.BaseData{Url: "https://api.moyuduck.com/random/xiezhen"}, &resp))
}

func TestPostUrl(t *testing.T) {
	var resp []byte
	t.Log(uxnet.GetHttpClient().Post(uxnet.BaseData{Url: "https://oapi.dingtalk.com/robot/send"}, uxnet.Json, nil, &resp))
}
