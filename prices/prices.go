package prices

import (
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func GetPrices() string {
	response, err := http.Get("https://cbr.ru/")
	if err != nil {
		log.Fatal(err)
	}

	Body := response.Body

	doc, err := goquery.NewDocumentFromReader(Body)
	if err != nil {
		log.Fatal(err)
	}
	dollarTable := doc.Find("div.main-indicator_rate")
	dollarCurrency, _ := dollarTable.Children().Html()
	buyDollar, _ := dollarTable.Children().Next().Html()
	sellDollar, _ := dollarTable.Children().Next().Next().Html()
	text := strings.TrimSpace(dollarCurrency) + " \nBuy: " + strings.TrimSpace(buyDollar) + " \nSell: " + strings.TrimSpace(sellDollar) + "\nSlava pidor"
	return text
}
