package main

import (
	"fmt"
	ts "telesign/rest"
	jp "github.com/buger/jsonparser"
)

const (
	PHONE_NUMBER string = "Phone Number to test" 
)

func main() {
	ts.SetCustomerID("Your Customer ID")
	ts.SetSecretKey("Your Secret Key")

	// Empty params dictionary
	params map[string]string

	// Send a PhoneID request
	fmt.Printf("\nSend phoneid request ------ \n")
	message_response := ts.PhoneID(PHONE_NUMBER)
	fmt.Println(string(message_response.Body), "\n")
	
	// Parse out the transaction reference_id from JSON payload
	ref_id, _, _, err := jp.Get(message_response.Body, "reference_id")
	if err != nil || ref_id == nil {
		fmt.Printf("Failed to retrieve a valid reference_id or nil")
	}
	fmt.Println("ref_id: %s", string(ref_id))

}