package rss

import (
	"encoding/xml"
	"fmt"
	"holidayRemind/common"
	"io"
	"net/http"
	"regexp"
	"sort"
	"strings"
)

func Request(configs []RequestConfig, token string) {
	rss := Rss{}
	for _, config := range configs {
		err := getRssInfo(config.Url+"/feed", &rss, Sspai)
		if err != nil {
			fmt.Printf("getRssInfo Error:%v\n", err.Error())
			return
		}
		if config.IsDingTalk {
			sendRssMessage(&rss.Channel, token)
		}
		if config.IsEmail {
			sendRssEmail(&rss.Channel)
		}
	}
}

func getRssInfo(url string, rss *Rss, channelType ChannelType) error {
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
		defer common.CloseBody(rssResp.Body)
		body, err := io.ReadAll(rssResp.Body)
		if err != nil {
			return err
		}
		err = xml.Unmarshal(body, rss)
		if err != nil {
			return err
		}
		sort.SliceStable(rss.Channel.Item, func(i, j int) bool {
			return rss.Channel.Item[i].PubDate > rss.Channel.Item[j].PubDate
		})
		for i := 0; i < len(rss.Channel.Item); i++ {
			switch channelType {
			case Sspai:
				cleanSspaiDescription(&rss.Channel.Item[i].Description)
			case Zztt:
				cleanZzttInfo(&rss.Channel.Item[i], &rss.Channel, &i)
			}
		}
	}
	return nil
}

func cleanZzttInfo(item *Item, channel *Channel, i *int) {
	if item.Title == "黑料不打烊—精选专区" || item.Title == "." || item.Title == ".   " {
		channel.Item = append(channel.Item[:*i], channel.Item[*i+1:]...)
		// form the remove item index to start iterate next item
		*i--
	} else {
		reg, _ := regexp.Compile("萝莉约啪.*偷窥视频")
		item.Description = reg.ReplaceAllString(item.Description, "")
	}
}

func cleanSspaiDescription(description *string) {
	reg := regexp.MustCompile(`.*<a`)
	results := reg.FindAllString(*description, -1)
	if results != nil {
		*description = reg.FindAllString(*description, -1)[0]
		*description = strings.Replace(*description, "<a", "", -1)
	}
}
