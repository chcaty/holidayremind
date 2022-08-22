package moyudock

import (
	"fmt"
	"holidayRemind/internal/pkg"
	"holidayRemind/internal/pkg/net"
)

const baseUrl = "https://api.moyuduck.com/"

func GetHotTop(topSiteInfo *Response[HotTopSite]) error {
	err := getRequest(topSiteInfo, "hot/all")
	if err != nil {
		return err
	}
	return nil
}

func GetHoliday(holidayInfo *Response[Holiday]) error {
	err := getRequest(holidayInfo, "holiday")
	if err != nil {
		return err
	}
	return nil
}

func GetImage(imageInfo *Response[string]) error {
	err := getRequest(imageInfo, "random/xiezhen")
	if err != nil {
		return err
	}
	return nil
}

func getRequest[T ResponseData](data *T, path string) error {
	var err error
	var resp []byte
	requestData := net.RequestBaseData{
		Url:     baseUrl + path,
		Headers: nil,
		Params:  net.DefaultHeader,
	}
	client := net.GetSimpleHttpClient()
	err = client.Get(&resp, requestData)
	if err != nil {
		return err
	}
	err = pkg.ToObjectByBytes(data, resp)
	if err != nil {
		return fmt.Errorf("get struct fail. error: %w", err)
	}
	return nil
}
