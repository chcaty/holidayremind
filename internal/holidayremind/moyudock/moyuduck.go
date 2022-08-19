package moyudock

import (
	"encoding/json"
	"fmt"
	"holidayRemind/internal/pkg/net"
)

func GetHotTop(top *Response[HotTopSite]) error {
	var err error
	var resp []byte
	requestData := net.RequestBaseData{
		Url:     baseUrl + "hot/all",
		Headers: nil,
		Params:  net.DefaultHeader,
	}
	err = net.Get(&resp, requestData)
	if err != nil {
		return err
	}
	err = json.Unmarshal(resp, top)
	if err != nil {
		return fmt.Errorf("get hottop struct fail. error: %w", err)
	}
	return nil
}

func GetHoliday(holidayInfo *Response[Holiday]) error {
	var err error
	var resp []byte
	requestData := net.RequestBaseData{
		Url:     baseUrl + "holiday",
		Headers: nil,
		Params:  net.DefaultHeader,
	}
	err = net.Get(&resp, requestData)
	if err != nil {
		return err
	}
	err = json.Unmarshal(resp, holidayInfo)
	if err != nil {
		return fmt.Errorf("get holiday struct fail. error: %w", err)
	}
	return nil
}

func GetImage(imageUrl *Response[string]) error {
	var err error
	var resp []byte
	requestData := net.RequestBaseData{
		Url:     baseUrl + "random/xiezhen",
		Headers: nil,
		Params:  net.DefaultHeader,
	}
	err = net.Get(&resp, requestData)
	if err != nil {
		return err
	}
	err = json.Unmarshal(resp, imageUrl)
	if err != nil {
		return fmt.Errorf("get image struct fail. error: %w", err)
	}
	return nil
}
