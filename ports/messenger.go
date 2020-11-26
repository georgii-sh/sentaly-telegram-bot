package ports

// Messenger port
type Messenger interface {
	SendText(chatId int, text string) (string, error)
}