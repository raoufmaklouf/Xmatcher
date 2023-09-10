package main

import (
	"math/rand"
	"regexp"
	"time"
)

func randomString(length int) string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano())

	b := make([]byte, length)

	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}

	return string(b)
}

func randomNumber(n int) int {
	rand.Seed(time.Now().UnixNano())

	// Generate the minimum and maximum possible values
	min := int64(1)
	max := int64(1)
	for i := 0; i < n; i++ {
		max *= 10
	}

	// Define the blacklist
	blacklist := []int{80, 443, 4, 8}

	// Remove blacklisted numbers
	for contains(blacklist, 80) || contains(blacklist, 443) {
		randomNumber := rand.Int63n(max-min+1) + min
		if !contains(blacklist, int(randomNumber)) {
			return int(randomNumber)
		}
	}

	return 0 // Default value if all numbers are blacklisted
}

func contains(slice []int, number int) bool {
	for _, item := range slice {
		if item == number {
			return true
		}
	}
	return false
}

func xMatch(rg string, str string) bool {
	match, _ := regexp.MatchString(rg, str)
	if match == true {
		return true
	} else {
		return false

	}

}

func getRandomUserAgent() string {
	userAgents := []string{
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.150 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:87.0) Gecko/20100101 Firefox/87.0",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.150 Safari/537.36",
		"Mozilla/5.0 (SymbianOS/9.2; U; Series60/3.1 NokiaE51-1/220.34.37; Profile/MIDP-2.0 Configuration/CLDC-1.1) AppleWebKit/413 (KHTML, like Gecko) Safari/413",
		"Mozilla/5.0 (SymbianOS/9.2; U; Series60/3.1 NokiaN95/11.0.026; Profile MIDP-2.0 Configuration/CLDC-1.1) AppleWebKit/413 (KHTML, like Gecko) Safari/413",
		"Mozilla/5.0 (SymbianOS/9.3; U; Series60/3.2 Nokia5630d-1/012.020; Profile MIDP-2.1 Configuration/CLDC-1.1) AppleWebKit/413 (KHTML, like Gecko) Safari/413",
		"Mozilla/5.0 (SymbianOS/9.3; U; Series60/3.2 NokiaN78-1/12.046; Profile/MIDP-2.1 Configuration/CLDC-1.1) AppleWebKit/413 (KHTML, like Gecko) Safari/413",
		"Mozilla/5.0 (SymbianOS/9.4; Series60/5.0 NokiaN97-3/21.2.045; Profile/MIDP-2.1 Configuration/CLDC-1.1;) AppleWebKit/525 (KHTML, like Gecko) BrowserNG/7.1.4",
		"Mozilla/5.0 (SymbianOS/9.4; U; Series60/5.0 Nokia5530c-2/10.0.050; Profile MIDP-2.1 Configuration/CLDC-1.1) AppleWebKit/525 (KHTML, like Gecko) Safari/525",
		"Mozilla/5.0 (SymbianOS/9.4; U; Series60/5.0 Nokia5800d-1/31.0.101; Profile MIDP-2.1 Configuration/CLDC-1.1) AppleWebKit/413 (KHTML, like Gecko) Safari/413",
		"Mozilla/5.0 (webOS/1.4.0; U; en-US) AppleWebKit/532.2 (KHTML, like Gecko) Version/1.0 Safari/532.2 Pre/1.0",
		"Mozilla/5.0 (X11; Linux i686; U;rv: 1.7.13) Gecko/20070322 Kazehakase/0.4.4.1",
		"Mozilla/5.0 (X11; U; Linux 2.4.2-2 i586; en-US; m18) Gecko/20010131 Netscape6/6.01",
		"Mozilla/5.0 (X11; U; Linux armv7l; en-GB; rv:1.9.2b6pre) Gecko/20100318 Firefox/3.5 Maemo Browser 1.7.4.8 RX-51 N900",
		"Mozilla/5.0 (X11; U; Linux i686; de-AT; rv:1.8.0.2) Gecko/20060309 SeaMonkey/1.0",
		"Mozilla/5.0 (X11; U; Linux i686; en-GB; rv:1.7.6) Gecko/20050405 Epiphany/1.6.1 (Ubuntu) (Ubuntu package 1.0.2)",
		"Mozilla/5.0 (X11; U; Linux i686; en-US; Nautilus/1.0Final) Gecko/20020408",
		"Mozilla/5.0 (X11; U; Linux i686; en-US; rv:0.9.3) Gecko/20010801",
		"Mozilla/5.0 (X11; U; Linux i686; en-US; rv:1.2b) Gecko/20021007 Phoenix/0.3",
	}

	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(userAgents))

	return userAgents[randomIndex]
}

func isUrl(url string) bool {
	s := false
	regex1, _ := regexp.MatchString("http", url)
	regex2, _ := regexp.MatchString("://", url)
	if regex1 == true && regex2 == true {

		s = true
	}
	return s
}

func containsList(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
			break
		}
	}
	return false
}

var Headers = []string{
	"Forwarded",
	"HTTP_FORWARDED",
	"HTTP_FORWARDED_FOR",
	"HTTP_X_FORWARDED",
	"HTTP_X_FORWARDED_FOR",
	"X-Forward-For",
	"X-Forwarded",
	"X-Forwarded-By",
	"X-Forwarded-For",
	"X-Forwarded-Host",
	"X-Forwarded-Server",
	"X-Forwarder-For",
	"X-Host",
	"Origin",
	"X-Original-URL",
	"Referer",
}
