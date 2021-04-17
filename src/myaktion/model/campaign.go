package model

type Campaign struct {
	ID              uint
	Name            string     `json:"name"`
	MinumumDonation float64    `json:"donationMinimum"`
	TargetAmount    float64    `json:"targetAmount"`
	Account         Account    `json:"account"`
	OrganizerName   string     `json:"organizerName"`
	Donations       []Donation `json:"donations"`
}
