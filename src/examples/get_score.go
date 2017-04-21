package main

import (
	"fmt"
	ts "telesign/rest"
	jp "github.com/buger/jsonparser"
)

const (
	PHONE_NUMBER string = "Phone Number to test" 
	ACCOUNT_LIFECYCLE_EVENT string = "create"
)

func main() {
	ts.SetCustomerID("Your Customer ID")
	ts.SetSecretKey("Your Secret Key")

	// Empty params dictionary
	var params map[string]string

	// Send a Score request
	fmt.Printf("\nSend score request ------ \n")
	message_response := ts.Score(PHONE_NUMBER, ACCOUNT_LIFECYCLE_EVENT, params)
	fmt.Println(string(message_response.Body), "\n")
	
	// Retrive reference_id from the transaction JSON response
	ref_id, _, _, err := jp.Get(message_response.Body, "reference_id")
	if err != nil || ref_id == nil {
		fmt.Printf("Failed to retrieve a valid reference_id or nil\n")
	}
	fmt.Println("ref_id: %s", string(ref_id))

}
