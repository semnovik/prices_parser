package main

import (
	"fmt"

	"awesomeProject/prices"
	"awesomeProject/telegram"
)

func main() {
	price := prices.GetPrices()
	fmt.Println(price)
	telegram.Setup()
}
