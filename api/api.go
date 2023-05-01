package api

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	netUrl "net/url"
)

type Headers map[string]string
type Querys map[string]string

func genBaseHeaders(token string) Headers {
	headers := Headers{}
	headers["Authorization"] = fmt.Sprintf("token %s", token)
	return headers
}

func httpSetting(url string, headers Headers, query Querys, body io.Reader, method string) *http.Request {
	// HTTP SETTING
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		panic(err)
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	params := netUrl.Values{}
	for key, value := range query {
		params.Set(key, value)
	}
	req.URL.RawQuery = params.Encode()
	return req
}

func getData(url string, headers Headers, query Querys) []byte {
	// HTTP SETTING
	req := httpSetting(url, headers, query, nil, "GET")

	// GET DATA
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return data
}

func postData(url string, headers Headers, query Querys, body io.Reader) {
	req := httpSetting(url, headers, query, body, "POST")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println("success!")
}

func delData(url string, headers Headers, query Querys) {
	req := httpSetting(url, headers, query, nil, "DELETE")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}
