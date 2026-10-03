package main

import (
	"context"
	"github.com/Nvoip/nvoip-go/internal/exampleutil"
	"os"
)

func main() {
	sms := os.Getenv("NVOIP_OTP_SMS")
	if sms == "" {
		sms = os.Getenv("NVOIP_TARGET_NUMBER")
	}
	payload, err := exampleutil.OTPPayload(sms, os.Getenv("NVOIP_OTP_VOICE"), os.Getenv("NVOIP_OTP_EMAIL"))
	if err != nil {
		panic(err)
	}
	client := exampleutil.NewClientFromEnv()
	bearer := exampleutil.AccessTokenOrCreate(context.Background(), client)
	exampleutil.PrintJSON(client.SendOTP(context.Background(), bearer, payload))
}
