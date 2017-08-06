package main

import (
	"net/http"
	"strconv"
	"time"
)

func watchSites() {
	for {

		// append slice with your sites
		// with quotation marks and commas
		// like this: "https://url1","https://url2
		urls := []string{}

		for _, url := range urls {

			resp, err := http.Get(url)
			if err == nil {
				if resp.StatusCode != http.StatusOK {
					// removing https from the url
					// otherwise email body will not be sent
					url := url[8:]
					msg := url + " " + "status is" + " " +
					strconv.Itoa(resp.StatusCode)
					email(msg)
				}
				time.Sleep(10 * time.Second)
			}
		}
	}
}
