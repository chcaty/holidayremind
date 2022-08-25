package uxnet

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

var client *http.Client
var httpClient *HttpClient
var onceClient sync.Once
var onceHttpClient sync.Once

type HttpClient struct {
	Timeout time.Duration
}

type Requester interface {
	Get(data BaseData, response *[]byte) error
	Post(data BaseData, contentType contentType, body any, response *[]byte) error
	PostByJson(data BaseData, body any, response *[]byte) error
}

// GetHttpClient 生成一个HttpClient对象
func GetHttpClient() *HttpClient {
	onceHttpClient.Do(func() {
		httpClient = &HttpClient{
			Timeout: 5 * time.Second,
		}
	})
	return httpClient
}

func (s *HttpClient) GeClient() *http.Client {
	onceClient.Do(func() {
		client = &http.Client{
			Timeout: s.Timeout,
		}
	})
	return client
}

// Get http get method
func (s *HttpClient) Get(data BaseData, response *[]byte) error {
	//new request
	req, err := http.NewRequest("GET", data.Url, nil)
	if err != nil {
		return errors.New("new request is fail ")
	}
	//add params
	setParams(data.Params, req)
	//add headers
	setHeaders(data.Headers, req)
	//http client
	client := s.GeClient()
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
func (s *HttpClient) Post(data BaseData, contentType contentType, body any, response *[]byte) error {
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
	setParams(data.Params, req)
	//add headers
	setHeaders(data.Headers, req)
	//set Content-type
	req.Header.Set("Content-type", string(contentType))
	//http client
	client := s.GeClient()
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
func (s *HttpClient) PostByJson(data BaseData, body any, response *[]byte) error {
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

func setParams(params map[string]string, req *http.Request) {
	q := req.URL.Query()
	if params != nil {
		for key, val := range params {
			q.Add(key, val)
		}
		req.URL.RawQuery = q.Encode()
	}
}

func setHeaders(headers map[string]string, req *http.Request) {
	if headers != nil {
		for key, val := range headers {
			req.Header.Add(key, val)
		}
	}
}
