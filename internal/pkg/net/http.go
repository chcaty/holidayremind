package net

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Get 发送GET请求
// url:请求地址
// response:请求返回的内容
func Get(url string, response *string) {
	client := http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	defer CloseBody(resp.Body)
	if err != nil {
		panic(err)
	}
	result, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	*response = string(result)
	return
}

// Post 发送POST请求
// url:请求地址，body:POST请求提交的数据,param:POST请求url param数据,contentType:请求体格式，如：application/json,content:请求放回的内容
func Post(url string, body interface{}, param map[string]string, contentType ContentType, content *string) {
	jsonStr, _ := json.Marshal(body)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Add("content-type", string(contentType))
	if err != nil {
		panic(err)
	}
	defer CloseBody(req.Body)

	params := req.URL.Query()
	for i, v := range param {
		params.Add(i, v)
	}
	req.URL.RawQuery = params.Encode()

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer CloseBody(resp.Body)

	result, _ := io.ReadAll(resp.Body)
	*content = string(result)
	return
}

func CloseBody(body io.ReadCloser) {
	func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Printf("Request Body Close error: %s", err.Error())
		}
	}(body)
}
