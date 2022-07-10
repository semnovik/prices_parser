package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strings"
)

func main() {
	response, err := http.Get("https://cbr.ru/")
	if err != nil {
		log.Fatal(err)
	}

	Body := response.Body

	doc, err := goquery.NewDocumentFromReader(Body)
	if err != nil {
		log.Fatal(err)
	}
	dollatTable := doc.Find("div.main-indicator_rate")
	dollarCurrency, _ := dollatTable.Children().Html()
	buyDollar, _ := dollatTable.Children().Next().Html()
	sellDollar, _ := dollatTable.Children().Next().Next().Html()
	fmt.Println(dollarCurrency, strings.TrimSpace(buyDollar), strings.TrimSpace(sellDollar))

}
