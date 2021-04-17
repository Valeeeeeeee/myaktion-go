package model

import "gorm.io/gorm"

type Status string

const (
	TRANSFERRED Status = "Transferred"
	IN_PROCESS  Status = "In Process"
)

type Donation struct {
	gorm.Model
	CampaignID       uint
	Amount           float64 `gorm:"notNull;check:amount >= 1.0" json:"amount"`
	ReceiptRequested bool    `gorm:"notNull" json:"receiptRequested"`
	DonorName        string  `gorm:"notNull;size:40" json:"donorName"`
	Status           Status  `gorm:"notNull;type:ENUM('Transferred','In Process')" json:"status"`
	Account          Account `gorm:"embedded;embeddedPrefix:account_" json:"account"`
}
