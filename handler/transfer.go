package handler

import (
	"encoding/json"
	"fmt"
	"github.com/bachmidiqbal/bank/account"
	"github.com/bachmidiqbal/bank/bankcore"
	"github.com/bachmidiqbal/bank/model"
	"net/http"
	"strconv"
)

func Transfer(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		fmt.Fprintf(w, "Invalid http method!")
		return
	}

	var reqModel model.Request

	if err := json.NewDecoder(req.Body).Decode(&reqModel); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if fromNumber, err := strconv.ParseFloat(reqModel.From, 64); err != nil {
		fmt.Fprintf(w, "Invalid from account number!")
	} else if toNumber, err := strconv.ParseFloat(reqModel.To, 64); err != nil {
		fmt.Fprintf(w, "Invalid to account number!")
	} else if amount, err := strconv.ParseFloat(reqModel.Amount, 64); err != nil {
		fmt.Fprintf(w, "Invalid amount!")
	} else {
		fromAccount := account.Accounts[fromNumber]
		toAccount := account.Accounts[toNumber]

		if err := fromAccount.Withdraw(amount); err != nil {
			fmt.Fprintf(w, "Failed to withdraw amount from sender account!")
			return
		}

		if err := toAccount.Deposit(amount); err != nil {
			fmt.Fprintf(w, "Failed to deposit amount to receiver account!")

			if err := fromAccount.Deposit(amount); err != nil {
				fmt.Fprintf(w, "Failed to rollback amount to sender account!")
			}

			return
		}

		res := [2]*bankcore.Account{fromAccount, toAccount}

		jbyte, _ := json.Marshal(res)

		fmt.Fprintf(w, string(jbyte))
	}
}
