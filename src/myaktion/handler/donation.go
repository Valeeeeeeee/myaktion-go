package handler

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/Valeeeeeeee/myaktion-go/src/myaktion/model"
	"github.com/Valeeeeeeee/myaktion-go/src/myaktion/service"
)

func AddDonation(w http.ResponseWriter, r *http.Request) {
	id, err := getId(r)
	if err != nil {
		log.Errorf("Error getting id: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	donation, err := getDonationFromRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedCampaign, err := service.AddDonation(id, donation)
	if err != nil {
		log.Errorf("Error calling service AddDonation: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sendJson(w, *updatedCampaign)
}

func getDonationFromRequest(r *http.Request) (*model.Donation, error) {
	var donation *model.Donation
	err := json.NewDecoder(r.Body).Decode(&donation)
	if err != nil {
		log.Errorf("Can't serialize request body to donation struct: %v", err)
		return nil, err
	}
	return donation, nil
}
