package net

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

var instance *http.Client
var once sync.Once

type SimpleHttpClient struct {
}

type SimpleHttpClientRequest interface {
	Get(data RequestBaseData, response *[]byte) error
	Post(data RequestBaseData, contentType ContentType, body any, response *[]byte) error
	PostByJson(data RequestBaseData, body any, response *[]byte) error
}

func (s *SimpleHttpClient) GetHttpClient() *http.Client {
	once.Do(func() {
		instance = &http.Client{
			Timeout: 5 * time.Second,
		}
	})
	return instance
}

// Get http get method
func (s *SimpleHttpClient) Get(data RequestBaseData, response *[]byte) error {
	//new request
	req, err := http.NewRequest("GET", data.Url, nil)
	if err != nil {
		return errors.New("new request is fail ")
	}
	//add params
	setParams(data, req)
	//add headers
	setHeaders(data, req)
	//http client
	client := s.GetHttpClient()
	log.Printf("Go GET URL : %s", req.URL.String())

	//傳送請求
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	//一定要關閉res.Body
	defer closeBody(&res.Body)
	//讀取body
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	*response = resBody
	return nil
}

// Post http post method
func (s *SimpleHttpClient) Post(data RequestBaseData, contentType ContentType, body any, response *[]byte) error {
	var err error
	//add post body
	var bodyJson []byte
	if body != nil {
		bodyJson, err = json.Marshal(body)
		if err != nil {
			return errors.New("http post body to json failed")
		}
	}
	req, err := http.NewRequest("POST", data.Url, bytes.NewBuffer(bodyJson))
	if err != nil {
		return errors.New("new request is fail ")
	}
	//add params
	setParams(data, req)
	//add headers
	setHeaders(data, req)
	//set Content-type
	req.Header.Set("Content-type", string(contentType))
	//http client
	client := s.GetHttpClient()
	log.Printf("Go POST URL : %s", req.URL.String())

	//傳送請求
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	//一定要關閉res.Body
	defer closeBody(&res.Body)
	//讀取body
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	*response = resBody
	return nil
}

// PostByJson http post method with header content-type:application/json
func (s *SimpleHttpClient) PostByJson(data RequestBaseData, body any, response *[]byte) error {
	err := s.Post(data, Json, body, response)
	if err != nil {
		return err
	}
	return nil
}

// closeBody close response body
func closeBody(body *io.ReadCloser) {
	func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("Request Body Close Fail. Error: %s", err.Error())
		}
	}(*body)
}

// GetSimpleHttpClient 生成一个SimpleHttpClient对象
func GetSimpleHttpClient() *SimpleHttpClient {
	return new(SimpleHttpClient)
}

func setParams(data RequestBaseData, req *http.Request) {
	q := req.URL.Query()
	if data.Params != nil {
		for key, val := range data.Params {
			q.Add(key, val)
		}
		req.URL.RawQuery = q.Encode()
	}
}

func setHeaders(data RequestBaseData, req *http.Request) {
	if data.Headers != nil {
		for key, val := range data.Headers {
			req.Header.Add(key, val)
		}
	}
}
