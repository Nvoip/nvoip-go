package main

import (
	"context"
	"os"

	"github.com/Nvoip/nvoip-go/internal/exampleutil"
)

func main() {
	client := exampleutil.NewClientFromEnv()
	accessToken := exampleutil.AccessTokenOrCreate(context.Background(), client)
	exampleutil.PrintJSON(client.CreateCall(
		context.Background(),
		accessToken,
		exampleutil.MustEnv("NVOIP_CALLER"),
		firstNonEmpty(os.Getenv("NVOIP_TARGET_NUMBER"), "11999999999"),
	))
}

func firstNonEmpty(value, fallback string) string {
	if value == "" {
		return fallback
	}
	return value
}
