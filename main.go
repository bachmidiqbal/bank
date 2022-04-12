package main

import (
	"github.com/bachmidiqbal/bank/handler"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/statement", handler.Statement)
	mux.HandleFunc("/deposit", handler.Deposit)
	mux.HandleFunc("/withdraw", handler.Withdraw)
	mux.HandleFunc("/create-account", handler.CreateAccount)
	mux.HandleFunc("/transfer", handler.Transfer)
	log.Fatal(http.ListenAndServe(":8000", mux))
}
