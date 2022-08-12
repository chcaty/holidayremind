package rss

import (
	"encoding/xml"
	"fmt"
	"holidayRemind/internal/pkg"
	"io"
	"net/http"
	"regexp"
	"sort"
	"strings"
)

func Request(data RequestData, rss *Rss) error {
	err := getRssInfo(rss, data.Url)
	cleanRssInfo(rss, data.ChannelType)
	if err != nil {
		return err
	}
	return nil
}

func getRssInfo(rss *Rss, url string) error {
	rssReq, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	params := rssReq.URL.Query()
	rssReq.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.0.0 Safari/537.36")
	rssReq.URL.RawQuery = params.Encode()

	rssResp, err := http.DefaultClient.Do(rssReq)
	if err != nil {
		return err
	}
	if rssResp != nil {
		defer pkg.CloseBody(rssResp.Body)
		body, err := io.ReadAll(rssResp.Body)
		if err != nil {
			return fmt.Errorf("get response body fail. error: %w", err)
		}
		err = xml.Unmarshal(body, rss)
		if err != nil {
			return fmt.Errorf("get rss struct fail. error: %w", err)
		}
		sort.SliceStable(rss.Channel.Item, func(i, j int) bool {
			return rss.Channel.Item[i].PubDate > rss.Channel.Item[j].PubDate
		})
	}
	return nil
}

func cleanRssInfo(rss *Rss, channelType ChannelType) {
	switch channelType {

	case Sspai:
		cleanSspaiDescription(&rss.Channel)
	case Zztt:
		cleanZzttInfo(&rss.Channel)
	}
}

func cleanZzttInfo(channel *Channel) {
	for i := 0; i < len(channel.Item); i++ {
		if channel.Item[i].Title == "黑料不打烊—精选专区" || channel.Item[i].Title == "." || channel.Item[i].Title == ".   " {
			channel.Item = append(channel.Item[:i], channel.Item[i+1:]...)
			// form the remove item index to start iterate next item
			i--
		} else {
			reg, _ := regexp.Compile("萝莉约啪.*偷窥视频")
			channel.Item[i].Description = reg.ReplaceAllString(channel.Item[i].Description, "")
		}
	}
}

func cleanSspaiDescription(channel *Channel) {
	for i := 0; i < len(channel.Item); i++ {
		reg := regexp.MustCompile(`.*<a`)
		results := reg.FindAllString((*channel).Item[i].Description, -1)
		if results != nil {
			(*channel).Item[i].Description = reg.FindAllString((*channel).Item[i].Description, -1)[0]
			(*channel).Item[i].Description = strings.Replace((*channel).Item[i].Description, "<a", "", -1)
		}
	}
}
