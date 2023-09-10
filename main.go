package main

import (
	"bufio"
	"flag"
	"os"
	"sync"
)

var wg sync.WaitGroup

func main() {

	proxy := flag.String("proxy", "null", "proxy exemple: http://127.0.0.1:8080")
	matchitem := flag.String("mt", "null", "match item [b = body / s = scode / h = headers]")
	str := flag.String("ms", "null", "match string ")
	flag.Parse()
	var client, tr = createHTTPClient(*proxy, true)
	defer client.CloseIdleConnections()
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {

		line := scanner.Text()
		if isUrl(line) == true {
			wg.Add(1)

			go func() {
				defer wg.Done()
				matcher(line, *proxy, client, tr, *matchitem, *str)
			}()

		}

	}
	wg.Wait()

}
