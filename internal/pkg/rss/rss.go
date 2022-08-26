package rss

import (
	"encoding/xml"
	"fmt"
	"holidayRemind/internal/pkg/uxnet"
	"sort"
)

func Request(url string, rss *Rss) error {
	err := getRssInfo(rss, url)
	if err != nil {
		return err
	}
	return nil
}

func getRssInfo(rss *Rss, url string) error {
	var err error
	var resp []byte
	requestData := uxnet.BaseData{
		Url:     url,
		Headers: uxnet.DefaultHeader,
		Params:  nil,
	}
	client := uxnet.GetHttpClient()
	err = client.Get(requestData, &resp)
	if err != nil {
		return err
	}
	err = xml.Unmarshal(resp, rss)
	if err != nil {
		return fmt.Errorf("get rss struct fail. error: %w", err)
	}
	sort.SliceStable(rss.Channel.Item, func(i, j int) bool {
		return rss.Channel.Item[i].PubDate > rss.Channel.Item[j].PubDate
	})
	return nil
}
