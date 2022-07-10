package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

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
	BotMessage.Text = prices.GetPrices()
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
