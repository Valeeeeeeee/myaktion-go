package model

type Status string

const (
	TRANSFERRED Status = "Transferred"
	IN_PROCESS  Status = "In Process"
)

type Donation struct {
	Amount           float64 `json:"amount"`
	ReceiptRequested bool    `json:"receiptRequested"`
	DonorName        string  `json:"donorName"`
	Status           Status  `json:"status"`
	Account          Account `json:"account"`
}
