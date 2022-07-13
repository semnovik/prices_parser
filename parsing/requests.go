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
	responseDollar, err := http.Get("https://quote.rbc.ru/ticker/59111")
	if err != nil {
		log.Fatal(err)
	}

	BodyDollar := responseDollar.Body

	docDollar, err := goquery.NewDocumentFromReader(BodyDollar)
	if err != nil {
		log.Fatal(err)
	}
	currentDollarPrice := docDollar.Find(".chart__info__sum").Text()

	responseEuro, err := http.Get("https://quote.rbc.ru/ticker/59090")
	if err != nil {
		log.Fatal(err)
	}

	BodyEuro := responseEuro.Body

	docEuro, err := goquery.NewDocumentFromReader(BodyEuro)
	if err != nil {
		log.Fatal(err)
	}
	euroPrice := docEuro.Find(".chart__info__sum").Text()

	return "Текущий курс доллара: " + strings.TrimSpace(currentDollarPrice) + "\n" + "Текущий курс евро: " + strings.TrimSpace(euroPrice)
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

func GetWeatherSosnogorsk() string {
	response, err := http.Get("https://weather.rambler.ru/v-sosnogorske/")
	if err != nil {
		log.Fatal(err)
	}

	Body := response.Body

	doc, err := goquery.NewDocumentFromReader(Body)
	if err != nil {
		log.Fatal(err)
	}

	table := doc.Find(".J8Tp")
	where, _ := table.Children().Children().Children().Children().Html() // Погода в Сосногорске
	when, _ := table.Children().Next().Children().Html()                 // День недели, дата
	howIt, _ := doc.Find(".TWnE").Html()
	degree, _ := doc.Find(".T8o8").Children().Children().Children().Children().Next().Html()

	return where + ":" + "\n" + when + "\n" + "\n" + howIt + "\n" + "Температура: " + strings.ReplaceAll(degree, "<!-- -->", "")
}
func GetWeatherKaliningrad() string {
	response, err := http.Get("https://weather.rambler.ru/v-kaliningrade/")
	if err != nil {
		log.Fatal(err)
	}

	Body := response.Body

	doc, err := goquery.NewDocumentFromReader(Body)
	if err != nil {
		log.Fatal(err)
	}

	table := doc.Find(".J8Tp")
	where, _ := table.Children().Children().Children().Children().Html() // Погода в Сосногорске
	when, _ := table.Children().Next().Children().Html()                 // День недели, дата
	howIt, _ := doc.Find(".TWnE").Html()
	degree, _ := doc.Find(".T8o8").Children().Children().Children().Children().Next().Html()

	return where + ":" + "\n" + when + "\n" + "\n" + howIt + "\n" + "Температура: " + strings.ReplaceAll(degree, "<!-- -->", "")
}
