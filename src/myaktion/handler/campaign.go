package handler

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/Valeeeeeeee/myaktion-go/src/myaktion/model"
	"github.com/Valeeeeeeee/myaktion-go/src/myaktion/service"
)

func CreateCampaign(w http.ResponseWriter, r *http.Request) {
	var campaign *model.Campaign
	campaign, err := getCampaignFromRequest(r)
	entry := log.WithField("url", r.URL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := service.CreateCampaign(campaign); err != nil {
		entry.Errorf("Error calling service CreateCampaign: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sendJson(w, campaign)
}

func GetCampaigns(w http.ResponseWriter, _ *http.Request) {
	campaigns, err := service.GetCampaigns()
	if err != nil {
		log.Errorf("Error calling service GetCampaigns: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if len(campaigns) == 0 {
		sendJson(w, result{Message: "There are no campaigns"})
		return
	}
	sendJson(w, campaigns)
}

func GetCampaign(w http.ResponseWriter, r *http.Request) {
	id, err := getId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	var campaign *model.Campaign
	campaign, err = service.GetCampaign(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if campaign == nil {
		log.Warnf("No campaign found")
		sendJson(w, result{Message: "No campaign found"})
		return
	}

	sendJson(w, *campaign)
}

func UpdateCampaign(w http.ResponseWriter, r *http.Request) {
	id, err := getId(r)
	if err != nil {
		log.Errorf("Error getting id: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	campaign, err := getCampaignFromRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	campaign, err = service.UpdateCampaign(id, campaign)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if campaign == nil {
		http.Error(w, "404 Campaign not found", http.StatusNotFound)
		return
	}

	sendJson(w, *campaign)
}

func DeleteCampaign(w http.ResponseWriter, r *http.Request) {
	id, err := getId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = service.DeleteCampaign(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	sendJson(w, result{Message: "OK"})
}

func getCampaignFromRequest(r *http.Request) (*model.Campaign, error) {
	var campaign model.Campaign
	err := json.NewDecoder(r.Body).Decode(&campaign)
	if err != nil {
		log.Errorf("Can't serialize request body to campaign struct: %v", err)
		return nil, err
	}
	return &campaign, nil
}
