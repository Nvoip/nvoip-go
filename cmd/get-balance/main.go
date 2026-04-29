package main

import (
	"context"

	"github.com/Nvoip/nvoip-go/internal/exampleutil"
)

func main() {
	client := exampleutil.NewClientFromEnv()
	accessToken := exampleutil.AccessTokenOrCreate(context.Background(), client)
	exampleutil.PrintJSON(client.GetBalance(context.Background(), accessToken))
}
