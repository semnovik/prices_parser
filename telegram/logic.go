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

	"awesomeProject/parsing"
)

var counter = 0

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
	messageFromChannel := update.Message.Text
	rand.Seed(time.Now().Unix())
	switch {
	case readFromString(messageFromChannel, "доллар"):
		BotMessage.Text = parsing.GetPrices() + "\n" + parsing.GetCurrentCurrencyUSD() + "\n" + parsing.GetAliCurrency()
	case readFromString(messageFromChannel, "кто пидор"):
		BotMessage.Text = "Сегодня пидор Друля"
	case readFromString(messageFromChannel, "сквад"):
		BotMessage.Text = "Внимание @Semanovik @AlexNicker @Andrey @Vyacheslov и Вован\nСегодня сквад в 22 МСК "
	case readFromString(messageFromChannel, "сема"):
		name := update.Message.Chat.FirstName
		BotMessage.Text = "@Semanovik пидор на " + strconv.Itoa(rand.Intn(100)) + "%" + "\n" + name + " пидор на все 100%"
	case readFromString(messageFromChannel, "леха"):
		BotMessage.Text = "@AlexNicker пидор на " + strconv.Itoa(rand.Intn(100)) + "%"
	case readFromString(messageFromChannel, "друля"):
		BotMessage.Text = "@Andrey пидор на " + strconv.Itoa(rand.Intn(100)) + "%"
	case readFromString(messageFromChannel, "слава"):
		BotMessage.Text = "@Vyacheslov пидор на " + strconv.Itoa(rand.Intn(100)) + "%"
	case readFromString(messageFromChannel, "вован"):
		BotMessage.Text = "@Gulyanda пидор по жизни"
	case readFromString(messageFromChannel, "как меня зовут"):
		name := ""
		if name = update.Message.Chat.FirstName; len(name) < 1 {
			name = update.Message.FromObj.FirstName
		}
		BotMessage.Text = "Тебя зовут " + name
	case readFromString(messageFromChannel, "сосногорск"):
		BotMessage.Text = parsing.GetWeatherSosnogorsk()
	case readFromString(messageFromChannel, "калининград"):
		BotMessage.Text = parsing.GetWeatherKaliningrad()
	case readFromString(messageFromChannel, "на ком катать"):
		heroToPlay := WhoToPlayFor()
		BotMessage.Text = "Ты сегодня ебашишь на " + heroToPlay
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
