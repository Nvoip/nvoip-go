package main

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/Nvoip/nvoip-go/internal/exampleutil"
)

func main() {
	client := exampleutil.NewClientFromEnv()
	accessToken := exampleutil.AccessTokenOrCreate(context.Background(), client)

	payload := map[string]any{
		"idTemplate": exampleutil.MustEnv("NVOIP_WA_TEMPLATE_ID"),
		"instance":   exampleutil.MustEnv("NVOIP_WA_INSTANCE"),
		"language":   firstNonEmpty(os.Getenv("NVOIP_WA_LANGUAGE"), "pt_BR"),
	}
	recipientType := strings.ToLower(strings.TrimSpace(os.Getenv("NVOIP_WA_RECIPIENT_TYPE")))
	recipientValue := strings.TrimSpace(os.Getenv("NVOIP_WA_RECIPIENT_VALUE"))
	if recipientType == "" {
		destination := exampleutil.MustEnv("NVOIP_WA_DESTINATION")
		if !regexp.MustCompile(`^\+?[0-9]{8,20}$`).MatchString(destination) {
			panic("NVOIP_WA_DESTINATION must be a phone number; use recipient for BSUID")
		}
		payload["destination"] = destination
	} else {
		if recipientType != "phone" && recipientType != "bsuid" && recipientType != "parent_bsuid" {
			panic("NVOIP_WA_RECIPIENT_TYPE must be phone, bsuid or parent_bsuid")
		}
		if recipientValue == "" {
			panic("NVOIP_WA_RECIPIENT_VALUE is required with NVOIP_WA_RECIPIENT_TYPE")
		}
		if strings.HasPrefix(recipientValue, "@") {
			panic("@username is not a WhatsApp recipient; use a BSUID or parent BSUID")
		}
		if recipientType == "phone" && !regexp.MustCompile(`^\+?[0-9]{8,20}$`).MatchString(recipientValue) {
			panic("A phone recipient must contain only an optional leading + and 8 to 20 digits")
		}
		if recipientType != "phone" && (strings.ContainsAny(recipientValue, " \t\r\n") || len(recipientValue) > 256) {
			panic("A BSUID must be an opaque value without whitespace (maximum 256 characters)")
		}
		payload["recipient"] = map[string]string{"type": recipientType, "value": recipientValue}
	}
	if bodyVariables := exampleutil.JSONArrayEnv("NVOIP_WA_BODY_VARIABLES"); bodyVariables != nil {
		payload["bodyVariables"] = bodyVariables
	}
	if headerVariables := exampleutil.JSONArrayEnv("NVOIP_WA_HEADER_VARIABLES"); headerVariables != nil {
		payload["headerVariables"] = headerVariables
	}
	toFlow := exampleutil.BoolEnv("NVOIP_WA_TO_FLOW", false)
	if toFlow && (recipientType == "bsuid" || recipientType == "parent_bsuid") {
		panic(fmt.Sprintf("WhatsApp Flow and attendance require a phone recipient, not %s", recipientType))
	}
	payload["functions"] = map[string]bool{"to_flow": toFlow}

	exampleutil.PrintJSON(client.SendWhatsAppTemplate(context.Background(), accessToken, payload))
}

func firstNonEmpty(value, fallback string) string {
	if value == "" {
		return fallback
	}
	return value
}
