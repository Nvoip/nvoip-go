package main

import (
	"context"
	"os"

	"github.com/Nvoip/nvoip-go/internal/exampleutil"
)

func main() {
	client := exampleutil.NewClientFromEnv()
	accessToken := exampleutil.AccessTokenOrCreate(context.Background(), client)

	payload := map[string]any{
		"idTemplate":  exampleutil.MustEnv("NVOIP_WA_TEMPLATE_ID"),
		"destination": exampleutil.MustEnv("NVOIP_WA_DESTINATION"),
		"instance":    exampleutil.MustEnv("NVOIP_WA_INSTANCE"),
		"language":    firstNonEmpty(os.Getenv("NVOIP_WA_LANGUAGE"), "pt_BR"),
	}
	if bodyVariables := exampleutil.JSONArrayEnv("NVOIP_WA_BODY_VARIABLES"); bodyVariables != nil {
		payload["bodyVariables"] = bodyVariables
	}
	if headerVariables := exampleutil.JSONArrayEnv("NVOIP_WA_HEADER_VARIABLES"); headerVariables != nil {
		payload["headerVariables"] = headerVariables
	}
	payload["functions"] = map[string]bool{
		"to_flow": exampleutil.BoolEnv("NVOIP_WA_TO_FLOW", false),
	}

	exampleutil.PrintJSON(client.SendWhatsAppTemplate(context.Background(), accessToken, payload))
}

func firstNonEmpty(value, fallback string) string {
	if value == "" {
		return fallback
	}
	return value
}
