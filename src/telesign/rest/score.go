package rest

import (
    "net/url"
)

const (
    SCORE_RESOURCE string = "/v1/score/"
)

func Score(phone_number string, account_lifecycle_event string, params map[string]string) TSResponse {
    /*
    Score is an API that delivers reputation scoring based on phone number intelligence, 
    traffic patterns, machine learning, and a global data consortium.

    See https://developer.telesign.com/docs/rest_api-phoneid-score for detailed API documentation.
    */

    resource := STATUS_RESOURCE + phone_number
    fields := url.Values{}
    fields.Set("account_lifecycle_event", account_lifecycle_event)
    for key, value := range params {
        fields.Add(key, value)
    }

    return Post(resource, fields)
}