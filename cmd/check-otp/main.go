package main

import (
	"context"

	"github.com/Nvoip/nvoip-go/internal/exampleutil"
)

func main() {
	client := exampleutil.NewClientFromEnv()
	exampleutil.PrintJSON(client.CheckOTP(
		context.Background(),
		exampleutil.AccessTokenOrCreate(context.Background(), client),
		exampleutil.MustEnv("NVOIP_OTP_CODE"),
		exampleutil.MustEnv("NVOIP_OTP_KEY"),
	))
}
