package main

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

func checkError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

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
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkError(err)

	river := doc.Find("div.river").Size()

	log.Println(river)
}
