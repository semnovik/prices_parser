package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

var BotUrl, _ = os.ReadFile("url")

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

func Respond(update Update) {
	var BotMessage BotMessage
	BotMessage.ChatId = update.Message.Chat.ChatId
	messageFromChannel := update.Message.Text

	nameFromChat := update.Message.Chat.FirstName
	nameFromChannel := update.Message.FromObj.FirstName

	BotMessage.Text = ChoseAnswer(messageFromChannel, nameFromChat, nameFromChannel)

	SendMessage(BotMessage)
}

func Setup() {
	offset := 0
	for {
		updates, err := GetUpdates(string(BotUrl), offset)
		if err != nil {
			log.Println("Something went wrong: ", err)
		}
		for _, update := range updates {
			Respond(update)
			offset = update.UpdateId + 1
		}
		fmt.Println(updates)
	}
}

func SendMessage(BotMessage BotMessage) {
	buf, err := json.Marshal(BotMessage)
	if err != nil {
		log.Println("Something went wrong")
	}
	_, err = http.Post(string(BotUrl)+"/sendMessage", "application/json", bytes.NewBuffer(buf))
	if err != nil {
		log.Println("Something went wrong during sending message")
	}
}
