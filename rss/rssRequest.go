package rss

import (
	"encoding/xml"
	"fmt"
	"holidayRemind/common"
	"holidayRemind/dingtalk"
	"io"
	"net/http"
	"regexp"
	"strings"
)

func SspaiRssRequest(token string) {
	rss := common.Rss{}
	err := getRssInfo("https://sspai.com/feed", &rss, Sspai)
	if err != nil {
		fmt.Printf("getRssInfo Error:%v\n", err.Error())
		return
	}
	sendRssMessage(&rss.Channel, token)
}

func ZzttRssRequest(token string) {
	rss := common.Rss{}
	err := getRssInfo("https://zztt15.com/feed", &rss, Zztt)
	if err != nil {
		fmt.Printf("getRssInfo Error:%v\n", err.Error())
		return
	}
	sendRssMessage(&rss.Channel, token)
}

func getRssInfo(url string, rss *common.Rss, channelType Channel) error {
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
		channel := rss.Channel
		for i := 0; i < len(channel.Item); i++ {
			if channelType == Sspai {
				cleanSspaiDescription(&channel.Item[i].Description)
			}
			if channelType == Zztt {
				cleanZzttInfo(&channel.Item[i], &channel, i)
			}
		}
	}
	return nil
}

func cleanZzttInfo(item *common.Item, channel *common.Channel, i int) {
	if item.Title == "黑料不打烊—精选专区" || item.Title == "." {
		channel.Item = append(channel.Item[:i], channel.Item[i+1:]...)
		i-- // form the remove item index to start iterate next item
	} else {
		item.Description = strings.Replace(item.Description, "萝莉约啪 成人抖阴 色漫肉番学生援交 B站约炮 偷窥视频", "", -1)
	}
}

func sendRssMessage(channel *common.Channel, token string) {
	content := ""
	for _, item := range channel.Item {
		content += fmt.Sprintf(ReqRssContent, item.Title, item.Link, item.Description)
	}
	text := fmt.Sprintf(ReqRssMD, channel.Title, content)
	message := &common.Message{
		Title:    channel.Title + "今日推送",
		Text:     text,
		Token:    token,
		Tel:      "",
		IsRemind: false,
	}
	_, err := dingtalk.SendMessage(*message)
	if err != nil {
		fmt.Printf("sendMessage Error:%v\n", err.Error())
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
