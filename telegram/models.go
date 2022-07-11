package telegram

type Update struct {
	UpdateId int     `json:"update_id"`
	Message  Message `json:"message"`
}

type Message struct {
	Chat    Chat   `json:"chat"`
	Text    string `json:"text"`
	FromObj From   `json:"from"`
}

type Chat struct {
	ChatId    int    `json:"id"`
	FirstName string `json:"first_name"`
}

type RestResp struct {
	Result []Update `json:"result"`
}

type BotMessage struct {
	ChatId int    `json:"chat_id"`
	Text   string `json:"text"`
}

type From struct {
	FirstName string `json:"first_name"`
}
