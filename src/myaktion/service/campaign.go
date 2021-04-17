package service

import (
	"errors"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"github.com/Valeeeeeeee/myaktion-go/src/myaktion/db"
	"github.com/Valeeeeeeee/myaktion-go/src/myaktion/model"
)

var (
//campaignStore = make(map[uint]*model.Campaign)
)

func CreateCampaign(campaign *model.Campaign) error {
	result := db.DB.Create(campaign)
	if result.Error != nil {
		return result.Error
	}
	log.Infof("Successfully stored new campaign with ID %v in database.", campaign.ID)
	log.Tracef("Stored: %v", campaign)
	return nil
}

func GetCampaigns() ([]model.Campaign, error) {
	var campaigns []model.Campaign
	result := db.DB.Preload("Donations").Find(&campaigns)
	if result.Error != nil {
		return nil, result.Error
	}
	log.Tracef("Retrieved: %v", campaigns)
	return campaigns, nil
}

func GetCampaign(id uint) (*model.Campaign, error) {
	campaign := new(model.Campaign)
	result := db.DB.Preload("Donations").First(campaign, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}
	log.Tracef("Retrieved: %v", campaign)
	return campaign, nil
}

func UpdateCampaign(id uint, campaign *model.Campaign) (*model.Campaign, error) {
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
	savedCampaign.MinimumDonation = campaign.MinimumDonation
	savedCampaign.TargetAmount = campaign.TargetAmount
	savedCampaign.Account = campaign.Account
	savedCampaign.OrganizerName = campaign.OrganizerName
	result := db.DB.Save(savedCampaign)
	// result := db.DB.Model(&savedCampaign).Updates(campaign)
	// result := db.DB.Model(&savedCampaign).Updates(map[string]interface{}{
	// 	"name":             campaign.Name,
	// 	"minimum_donation": campaign.MinimumDonation,
	// 	"target_amount":    campaign.TargetAmount,
	// 	// "account":          campaign.Account,
	// 	"organizer_name": campaign.OrganizerName})
	if result.Error != nil {
		return nil, result.Error
	}

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
	// delete(campaignStore, id)
	result := db.DB.Delete(&campaign)
	if result.Error != nil {
		return result.Error
	}
	entry := log.WithField("ID", id)
	entry.Info("Successfully deleted campaign.")
	return nil
}
