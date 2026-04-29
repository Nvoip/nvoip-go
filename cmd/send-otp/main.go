package main

import (
	"context"
	"os"

	"github.com/Nvoip/nvoip-go/internal/exampleutil"
)

func main() {
	client := exampleutil.NewClientFromEnv()
	accessToken := exampleutil.AccessTokenOrCreate(context.Background(), client)

	payload := map[string]any{}
	if sms := firstNonEmpty(os.Getenv("NVOIP_OTP_SMS"), os.Getenv("NVOIP_TARGET_NUMBER")); sms != "" {
		payload["sms"] = sms
	}
	if voice := os.Getenv("NVOIP_OTP_VOICE"); voice != "" {
		payload["voice"] = voice
	}
	if email := os.Getenv("NVOIP_OTP_EMAIL"); email != "" {
		payload["email"] = email
	}

	exampleutil.PrintJSON(client.SendOTP(context.Background(), accessToken, payload))
}

func firstNonEmpty(value, fallback string) string {
	if value == "" {
		return fallback
	}
	return value
}
