package ports

import (
	"io"
)

// Update is a Telegram object that the handler receives every time an user interacts with the bot.
type Update struct {
	UpdateId int     `json:"update_id"`
	Message  Message `json:"message"`
}

// Message is a Telegram object that can be found in an update.
type Message struct {
	Text     string   `json:"text"`
	Chat     Chat     `json:"chat"`
}

// Chat indicates the conversation to which the message belongs.
type Chat struct {
	Id int `json:"id"`
}

// BotService port
type BotService interface {
	ParseRequest(r io.ReadCloser) (*Update, error)
	ProcessRequest(u *Update) error
}