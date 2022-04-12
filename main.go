package main

import (
	"github.com/bachmidiqbal/bank/account"
	"github.com/bachmidiqbal/bank/bankcore"
	"github.com/bachmidiqbal/bank/handler"
	"log"
	"net/http"
)

func main() {
	account.Accounts[1001] = &bankcore.Account{
		Customer: bankcore.Customer{
			Name:    "John",
			Address: "Los Angeles, California",
			Phone:   "(213) 555 0147",
		},
		Number: 1001,
	}

	http.HandleFunc("/statement", handler.Statement)
	http.HandleFunc("/deposit", handler.Deposit)
	http.HandleFunc("/withdraw", handler.Withdraw)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
