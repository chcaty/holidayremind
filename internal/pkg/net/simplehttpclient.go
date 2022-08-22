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
	Get(response *[]byte, data RequestBaseData) error
	Post(response *[]byte, data RequestBaseData, contentType ContentType, body any) error
	PostByJson(response *[]byte, data RequestBaseData, body any) error
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
func (s *SimpleHttpClient) Get(response *[]byte, data RequestBaseData) error {
	//new request
	req, err := http.NewRequest("GET", data.Url, nil)
	if err != nil {
		return errors.New("new request is fail ")
	}
	//add params
	setParams(req, data)
	//add headers
	setHeaders(req, data)
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
func (s *SimpleHttpClient) Post(response *[]byte, data RequestBaseData, contentType ContentType, body any) error {
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
	setParams(req, data)
	//add headers
	setHeaders(req, data)
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
func (s *SimpleHttpClient) PostByJson(response *[]byte, data RequestBaseData, body any) error {
	err := s.Post(response, data, Json, body)
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

func setParams(req *http.Request, data RequestBaseData) {
	q := req.URL.Query()
	if data.Params != nil {
		for key, val := range data.Params {
			q.Add(key, val)
		}
		req.URL.RawQuery = q.Encode()
	}
}

func setHeaders(req *http.Request, data RequestBaseData) {
	if data.Headers != nil {
		for key, val := range data.Headers {
			req.Header.Add(key, val)
		}
	}
}
