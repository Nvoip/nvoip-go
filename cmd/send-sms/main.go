package main

import (
	"context"
	"os"

	"github.com/Nvoip/nvoip-go/internal/exampleutil"
)

func main() {
	client := exampleutil.NewClientFromEnv()
	accessToken := exampleutil.AccessTokenOrCreate(context.Background(), client)
	exampleutil.PrintJSON(client.SendSMS(
		context.Background(),
		accessToken,
		firstNonEmpty(os.Getenv("NVOIP_TARGET_NUMBER"), "11999999999"),
		firstNonEmpty(os.Getenv("NVOIP_SMS_MESSAGE"), "Mensagem de teste Nvoip"),
		false,
	))
}

func firstNonEmpty(value, fallback string) string {
	if value == "" {
		return fallback
	}
	return value
}
