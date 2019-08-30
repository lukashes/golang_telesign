package rest

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const (
	host = "https://rest-api.telesign.com"
)

type Client struct {
	customerID string
	secretKey  string
	host       string
	httpClient *http.Client
}

func NewClient(customerID, secretKey string) *Client {
	return &Client{
		customerID: customerID,
		secretKey:  secretKey,
		httpClient: &http.Client{
			Transport: &http.Transport{
				MaxIdleConnsPerHost: MaxIdleConnections,
			},
			Timeout: time.Duration(RequestTimeout) * time.Second,
		},
		host: host,
	}
}

func (c *Client) setHost(host string) {
	c.host = host
}

func (c *Client) setCustomerID(cid string) {
	c.customerID = cid
}

func (c *Client) setSecretKey(key string) {
	c.secretKey = key
}

func (c *Client) execute(resource string, method string, fields url.Values) (res TSResponse, _ error) {
	var (
		encoded  = fields.Encode()
		req, err = http.NewRequest(method, c.host+resource, bytes.NewBufferString(encoded))
	)
	if err != nil {
		return res, err
	}
	headers := GenerateTelesignHeaders(
		c.customerID,
		c.secretKey,
		method,
		resource,
		encoded,
		"",
		"",
		"",
	)
	for k, v := range headers {
		req.Header.Add(k, v)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}

	res.Ok = (resp.StatusCode == http.StatusOK)
	res.StatusCode = resp.StatusCode
	res.Header = resp.Header
	res.Body = body

	return res, nil
}

func (c *Client) Get(resource string, fields url.Values) (TSResponse, error) {
	return c.execute(resource, http.MethodGet, fields)
}

func (c *Client) Post(resource string, fields url.Values) (TSResponse, error) {
	return c.execute(resource, http.MethodPost, fields)
}

func (c *Client) Put(resource string, fields url.Values) (TSResponse, error) {
	return c.execute(resource, http.MethodPut, fields)
}

func (c *Client) Delete(resource string, fields url.Values) (TSResponse, error) {
	return c.execute(resource, http.MethodDelete, fields)
}

func (c *Client) Ping() (TSResponse, error) {
	return c.Get(PING_RESOURCE, url.Values{})
}

// Message sends a message to the target phone.
// See https://developer.telesign.com/v2.0/docs/messaging-api for detailed API documentation.
func (c *Client) Message(phone string, message string, messageType string, params map[string]string) (TSResponse, error) {
	fields := url.Values{}
	fields.Set("phone_number", phone)
	fields.Add("message", message)
	fields.Add("message_type", messageType)
	for key, value := range params {
		fields.Add(key, value)
	}

	return c.Post(MESSAGING_RESOURCE, fields)
}
