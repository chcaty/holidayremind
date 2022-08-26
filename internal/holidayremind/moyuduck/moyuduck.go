package moyuduck

import (
	"holidayRemind/internal/pkg/uxjson"
	"holidayRemind/internal/pkg/uxnet"
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
	requestData := uxnet.BaseData{
		Url:     baseUrl + path,
		Headers: uxnet.DefaultHeader,
		Params:  nil,
	}
	client := uxnet.GetHttpClient()
	err = client.Get(requestData, &resp)
	if err != nil {
		return err
	}
	err = uxjson.ToObjectByBytes(resp, data)
	if err != nil {
		return err
	}
	return nil
}
