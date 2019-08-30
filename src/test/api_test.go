package main

import (
	"fmt"
	ts "github.com/lukashes/golang_telesign/src/telesign/rest"
	"testing"
)

const (
	CustomerID string = "FFFFFFFF-EEEE-DDDD-1234-AB1234567890"
	SecretKey  string = "TE8sTgg45yusumoN6BYsBVkhyRJ5czgsnCehZaOYldPJdmFh6NeX8kunZ2zU1YWaUw/0wV6xfw=="
)

func Test_generate_telesign_headers_with_post(tst *testing.T) {
	method_name := "POST"
	curr_date := "Wed, 14 Dec 2016 18:20:12 -0000"
	nonce := "A1592C6F-E384-4CDB-BC42-C3AB970369E9"
	resource := "/v1/resource"
	body_params_url_encoded := "test=param"

	expected_authorization_header := `TSA FFFFFFFF-EEEE-DDDD-1234-AB1234567890:A0ky++zMbuQv0X7t2XMNY4omcdtwrUhNrn3H+xvomLk=`

	actual_headers := ts.GenerateTelesignHeaders(
		CustomerID,
		SecretKey,
		method_name,
		resource,
		body_params_url_encoded,
		curr_date,
		nonce,
		"unit_test")

	if expected_authorization_header != actual_headers["Authorization"] {
		fmt.Println("Failed to validate POST Authorization, ", expected_authorization_header, actual_headers["Authorization"])
	}
}

func Test_make_string_to_sign_with_post_unicode_content(tst *testing.T) {
	method_name := "POST"
	curr_date := "Wed, 14 Dec 2016 18:20:12 -0000"
	nonce := "A1592C6F-E384-4CDB-BC42-C3AB970369E9"
	resource := "/v1/resource"
	body_params_url_encoded := "test=%CF%BF"

	expected_authorization_header := `TSA FFFFFFFF-EEEE-DDDD-1234-AB1234567890:XM96Yn9jRqMW9lSmEatzO8U+BEJUS2I4sbTHJzSZSeQ=`

	actual_headers := ts.GenerateTelesignHeaders(
		CustomerID,
		SecretKey,
		method_name,
		resource,
		body_params_url_encoded,
		curr_date,
		nonce,
		"unit_test")

	if expected_authorization_header != actual_headers["Authorization"] {
		fmt.Println("Failed to validate Unicode Authorization, ", expected_authorization_header, actual_headers["Authorization"])
	}
}

func Test_make_string_to_sign_with_get(tst *testing.T) {
	method_name := "GET"
	curr_date := "Wed, 14 Dec 2016 18:20:12 -0000"
	nonce := "A1592C6F-E384-4CDB-BC42-C3AB970369E9"
	resource := "/v1/resource"

	expected_authorization_header := `TSA FFFFFFFF-EEEE-DDDD-1234-AB1234567890:e52pAhcuAcza7AGLbJX9+W1odHkZx7gcKePveMusLM4=`

	actual_headers := ts.GenerateTelesignHeaders(
		CustomerID,
		SecretKey,
		method_name,
		resource,
		"",
		curr_date,
		nonce,
		"unit_test")

	if expected_authorization_header != actual_headers["Authorization"] {
		fmt.Println("Failed to validate GET Authorization, ", expected_authorization_header, actual_headers["Authorization"])
	}
}
