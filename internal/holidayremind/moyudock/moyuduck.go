package moyudock

import (
	"fmt"
	"holidayRemind/internal/pkg"
	"holidayRemind/internal/pkg/net"
)

const baseUrl = "https://api.moyuduck.com/"

func GetHotTop(topSiteInfo *Response[HotTopSite]) error {
	err := getRequest("hot/all", topSiteInfo)
	if err != nil {
		return err
	}
	return nil
}

func GetHoliday(holidayInfo *Response[Holiday]) error {
	err := getRequest("holiday", holidayInfo)
	if err != nil {
		return err
	}
	return nil
}

func GetImage(imageInfo *Response[string]) error {
	err := getRequest("random/xiezhen", imageInfo)
	if err != nil {
		return err
	}
	return nil
}

func getRequest[T ResponseData](path string, data *T) error {
	var err error
	var resp []byte
	requestData := net.RequestBaseData{
		Url:     baseUrl + path,
		Headers: nil,
		Params:  net.DefaultHeader,
	}
	client := net.GetSimpleHttpClient()
	err = client.Get(requestData, &resp)
	if err != nil {
		return err
	}
	err = pkg.ToObjectByBytes(resp, data)
	if err != nil {
		return fmt.Errorf("get struct fail. error: %w", err)
	}
	return nil
}
