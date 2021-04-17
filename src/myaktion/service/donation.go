package service

import (
	log "github.com/sirupsen/logrus"

	"github.com/Valeeeeeeee/myaktion-go/src/myaktion/db"
	"github.com/Valeeeeeeee/myaktion-go/src/myaktion/model"
)

func AddDonation(id uint, donation *model.Donation) error {
	donation.CampaignID = id
	result := db.DB.Create(donation)
	if result.Error != nil {
		return result.Error
	}
	entry := log.WithField("ID", id)
	entry.Info("Successfully added donation to campaign.")
	entry.Tracef("Stored: %v", donation)
	return nil
}
