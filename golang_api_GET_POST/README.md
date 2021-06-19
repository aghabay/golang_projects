# GoLang Example API project.
```sh
Example POST request to add a new transaction.
Do not send transactionID. Because "getNextID" function will create it automatically.

{
	"receiverID": "Alibay",
	"currency": "USD",
	"senderID": "Mahammadali",
	"amount": 555
}

We can make a GET request to retrieve all transactions:
http://localhost:9000/transactions
```
