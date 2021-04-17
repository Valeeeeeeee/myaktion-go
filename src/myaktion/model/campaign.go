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

func (c *Campaign) AfterFind(tx *gorm.DB) (err error) {
	var sum float64
	result := tx.Model(&Donation{}).Select("ifnull(sum(amount),0)").Where("campaign_id = ?", c.ID).Scan(&sum)
	if result.Error != nil {
		return result.Error
	}
	c.AmountDonatedSoFar = sum
	return nil
}
