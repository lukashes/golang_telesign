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
	var params map[string]string

	// Start a call to listed phone_number 
	fmt.Printf("\nSend call request------ \n")
	message_response := ts.Call(PHONE_NUMBER, "", "ARN", params)
	fmt.Println(string(message_response.Body), "\n")
	
	// Parse out the transaction reference_id from payload
	ref_id, _, _, err := jp.Get(message_response.Body, "reference_id")
	if err != nil || ref_id == nil {
		fmt.Printf("Failed to retrieve a valid reference_id or nil\n")
	}
	fmt.Println("ref_id: %s", string(ref_id))

	// Get call status
	fmt.Printf("\nGet status ------ \n")
	status_response := ts.CallStatus(string(ref_id), params)
	fmt.Println(string(status_response.Body), "\n")

}