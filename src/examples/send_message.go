package main

import (
	"fmt"
	ts "telesign/rest"
	jp "github.com/buger/jsonparser"
)

const (
	PHONE_NUMBER string = "Phone Number to test" 
	MSSG string = ""
	MSSG_TYPE string = "ARN"
)

func main() {
	ts.SetCustomerID("Your Customer ID")
	ts.SetSecretKey("Your Secret Key")

	// Empty params dictionary
	var params map[string]string

	// Send a SMS Message with ARN message type
	fmt.Printf("\nSend message request------ \n")
	message_response := ts.Message(PHONE_NUMBER, MSSG, MSSG_TYPE, params)
	fmt.Println(string(message_response.Body), "\n")
	
	// Retrive reference_id from the JSON response
	ref_id, _, _, err := jp.Get(message_response.Body, "reference_id")
	if err != nil || ref_id == nil {
		fmt.Printf("Failed to retrieve a valid reference_id or nil\n")
	}
	fmt.Println("ref_id: %s", string(ref_id))

	// Get message status
	fmt.Printf("\nGet status ------ \n")
	status_response := ts.MessageStatus(string(ref_id), params)
	fmt.Println(string(status_response.Body), "\n")

}