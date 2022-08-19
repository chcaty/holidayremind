package net

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

// Get http get method
func Get(response *[]byte, data RequestBaseData) error {
	//new request
	req, err := http.NewRequest("GET", data.Url, nil)
	if err != nil {
		return errors.New("new request is fail ")
	}
	//add params
	q := req.URL.Query()
	if data.Params != nil {
		for key, val := range data.Params {
			q.Add(key, val)
		}
		req.URL.RawQuery = q.Encode()
	}
	//add headers
	if data.Headers != nil {
		for key, val := range data.Headers {
			req.Header.Add(key, val)
		}
	}
	//http client
	client := &http.Client{}
	log.Printf("Go GET URL : %s \n", req.URL.String())

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
func Post(response *[]byte, data RequestBaseData, contentType ContentType, body any) error {
	//add post body
	var bodyJson []byte
	var req *http.Request
	if body != nil {
		var err error
		bodyJson, err = json.Marshal(body)
		if err != nil {
			return errors.New("http post body to json failed")
		}
	}
	req, err := http.NewRequest("POST", data.Url, bytes.NewBuffer(bodyJson))
	if err != nil {
		return errors.New("new request is fail: %v \n")
	}
	//add params
	q := req.URL.Query()
	if data.Params != nil {
		for key, val := range data.Params {
			q.Add(key, val)
		}
		req.URL.RawQuery = q.Encode()
	}
	//add headers
	if data.Headers != nil {
		for key, val := range data.Headers {
			req.Header.Add(key, val)
		}
	}
	//set Content-type
	req.Header.Set("Content-type", string(contentType))
	//http client
	client := &http.Client{}
	log.Printf("Go POST URL : %s \n", req.URL.String())

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
func PostByJson(response *[]byte, data RequestBaseData, body any) error {
	err := Post(response, data, Json, body)
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
			fmt.Printf("Request Body Close error: %s", err.Error())
		}
	}(*body)
}
