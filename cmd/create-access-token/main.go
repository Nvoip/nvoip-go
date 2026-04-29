package main

import (
	"context"

	"github.com/Nvoip/nvoip-go/internal/exampleutil"
)

func main() {
	client := exampleutil.NewClientFromEnv()
	exampleutil.PrintJSON(client.CreateAccessToken(
		context.Background(),
		exampleutil.MustEnv("NVOIP_NUMBERSIP"),
		exampleutil.MustEnv("NVOIP_USER_TOKEN"),
	))
}
