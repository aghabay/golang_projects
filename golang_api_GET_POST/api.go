package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Money struct {
	TransactionID int    `json:"transactionID"`
	Receiver      string `json:"receiverID"`
	Currency      string `json:"currency"`
	Sender        string `json:"senderID"`
	Amount        int    `json:"amount"`
}

var transactionList []Money

func init() {
	transactionListJSON := `[
		{
		  "transactionID": 1,
		  "senderID": "Mahammadali",
		  "receiverID": "Alibay",
		  "amount": 100,
		  "currency": "AZN"
		},
		{
		  "transactionID": 2,
		  "senderID": "Alibay",
		  "receiverID": "Mahammadali",
		  "amount": 500,
		  "currency": "PLN"
		},
		{
		  "transactionID": 3,
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
	case http.MethodGet: // If method is GET, list all transactions.
		transactionJson, err := json.Marshal(transactionList)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(transactionJson)
	case http.MethodPost: // If method is POST, add new transaction
		var newTransaction Money
		bodyBytes, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = json.Unmarshal(bodyBytes, &newTransaction)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if newTransaction.TransactionID != 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Transaction successfully created.
		newTransaction.TransactionID = getNextID()
		transactionList = append(transactionList, newTransaction)
		w.WriteHeader(http.StatusCreated)
		return
	}
}

// getNextID creates a new ID
func getNextID() int {
	highestID := -1
	for _, Money := range transactionList {
		if highestID < Money.TransactionID {
			highestID = Money.TransactionID
		}
	}
	return highestID + 1
}

func main() {
	http.HandleFunc("/transactions", transactionHandler)
	http.ListenAndServe(":9000", nil)
}

/*
Example POST request
Do not send transactionID. Because "getNextID" function will create it.
{
	"receiverID": "Alibay",
	"currency": "USD",
	"senderID": "Mahammadali",
	"amount": 555
}
*/
