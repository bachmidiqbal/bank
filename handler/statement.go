package handler

import (
	"fmt"
	"github.com/bachmidiqbal/bank/account"
	"net/http"
	"strconv"
)

func Statement(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		fmt.Fprintf(w, "Invalid http method!")
		return
	}

	numberqs := req.URL.Query().Get("number")

	if numberqs == "" {
		fmt.Fprintf(w, "Account number is missing!")
		return
	}

	if number, err := strconv.ParseFloat(numberqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid account number!")
	} else {
		account, ok := account.Accounts[number]
		if !ok {
			fmt.Fprintf(w, "Account with number %v can't be found!", number)
		} else {
			fmt.Fprintf(w, account.Statement())
		}
	}
}
