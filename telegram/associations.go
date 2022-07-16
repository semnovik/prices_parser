package telegram

import (
	"awesomeProject/parsing"
	"awesomeProject/telegram/answers"
	"math/rand"
	"strconv"
	"time"
)

var Heroes = []string{"Abaddon", "Alchemist", "Axe", "Beastmaster", "Brewmaster", "Bristleback", "Centaur Warrunner", "Chaos Knight", "Clockwerk", "Dawnbreaker", "Doom", "Dragon Knight", "Earth Spirit, лох", "Earthshaker", "Elder Titan", "Huskar", "Io, неудачник", "Kunkka", "Legion Commander", "Lifestealer", "Lycan", "Magnus", "Marci", "Mars", "Night Stalker", "Omniknight", "Phoenix", "Primal Beast", "Pudge", "Sand King", "Slardar", "Snapfire", "Spirit Breaker", "Sven", "Tidehunter", "Timbersaw", "Tiny", "Treant Protector", "Tusk", "Underlord", "Undying", "Wraith King", "Anti-Mage", "Arc Warden, лошара", "Bloodseeker", "Bounty Hunter", "Broodmother, паучок вонючий", "Clinkz", "Drow Ranger", "Ember Spirit", "Faceless Void", "Gyrocopter", "Hoodwink", "Juggernaut, настало время ебашить", "Lone Druid, гг мы проебали", "Luna", "Medusa", "Meepo, трону пизда", "Mirana", "Monkey King", "Morphling", "Naga Siren", "Nyx Assassin", "Pangolier", "Phantom Assassin", "Phantom Lancer", "Razor", "Riki", "Shadow Fiend", "Slark", "Sniper", "Spectre", "Templar Assassin", "Terrorblade", "Troll Warlord", "Ursa", "Vengeful Spirit", "Venomancer", "Viper", "Weaver", "Ancient Apparition", "Bane", "Batrider", "Chen, можно ливать", "Crystal Maiden", "Dark Seer", "Dark Willow", "Dazzle, на нем ебашит только Сема", "Death Prophet", "Disruptor", "Enchantress", "Enigma", "Grimstroke", "Invoker", "Jakiro", "Keeper of the Light", "Leshrac", "Lich", "Lina", "Lion", "Nature’s Prophet", "Necrophos", "Ogre Magi", "Oracle", "Outworld Devourer", "Puck", "Pugna", "Queen of Pain", "Rubick", "Shadow Demon", "Shadow Shaman", "Silencer", "Skywrath Mage", "Storm Spirit", "Techies, все в минах нахуй", "Tinker", "Visage", "Void Spirit", "Warlock", "Windranger", "Winter Wyvern", "Witch Doctor", "Zeus"}

func WhoToPlayFor() string {
	rand.Seed(time.Now().Unix())
	chanceToWin := rand.Intn(100)
	empty := ""
	switch {
	case chanceToWin < 20:
		empty = "\nГГ мы проебали"
	case chanceToWin > 90:
		empty = "\nИзи разъебыч"
	}
	return Heroes[rand.Intn(123)] + "\nШанс на победу " + strconv.Itoa(chanceToWin) + "%" + empty
}

func ChoseAnswer(message string, nameFromChat string, nameFromChannel string) string {
	response := ""
	switch {
	case readFromString(message, "дота в "):
		SetTimeForDota(message)
	case readFromString(message, "курс"):
		response = parsing.GetCurrentCurrencyUSD() + "\n" + parsing.GetAliCurrency()
	case readFromString(message, "кто пидор"):
		response = "Сегодня пидор Друля"
	case readFromString(message, "сквад"):
		response = RoarForSquad()
	case readFromString(message, "сема"):
		name := nameFromChannel
		response = "@Semanovik пидор на " + GetChance() + "\n" + name + " пидор на все 100%"
	case readFromString(message, "леха"):
		response = "@AlexNicker пидор на " + GetChance()
	case readFromString(message, "друля"):
		response = "@Andrey пидор на " + GetChance()
	case readFromString(message, "слава"):
		response = "@Vyacheslov пидор на " + GetChance()
	case readFromString(message, "вован"):
		response = "@Gulyanda пидор по жизни"
	case readFromString(message, "как меня зовут"):
		name := ""
		if name = nameFromChat; len(name) < 1 {
			name = nameFromChannel
		}
		response = answers.YourNameIsText + name
	case readFromString(message, "сосногорск"):
		response = parsing.GetWeatherSosnogorsk()
	case readFromString(message, "калининград"):
		response = parsing.GetWeatherKaliningrad()
	case readFromString(message, "на ком катать"):
		heroToPlay := WhoToPlayFor()
		response = answers.TodayYouPlayAtText + heroToPlay

	case readFromString(message, "бот помощь"):
		response = answers.WhatCanIDoText

	}
	return response
}
