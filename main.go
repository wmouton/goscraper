package main

import (
	"log"
	"net/http"
)

func main() {
	url := "https://techcrunch.com/"
	res, err := http.Get(url)
	defer func() {
		if err := res.Body.Close(); err != nil {
			log.Fatalln("defer failed: ", err)
		}
	}()
	if err != nil {
		log.Println(err)
	}

	if res.StatusCode > 400 {
		log.Println("Status Code: ", res.StatusCode)
	}
}
