package golang_telesign

import "github.com/lukashes/golang_telesign/src/telesign/rest"

type Client = rest.Client

func New(customerID, secretKey string) *Client {
	return rest.NewClient(customerID, secretKey)
}
