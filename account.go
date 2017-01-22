package telegram

type Account struct {
	ID int `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	UserName string `json:"username"`
}


type AccountResponse struct {
	Ok bool `json:"ok"`
	Account Account `json:"result"`
}