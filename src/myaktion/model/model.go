package model

type Status string

const (
	TRANSFERRED Status = "Transferred"
	IN_PROCESS  Status = "In Process"
)

type Donation struct {
	Amount           float64
	ReceiptRequested bool
	DonorName        string
	Status           Status
	Account          Account
}
