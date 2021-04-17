package model

import "gorm.io/gorm"

type Campaign struct {
	gorm.Model
	Name               string     `gorm:"notNull;size:30" json:"name"`
	MinimumDonation    float64    `gorm:"notNull;check:minimum_donation >= 1.0" json:"donationMinimum"`
	TargetAmount       float64    `gorm:"notNull;check:target_amount >= 10.0" json:"targetAmount"`
	Account            Account    `gorm:"embedded;embeddedPrefix:account_" json:"account"`
	OrganizerName      string     `gorm:"notNull" json:"organizerName"`
	Donations          []Donation `gorm:"foreignKey:CampaignID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"donations"`
	AmountDonatedSoFar float64    `gorm:"-"`
}
