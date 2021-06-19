package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type items struct {
	Name    string
	surname string // Pay attention to lowercase. Surname will not exported.
	Age     int
}

func main() {

	// Encode to JSON data format and convert it to string and print it out.
	data, _ := json.Marshal(&items{"Mahammadali", "Aghabayli", 23})
	fmt.Println(string(data))

	// Decode from string JSON format to Go type and print out the "Name"
	decoded := items{}
	err := json.Unmarshal([]byte(string(data)), &decoded)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(decoded.Name)
}
