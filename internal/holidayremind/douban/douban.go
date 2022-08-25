package douban

import (
	"holidayRemind/internal/pkg/uxjson"
	"holidayRemind/internal/pkg/uxmap"
	"holidayRemind/internal/pkg/uxnet"
)

const baseUrl = "https://m.douban.com/rexxar/api/v2/subject_collection/"

func GetWeeklyBestByType(bestType WeeklyBestType, params CollectionParams, response *CollectionResponse) error {
	paramsMap := map[string]string{}
	uxmap.GetMapByStruct(params, &paramsMap)
	err := getRequest(string(bestType)+"/items", paramsMap, response)
	if err != nil {
		return err
	}
	return nil
}

func getRequest(path string, params map[string]string, data *CollectionResponse) error {
	var err error
	var resp []byte
	header := uxnet.DefaultHeader
	header["Referer"] = "https://m.douban.com/"
	requestData := uxnet.BaseData{
		Url:     baseUrl + path,
		Headers: header,
		Params:  params,
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
