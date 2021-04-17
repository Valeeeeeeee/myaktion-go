package service

import (
	"errors"

	log "github.com/sirupsen/logrus"

	"github.com/Valeeeeeeee/myaktion-go/src/myaktion/model"
)

func AddDonation(id uint, donation *model.Donation) (*model.Campaign, error) {
	if existingCampaign, err := campaignStore[id]; err {
		existingCampaign.Donations = append(existingCampaign.Donations, *donation)
		entry := log.WithField("ID", id)
		entry.Info("Successfully added donation to campaign.")
		entry.Tracef("Updated: %v", existingCampaign)
		return existingCampaign, nil
	}
	return nil, errors.New("campaign not found")
}
