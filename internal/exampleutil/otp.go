package exampleutil

import "fmt"

// OTPPayload maps optional channels to the v3 contract's single phoneNumber.
func OTPPayload(sms, voice, email string) (map[string]any, error) {
	if sms != "" && voice != "" && sms != voice {
		return nil, fmt.Errorf("SMS and voice OTP must use the same phoneNumber; make separate requests for distinct destinations")
	}
	if sms == "" && voice == "" && email == "" {
		return nil, fmt.Errorf("missing OTP destination")
	}
	body := map[string]any{}
	methods := map[string]bool{}
	if sms != "" {
		body["phoneNumber"] = sms
		methods["sms"] = true
	}
	if voice != "" {
		body["phoneNumber"] = voice
		methods["torpedo"] = true
	}
	if email != "" {
		body["email"] = email
		methods["email"] = true
	}
	body["methods"] = methods
	return body, nil
}
