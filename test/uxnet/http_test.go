package uxnet

import (
	"holidayRemind/internal/pkg/uxnet"
	"testing"
)

func TestGetUrl(t *testing.T) {
	var resp []byte
	err := uxnet.GetHttpClient().Get(uxnet.BaseData{Url: "https://api.moyuduck.com/random/xiezhen"}, &resp)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(string(resp))
}

func TestPostUrl(t *testing.T) {
	var resp []byte
	err := uxnet.GetHttpClient().Post(uxnet.BaseData{Url: "https://oapi.dingtalk.com/robot/send"}, uxnet.Json, nil, &resp)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(string(resp))
}
