package model

type Request struct {
	Number string `json:"number"`
	Amount string `json:"amount"`
	From   string `json:"from"`
	To     string `json:"to"`
}
