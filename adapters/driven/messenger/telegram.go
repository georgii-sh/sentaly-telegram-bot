package messenger

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"sentaly.com/telegram-bot/ports"
)


type telegramMessenger struct {
	token string
}

// NewTelegramMessenger create new service
func NewTelegramMessenger(token string) ports.Messenger {
	return &telegramMessenger {
		token: token,
	} 
}

func (t *telegramMessenger) SendText(chatID int, text string) (string, error) {
	var apiURL string = "https://api.telegram.org/bot" + t.token + "/sendMessage"
	response, err := http.PostForm(
		apiURL,
		url.Values{
			"chat_id": {strconv.Itoa(chatID)},
			"text":    {text},
		})

	if err != nil {
		log.Printf("error when posting text to the chat: %s", err.Error())
		return "", err
	}
	defer response.Body.Close()

	var bodyBytes, errRead = ioutil.ReadAll(response.Body)
	if errRead != nil {
		log.Printf("error in parsing telegram answer %s", errRead.Error())
		return "", err
	}
	bodyString := string(bodyBytes)
	log.Printf("Body of Telegram Response: %s", bodyString)

	return bodyString, nil
}

