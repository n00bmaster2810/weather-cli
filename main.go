package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"log"
	"net/http"
	"os"
)

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
}

func main() {
	loadEnv()

	apiKey := os.Getenv("API_KEY")
	cityName := os.Getenv("CITY_NAME")

	if len(os.Args) >= 2 {
		cityName = os.Args[1]
	}

	res, err := http.Get("https://api.weatherapi.com/v1/current.json?key=" + apiKey + "&q=" + cityName + "&aqi=no")
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
