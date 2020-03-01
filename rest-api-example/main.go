package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	api_url := "https://httpbin.org/get"
	req, err := http.NewRequest(http.MethodGet, api_url, nil)
	if err != nil {
		panic(err)
	}
	client := http.Client{Timeout: 1 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	// TODO: replace ReadAll with something less dangerous
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v", string(body))
}
