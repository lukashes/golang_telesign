package rest

import (
    "net/url"
)

const (
    VERIFICATION_STATUS_RESOURCE string = "/v1/mobile/verification/status/"
)

func VerificationStatus(external_id string, params map[string]string) TSResponse {
    /*
    Retrieves the verification result for an AutoVerify transaction by external_id. To ensure a secure verification
    flow you must check the status using TeleSign's servers on your backend. Do not rely on the SDK alone to
    indicate a successful verification.

    See https://developer.telesign.com/docs/auto-verify-sdk#section-obtaining-verification-status for detailed API
    documentation.
    */
    
    resource := VERIFICATION_STATUS_RESOURCE + external_id
    fields := url.Values{}
    for key, value := range params {
        fields.Add(key, value)
    }

    return Get(resource, fields)
}