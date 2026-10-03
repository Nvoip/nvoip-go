package main

import (
	"context"

	"github.com/Nvoip/nvoip-go/v3/internal/exampleutil"
)

func main() {
	client := exampleutil.NewClientFromEnv()
	exampleutil.PrintJSON(client.CreateClientCredentialsToken(context.Background()))
}
