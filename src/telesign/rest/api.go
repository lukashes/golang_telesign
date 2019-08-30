package rest

import (
	"bytes"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

// Global Inits
var (
	defClient     *Client
	http_client   *http.Client
	rest_api_host string
	customer_id   string
	secret_key    string

	ts_err_log = log.New(os.Stdout, "ERROR ", log.LstdFlags)
)

// Constants
const (
	MaxIdleConnections int = 20
	RequestTimeout     int = 10

	PING_RESOURCE string = "/ping"
)

// init http_client
func init() {
	defClient = NewClient("", "")

	// Todo: Remove QA 0 account after tests are complete
	customer_id = "Your customer_id"
	secret_key = "Your secret_key"
}

// Configure API methods
func SetCustomerID(cid string) {
	defClient.setCustomerID(cid)
}

func SetSecretKey(key string) {
	defClient.setSecretKey(key)
}

func SetAPIHost(hostname string) {
	defClient.setHost(hostname)
}

// TSResponse response structure
type TSResponse struct {
	Ok         bool
	StatusCode int
	Header     http.Header
	Body       []byte
}

func pseudo_uuid() (uuid string) {

	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		ts_err_log.Printf("Failed to create UUID: %s", err)
		return
	}

	uuid = fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return
}

func set_optional_header_values(
	curr_date string,
	nonce string) (string, string) {

	// Setting default params
	if curr_date == "" {
		now := time.Now().UTC()
		curr_date = now.Format(time.RFC1123)
	}
	// Note: RFC1123 uses UTC, which does not comply with TS requirements on paper,
	// but works fine and is a better standard than using GMT
	if nonce == "" {
		nonce = pseudo_uuid()
	}

	return curr_date, nonce
}

func GenerateTelesignHeaders(
	customer_id string,
	secret_key string,
	method_name string,
	resource string,
	url_encoded_fields string,
	curr_date string,
	nonce string,
	user_agent string) (headers map[string]string) {
	// Returns a map of header name and value pairs to be added to the outgoing request

	curr_date, nonce = set_optional_header_values(
		curr_date,
		nonce)

	content_type := ""
	if method_name == "POST" || method_name == "PUT" {
		content_type = "application/x-www-form-urlencoded"
	}
	auth_method := "HMAC-SHA256"

	// Create the request string to sign
	var string_to_sign_builder bytes.Buffer

	string_to_sign_builder.WriteString(method_name)
	string_to_sign_builder.WriteString("\n")
	string_to_sign_builder.WriteString(content_type)
	string_to_sign_builder.WriteString("\n")
	string_to_sign_builder.WriteString(curr_date)
	string_to_sign_builder.WriteString("\nx-ts-auth-method:")
	string_to_sign_builder.WriteString(auth_method)
	string_to_sign_builder.WriteString("\nx-ts-nonce:")
	string_to_sign_builder.WriteString(nonce)

	if content_type != "" && url_encoded_fields != "" {
		string_to_sign_builder.WriteString("\n")
		string_to_sign_builder.WriteString(url_encoded_fields)
	}

	string_to_sign_builder.WriteString("\n")
	string_to_sign_builder.WriteString(resource)
	string_to_sign := string_to_sign_builder.String()

	// Create signature and auth header value using the secret key
	b64decoded_secret, err := base64.StdEncoding.DecodeString(secret_key)
	if err != nil {
		ts_err_log.Printf("Unable to decode secret key. Error: %s", err)
		return map[string]string{}
	}

	mac := hmac.New(sha256.New, b64decoded_secret)
	mac.Write([]byte(string_to_sign))
	signature := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	authorization := "TSA " + customer_id + ":" + signature

	// Assign headers into a response map
	headers = map[string]string{
		"Authorization":    authorization,
		"Date":             curr_date,
		"Content-Type":     content_type,
		"x-ts-auth-method": auth_method,
		"x-ts-nonce":       nonce}
	// "User-Agent": user_agent }

	return
}

func Get(resource string, fields url.Values) TSResponse {
	var res, err = defClient.Get(resource, fields)
	if err != nil {
		ts_err_log.Printf("Unexpected get error: %s", err)
	}
	return res
}

func Post(resource string, fields url.Values) TSResponse {
	var res, err = defClient.Post(resource, fields)
	if err != nil {
		ts_err_log.Printf("Unexpected post error: %s", err)
	}
	return res
}

func Put(resource string, fields url.Values) TSResponse {
	var res, err = defClient.Put(resource, fields)
	if err != nil {
		ts_err_log.Printf("Unexpected put error: %s", err)
	}
	return res
}

func Delete(resource string, fields url.Values) TSResponse {
	var res, err = defClient.Delete(resource, fields)
	if err != nil {
		ts_err_log.Printf("Unexpected delete error: %s", err)
	}
	return res
}

func Ping() TSResponse {
	var res, err = defClient.Ping()
	if err != nil {
		ts_err_log.Printf("Unexpected ping error: %s", err)
	}
	return res
}
