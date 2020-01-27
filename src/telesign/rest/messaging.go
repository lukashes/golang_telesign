package rest

import (
    "net/url"
)

const (
    MESSAGING_RESOURCE string = "/v1/messaging"
    STATUS_RESOURCE string = "/v1/messaging/"
)

func Message(phone_number string, message string, message_type string, params map[string]string) TSResponse {
    /*    
    Send a message to the target phone_number.
    See https://enterprise.telesign.com/api-reference/apis/sms-api for detailed API documentation.
    */

    fields := url.Values{}
    fields.Set("phone_number", phone_number)
    fields.Add("message", message)
    fields.Add("message_type", message_type)
    for key, value := range params {
        fields.Add(key, value)
    }
    
    return Post(MESSAGING_RESOURCE, fields)
}

func MessageStatus(reference_id string, params map[string]string) TSResponse {
    /*
    Retrieves the current status of the message.
    See https://enterprise.telesign.com/api-reference/apis/sms-api for detailed API documentation.
    */

    resource := STATUS_RESOURCE + reference_id
    fields := url.Values{}
    for key, value := range params {
        fields.Add(key, value)
    }
 
    return Get(resource, fields)
}