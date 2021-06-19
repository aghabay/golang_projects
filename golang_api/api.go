package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Money struct {
	Receiver string `json:"receiverID"`
	Currency string `json:"currency"`
	Sender   string `json:"senderID"`
	Amount   int    `json:"amount"`
}

var transactionList []Money

func init() {
	transactionListJSON := `[
		{
		  "senderID": "Mahammadali",
		  "receiverID": "Alibay",
		  "amount": 100,
		  "currency": "AZN"
		},
		{
		  "senderID": "Alibay",
		  "receiverID": "Mahammadali",
		  "amount": 500,
		  "currency": "PLN"
		},
		{
		  "senderID": "Mahammadali",
		  "receiverID": "Alibay",
		  "amount": 250,
		  "currency": "USD"
		}
	  ]`

	err := json.Unmarshal([]byte(transactionListJSON), &transactionList)
	if err != nil {
		log.Fatal(err)
	}
}

func transactionHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		transactionJson, err := json.Marshal(transactionList)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(transactionJson)
	}
}

func main() {
	http.HandleFunc("/transactions", transactionHandler)
	http.ListenAndServe(":9000", nil)
}
