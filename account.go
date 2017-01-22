package telegram

type Account struct {
	ID int `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	UserName string `json:"username"`
}

