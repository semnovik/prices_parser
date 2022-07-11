package parsing

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
	text := "Курс по ЦБ: " + strings.TrimSpace(dollarCurrency) + " \nВчера: " + strings.TrimSpace(buyDollar) + " \nСегодня: " + strings.TrimSpace(sellDollar) + "\n"
	return text
}

func GetCurrentCurrencyUSD() string {
	response, err := http.Get("https://quote.rbc.ru/ticker/59111")
	if err != nil {
		log.Fatal(err)
	}

	Body := response.Body

	doc, err := goquery.NewDocumentFromReader(Body)
	if err != nil {
		log.Fatal(err)
	}
	currentDollarPrice := doc.Find(".chart__info__sum").Text()
	return "Текущий курс доллара: " + strings.TrimSpace(currentDollarPrice)
}

func GetAliCurrency() string {
	response, err := http.Get("https://alicoup.ru/currency/")
	if err != nil {
		log.Fatal(err)
	}

	Body := response.Body

	doc, err := goquery.NewDocumentFromReader(Body)
	if err != nil {
		log.Fatal(err)
	}
	dollarAliTable, _ := doc.Find(".font-weight-light").Children().Next().Next().Children().Next().Html()
	aliCurrencyMessage := "Текущий курс USD на Алиэкспресс: " + strings.TrimSpace(dollarAliTable)
	return aliCurrencyMessage
}
