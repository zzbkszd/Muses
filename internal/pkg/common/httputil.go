package common

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func CreateFormQuery(params map[string]string) string {
	ps := ""
	for k, v := range params {
		ps += fmt.Sprintf("&%s=%s", k, v)
	}
	return ps[1:]
}

func DoGet(url string, params map[string]string, headers map[string]string) (string, error) {
	reqStr := CreateFormQuery(params)
	log.Printf("do get full url: %s?%s\n", url, reqStr)
	req, err := http.NewRequest(http.MethodGet, url+"?"+reqStr, nil)
	if err != nil {
		return "", err
	}
	if headers != nil {
		for k, v := range headers {
			req.Header.Add(k, v)
		}
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
