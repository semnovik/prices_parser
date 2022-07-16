package telegram

import (
	"awesomeProject/telegram/answers"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func readFromString(text string, subtext string) bool {
	return strings.Contains(strings.ToLower(text), subtext)
}

func GetChance() string {
	rand.Seed(time.Now().Unix())
	return strconv.Itoa(rand.Intn(100)) + "%"
}

func SetTimeForDota(message string) {
	index := strings.LastIndex(message, "в")
	answers.TimeForDota = message[index:]
}

func RoarForSquad() string {
	if answers.TimeForDota == "Время сквада не установлено" {
		return answers.TimeForDota
	} else {
		return answers.SquadCallText + answers.TimeForDota
	}
}
