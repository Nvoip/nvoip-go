package exampleutil

import "testing"

func TestOtpContractSupportsEmailOnlyAndRefusesMixedPhoneDestinations(t *testing.T) {
	p, err := OTPPayload("", "", "qa@example.invalid")
	if err != nil {
		t.Fatal(err)
	}
	m := p["methods"].(map[string]bool)
	if !m["email"] || m["sms"] || m["torpedo"] || p["phoneNumber"] != nil {
		t.Fatal(p)
	}
	p, err = OTPPayload("11999990000", "11999990000", "")
	if err != nil {
		t.Fatal(err)
	}
	m = p["methods"].(map[string]bool)
	if !m["sms"] || !m["torpedo"] {
		t.Fatal(p)
	}
	if _, err = OTPPayload("11999990000", "11999990001", ""); err == nil {
		t.Fatal("distinct phone destinations accepted")
	}
	if _, err = OTPPayload("", "", ""); err == nil {
		t.Fatal("empty destination accepted")
	}
}
