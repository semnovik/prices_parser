package telegram

import (
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
