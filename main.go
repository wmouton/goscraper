package main

import (
	"log"
	"net/http"
)

func main() {
	url := "https://techcrunch.com/"
	res, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}

	if res.StatusCode > 400 {
		log.Println("Status Code: ", res.StatusCode)
	}
}
