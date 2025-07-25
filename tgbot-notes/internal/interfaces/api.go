package interfaces

type Telegram interface {
	Send(chatID int64, text string) error
}
