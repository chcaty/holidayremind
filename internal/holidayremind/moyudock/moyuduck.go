package moyudock

import (
	"encoding/json"
	"fmt"
	"holidayRemind/internal/pkg/net"
)

func GetHotTop(top *Response[HotTopSite]) error {
	var err error
	resp := ""
	net.Get(baseUrl+"hot/all", &resp)
	err = json.Unmarshal([]byte(resp), top)
	if err != nil {
		return fmt.Errorf("get hottop struct fail. error: %w", err)
	}
	return nil
}

func GetHoliday(holidayInfo *Response[Holiday]) error {
	var err error
	resp := ""
	net.Get(baseUrl+"Holiday", &resp)
	err = json.Unmarshal([]byte(resp), holidayInfo)
	if err != nil {
		return fmt.Errorf("get holiday struct fail. error: %w", err)
	}
	return nil
}

func GetImage(imageUrl *Response[string]) error {
	var err error
	resp := ""
	net.Get(baseUrl+"Holiday", &resp)
	err = json.Unmarshal([]byte(resp), imageUrl)
	if err != nil {
		return fmt.Errorf("get image struct fail. error: %w", err)
	}
	return nil
}
