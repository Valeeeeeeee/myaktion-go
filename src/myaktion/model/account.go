package model

type Account struct {
	IBAN       string `json:"number"`
	Name       string `json:"name"`
	NameOfBank string `json:"bankName"`
}
