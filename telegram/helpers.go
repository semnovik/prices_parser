package telegram

import "strings"

func readFromString(text string, subtext string) bool {
	return strings.Contains(strings.ToLower(text), subtext)
}
