package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	res, err := http.Get("https://api.weatherapi.com/v1/current.json?key=b3f8320ed9d842de890213957250907&q=London&aqi=no")
	if err != nil {
		panic(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(res.Body)

	if res.StatusCode != 200 {
		panic(res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	fmt.Println(string(body))
}
