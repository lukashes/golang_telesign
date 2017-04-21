package main

import (
	"fmt"
	ts "telesign/rest"
)

const (
	EXTERNAL_ID string = "Your transactions external id"
)

func main() {
	ts.SetCustomerID("Your Customer ID")
	ts.SetSecretKey("Your Secret Key")

	// Empty params dictionary
	var params map[string]string
	status_response := ts.MessageStatus(string(EXTERNAL_ID), params)
	fmt.Println(string(status_response.Body))
	
}
