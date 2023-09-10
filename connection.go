package main

import (
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func simpleRequest(Url string, host string, Proxy string, input_headers string, headers_value string, userAgent string, client *http.Client, tr *http.Transport) (string, string, string) {
	var response string
	var scode string
	var headers string

	req, err := http.NewRequest("GET", Url, nil)
	if err != nil {
		return response, scode, headers
	}
	req.Host = host
	req.Header.Set("User-Agent", userAgent)
	if input_headers != "" {
		req.Header.Set(input_headers, headers_value)

	}

	resp, err := client.Do(req)
	if err != nil {
		return response, scode, headers
	}
	defer resp.Body.Close()
	headers = headerDump(resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	scode = strconv.Itoa(resp.StatusCode)
	response = string(body)

	return response, scode, headers
}

func headerDump(headers http.Header) string {
	var sb strings.Builder
	for name, values := range headers {
		for _, value := range values {
			head := name + ": " + value + "\n"
			sb.WriteString(head)
		}
	}
	return sb.String()
}

func createHTTPClient(proxy string, redirection bool) (*http.Client, *http.Transport) {
	var tr *http.Transport

	if proxy != "null" {
		proxyURL := proxy
		proxy, _ := url.Parse(proxyURL)
		tr = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			Proxy:           http.ProxyURL(proxy),
		}
	} else {
		tr = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	}
	if redirection == true {
		client := &http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
			Transport: tr,
		}
		return client, tr

	} else {
		client := &http.Client{

			Transport: tr,
		}
		return client, tr
	}

}
