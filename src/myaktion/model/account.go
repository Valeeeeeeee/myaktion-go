package model

type Account struct {
	IBAN       string `gorm:"notNull;size:22" json:"number"`
	Name       string `gorm:"notNull;size:60" json:"name"`
	NameOfBank string `gorm:"notNull;size:40" json:"bankName"`
}
