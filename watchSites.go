package main

import (
	"net/http"
	"strconv"
	"time"
)

func watchSites() {

	// append slice with your sites
	// with quotation marks and commas
	// like this: "https://url1","https://url2
	urls := []string{}

	for {
		for i, url := range urls {

			resp, err := http.Get(url)
			if err == nil {
				if resp.StatusCode != http.StatusOK {
					// removing https from the url
					// otherwise email body will not be sent
					url := url[8:]
					msg := url + " " + "status is" + " " +
						strconv.Itoa(resp.StatusCode)
					email(msg)
					// removing url so it will not spam
					urls = append(urls[:i], urls[i+1:]...)

				}
				time.Sleep(10 * time.Second)
			}
		}
	}
}
