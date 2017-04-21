package rest

import (
    "net/url"
)

const (
    PHONEID_RESOURCE string = "/v1/phoneid/"
)

func PhoneID(phone_number string, params map[string]string) TSResponse {
    /*    
    The PhoneID API provides a cleansed phone number, phone type, and telecom carrier information 
    to determine the best communication method - SMS or voice.

    See https://developer.telesign.com/docs/phoneid-api for detailed API documentation.
    */

    resource := PHONEID_RESOURCE + phone_number
    fields := url.Values{}
    for key, value := range params {
        fields.Add(key, value)
    }

    return Post(resource, fields)
}
