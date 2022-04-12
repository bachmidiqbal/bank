package handler

import (
	"encoding/json"
	"fmt"
	"github.com/bachmidiqbal/bank/account"
	"github.com/bachmidiqbal/bank/model"
	"net/http"
	"strconv"
)

func Deposit(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		fmt.Fprintf(w, "Invalid http method!")
		return
	}

	var reqModel model.Request

	if err := json.NewDecoder(req.Body).Decode(&reqModel); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	numberqs := reqModel.Number
	amountqs := reqModel.Amount

	if numberqs == "" {
		fmt.Fprintf(w, "Account number is missing!")
		return
	}

	if number, err := strconv.ParseFloat(numberqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid account number!")
	} else if amount, err := strconv.ParseFloat(amountqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid amount number!")
	} else {
		account, ok := account.Accounts[number]
		if !ok {
			fmt.Fprintf(w, "Account with number %v can't be found!", number)
		} else {
			err := account.Deposit(amount)
			if err != nil {
				fmt.Fprintf(w, "%v", err)
			} else {
				fmt.Fprintf(w, account.Statement())
			}
		}
	}
}
