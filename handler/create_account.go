package handler

import (
	"encoding/json"
	"fmt"
	"github.com/bachmidiqbal/bank/account"
	"github.com/bachmidiqbal/bank/bankcore"
	"net/http"
)

var accountId float64 = 1000

func CreateAccount(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		fmt.Fprintf(w, "Invalid http method!")
		return
	}

	var cust bankcore.Customer

	if err := json.NewDecoder(req.Body).Decode(&cust); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	accId := accountId
	accountId++

	account.Accounts[accId] = &bankcore.Account{
		Customer: cust,
		Number:   int32(accId),
	}

	fmt.Fprintf(w, account.Accounts[accId].Statement())
}
