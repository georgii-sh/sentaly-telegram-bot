package application

import (
	"encoding/json"
	"io"
	"log"

	"sentaly.com/telegram-bot/ports"
)

type delegatingSevice struct {
	messenger ports.Messenger
}

// NewDelegatingBotService create new service
func NewDelegatingBotService(m ports.Messenger) ports.BotService {
	return &delegatingSevice {
		messenger: m,
	} 
}

// ParseRequest handles request from telegram
func (s *delegatingSevice) ParseRequest(r io.ReadCloser) (*ports.Update, error) {
	var update ports.Update
	if err := json.NewDecoder(r).Decode(&update); err != nil {
		log.Printf("could not decode incoming update %s", err.Error())
		return nil, err
	}

	return &update, nil
}

// ProcessRequest handles incoming update from the Telegram web hook
func (s *delegatingSevice) ProcessRequest(u *ports.Update) error {
	var body, err = s.messenger.SendText(u.Message.Chat.Id, "Thank you for your message")
	if err != nil {
		log.Printf("got error %s from telegram, reponse body is %s", err.Error(), body)
		return err
	}

	log.Printf("punchline %s successfuly distributed to chat id %d", "$Message", u.Message.Chat.Id)

	return nil
}