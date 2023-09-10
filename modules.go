package main

import (
	"fmt"
	"net/http"
	"strings"
)

var vulnUrlsList = []string{}

func matcher(url string, proxy string, client *http.Client, tr *http.Transport, mt string, ms string) {
	hostName := strings.Split(url, "/")[2]

	userAgent := getRandomUserAgent()

	body, scode, headers := simpleRequest(url, hostName, proxy, "", "", userAgent, client, tr)
	//fmt.Println(headers)
	if mt == "b" {
		if xMatch(strings.ToLower(ms), strings.ToLower(body)) == true {
			fmt.Println(url)

		}
	}
	if mt == "s" {
		if xMatch(strings.ToLower(ms), strings.ToLower(scode)) == true {
			fmt.Println(url)
		}
	}
	if mt == "h" {
		if xMatch(strings.ToLower(ms), strings.ToLower(headers)) == true {
			fmt.Println(url)
		}
	}

}
