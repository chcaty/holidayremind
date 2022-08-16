package rss

import (
	"encoding/xml"
	"fmt"
	"holidayRemind/internal/pkg/net"
	"regexp"
	"sort"
	"strings"
)

func Request(data RequestData, rss *Rss) error {
	err := getRssInfo(rss, data.Url)
	if err != nil {
		return err
	}
	cleanRssInfo(rss, data.ChannelType)
	return nil
}

func getRssInfo(rss *Rss, url string) error {
	body := ""
	net.Get(url, &body)
	err := xml.Unmarshal([]byte(body), rss)
	if err != nil {
		return fmt.Errorf("get rss struct fail. error: %w", err)
	}
	sort.SliceStable(rss.Channel.Item, func(i, j int) bool {
		return rss.Channel.Item[i].PubDate > rss.Channel.Item[j].PubDate
	})
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
