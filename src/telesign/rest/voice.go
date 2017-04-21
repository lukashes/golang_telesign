package rest

import (
    "net/url"
)

const (
    VOICE_RESOURCE string = "/v1/voice"
    VOICE_STATUS_RESOURCE string = "/v1/voice/"
)

func Call(phone_number string, message string, message_type string, params map[string]string) TSResponse {
    /*    
    Send a voice call to the target phone_number.
    See https://developer.telesign.com/docs/voice-api for detailed API documentation.
    */

    fields := url.Values{}
    fields.Set("phone_number", phone_number)
    fields.Add("message", message)
    fields.Add("message_type", message_type)
    for key, value := range params {
        fields.Add(key, value)
    }

    return Post(VOICE_RESOURCE, fields)
}

func CallStatus(reference_id string, params map[string]string) TSResponse {
    /*
    Retrieves the current status of the voice call.
    See https://developer.telesign.com/docs/voice-api for detailed API documentation.
    */

    resource := VOICE_STATUS_RESOURCE + reference_id
    fields := url.Values{}
    for key, value := range params {
        fields.Add(key, value)
    }
    
    return Get(resource, fields)
}