package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Crawler(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("HTTP transport error is:", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Read io error is:", err)
	}

	fmt.Println(string(body))
	wg.Done()
}
