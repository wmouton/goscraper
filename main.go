package main

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"os"
	"strings"
)

func checkError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func writeFile(data, filename string) {
	file, err := os.Create(filename)
	defer func() {
		if err := file.Close(); err != nil {
			checkError(err)
		}
	}()
	checkError(err)

	if _, err = file.WriteString(data); err != nil {
		checkError(err)
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

	doc.Find("div.river").Find("div.post-block").Each(func(index int, item *goquery.Selection){
		h2 := item.Find("h2")
		title := strings.TrimSpace(h2.Text())
		url, _ := h2.Find("a").Attr("href")
		log.Println(title, url)
	})

	//log.Println(river)
	//writeFile(river,"index.html")
}
