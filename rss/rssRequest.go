package rss

import (
	"encoding/xml"
	"fmt"
	"holidayRemind/common"
	"holidayRemind/holiday"
	"io"
	"net/http"
	"regexp"
	"strings"
)

func SendRssRequest(token string) error {
	rssReq, err := http.NewRequest("GET",
		"https://sspai.com/feed", nil)
	if err != nil {
		return err
	}
	rssReq.Header.Set("Content-Type", "application/xml; charset=UTF-8")

	params := rssReq.URL.Query()
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
		//fmt.Printf("Response Body: %+v\n", string(body))
		sspaiRss := &common.Rss{}
		err = xml.Unmarshal(body, sspaiRss)
		if err != nil {
			return err
		}
		//fmt.Printf("rss: %v\n", sspaiRss)
		sspaiChannel := sspaiRss.Channel
		content := ""
		for _, item := range sspaiChannel.Item {
			getDescription(&item.Description)
			content += fmt.Sprintf(ReqRssContent, item.Title, item.Link, item.Description)
		}
		text := fmt.Sprintf(ReqRssMD, sspaiChannel.Title, content)
		message := &common.Message{
			Title:    sspaiChannel.Title + "今日推送",
			Text:     text,
			Token:    token,
			Tel:      "",
			IsRemind: false,
		}
		//fmt.Printf("Message %v\n", *message)
		_, err = holiday.CommonSendMessage(*message)
		if err != nil {
			fmt.Printf("sendMessage Error:%v\n", err.Error())
			return err
		}
	}
	return nil
}

func getDescription(description *string) {
	reg := regexp.MustCompile(`.*<a`)
	*description = reg.FindAllString(*description, -1)[0]
	*description = strings.Replace(*description, "<a", "", -1)
}
