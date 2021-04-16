package model

type Campaign struct {
	Name            string
	MinumumDonation float64
	TargetAmount    float64
	Account         Account
	OrganizerName   string
	Donations       []Donation
}
