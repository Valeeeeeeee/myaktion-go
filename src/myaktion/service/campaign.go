package service

import (
	"errors"

	log "github.com/sirupsen/logrus"

	"github.com/Valeeeeeeee/myaktion-go/src/myaktion/model"
)

var (
	campaignStore      = make(map[uint]*model.Campaign)
	actCampaignID uint = 1
)

func CreateCampaign(campaign *model.Campaign) error {
	campaign.ID = actCampaignID
	campaignStore[actCampaignID] = campaign
	actCampaignID += 1
	log.Infof("Successfully stored new campaign with ID %v in database.", campaign.ID)
	log.Tracef("Stored: %v", campaign)
	return nil
}

func GetCampaigns() ([]model.Campaign, error) {
	var campaigns []model.Campaign
	for _, campaign := range campaignStore {
		campaigns = append(campaigns, *campaign)
	}
	log.Tracef("Retrieved: %v", campaigns)
	return campaigns, nil
}

func GetCampaign(id uint) (*model.Campaign, error) {
	for _, campaign := range campaignStore {
		if campaign.ID == id {
			log.Tracef("Retrieved campaign: %v", campaign)
			return campaign, nil
		}
	}
	log.Tracef("Could not find campaign with id %v", id)
	return nil, nil
}

func UpdateCampaign(id uint, campaign *model.Campaign) (*model.Campaign, error) {
	// TODO check if exists -> else there could be a problem with the counter
	savedCampaign, err := GetCampaign(id)
	if err != nil {
		log.Errorf("An error occurred when getting campaign with ID %v", id)
		return nil, err
	}
	if savedCampaign == nil {
		log.Infof("No campaign exists with ID %v", id)
		return nil, nil
	}
	savedCampaign.Name = campaign.Name
	savedCampaign.MinumumDonation = campaign.MinumumDonation
	savedCampaign.TargetAmount = campaign.TargetAmount
	savedCampaign.Account = campaign.Account
	savedCampaign.OrganizerName = campaign.OrganizerName

	entry := log.WithField("ID", id)
	entry.Info("Successfully updated campaign.")
	entry.Tracef("Updated: %v", savedCampaign)
	return savedCampaign, nil
}

func DeleteCampaign(id uint) error {
	campaign, err := GetCampaign(id)
	if err != nil {
		log.Errorf("An error occurred when getting campaign with ID %v", id)
		return err
	}
	if campaign == nil {
		log.Infof("No campaign exists with ID %v", id)
		return errors.New("404 Campaign not found")
	}
	delete(campaignStore, id)
	entry := log.WithField("ID", id)
	entry.Info("Successfully deleted campaign.")
	return nil
}
