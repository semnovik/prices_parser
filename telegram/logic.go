package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"awesomeProject/prices"
)

func GetUpdates(url string, offset int) ([]Update, error) {
	resp, err := http.Get(url + "/getUpdates" + "?offset=" + strconv.Itoa(offset))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var restResp RestResp
	err = json.Unmarshal(body, &restResp)
	if err != nil {
		return nil, err
	}
	return restResp.Result, nil
}

func Respond(botUrl string, update Update) error {
	var BotMessage BotMessage
	BotMessage.ChatId = update.Message.Chat.ChatId
	rand.Seed(time.Now().Unix())
	if update.Message.Text == "/Доллар" {
		BotMessage.Text = prices.GetPrices()
	}
	if update.Message.Text == "/Кто пидор?" {
		BotMessage.Text = "Сегодня пидор Друля"
	}
	if update.Message.Text == "/Сквад" {
		BotMessage.Text = "Внимание @Semanovik @AlexNicker @Andrey @Vyacheslov и Вован\nСегодня сквад в 22 МСК "
	}
	if update.Message.Text == "/Сема" {
		BotMessage.Text = "@Semanovik пидор на " + strconv.Itoa(rand.Intn(100)) + "%"
	}
	if update.Message.Text == "/Леха" {
		BotMessage.Text = "@AlexNicker пидор на " + strconv.Itoa(rand.Intn(100)) + "%"
	}
	if update.Message.Text == "/Друля" {
		BotMessage.Text = "@Andrey пидор на " + strconv.Itoa(rand.Intn(100)) + "%"
	}
	if update.Message.Text == "/Слава" {
		BotMessage.Text = "@Vyacheslov пидор на " + strconv.Itoa(rand.Intn(100)) + "%"
	}
	if update.Message.Text == "/Вован" {
		BotMessage.Text = "@Вован пидор по жизни"
	}
	buf, err := json.Marshal(BotMessage)
	if err != nil {
		return err
	}
	_, err = http.Post(botUrl+"/sendMessage", "application/json", bytes.NewBuffer(buf))
	if err != nil {
		return err
	}
	return nil
}

func Setup() {
	botToken := "5418352036:AAEl4mn0Q91dkms3L-73B3WTlgTcwkzz3mc"
	botApi := "https://api.telegram.org/bot"
	botUrl := botApi + botToken
	offset := 0
	for {
		updates, err := GetUpdates(botUrl, offset)
		if err != nil {
			log.Println("Something went wrong: ", err)
		}
		for _, update := range updates {
			err = Respond(botUrl, update)
			offset = update.UpdateId + 1
		}
		fmt.Println(updates)
	}
}
