package telegram

type Message struct {
	ID int `json:"message_id"`
	From Account `json:"from"`
	Photo Photo `json:"photo"`
}