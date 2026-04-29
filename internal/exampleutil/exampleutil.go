package exampleutil

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/Nvoip/nvoip-go/nvoip"
)

func MustEnv(name string) string {
	value := strings.TrimSpace(os.Getenv(name))
	if value == "" {
		panic(fmt.Sprintf("missing required environment variable: %s", name))
	}
	return value
}

func NewClientFromEnv() *nvoip.Client {
	return nvoip.NewClient(
		os.Getenv("NVOIP_BASE_URL"),
		os.Getenv("NVOIP_OAUTH_CLIENT_ID"),
		os.Getenv("NVOIP_OAUTH_CLIENT_SECRET"),
	)
}

func AccessTokenOrCreate(ctx context.Context, client *nvoip.Client) string {
	if accessToken := strings.TrimSpace(os.Getenv("NVOIP_ACCESS_TOKEN")); accessToken != "" {
		return accessToken
	}

	payload, err := client.CreateAccessToken(ctx, MustEnv("NVOIP_NUMBERSIP"), MustEnv("NVOIP_USER_TOKEN"))
	if err != nil {
		panic(err)
	}

	accessToken, _ := payload["access_token"].(string)
	if accessToken == "" {
		panic("access_token not found in OAuth response")
	}

	return accessToken
}

func PrintJSON(payload any, err error) {
	if err != nil {
		panic(err)
	}

	raw, marshalErr := json.MarshalIndent(payload, "", "  ")
	if marshalErr != nil {
		panic(marshalErr)
	}

	fmt.Println(string(raw))
}

func JSONArrayEnv(name string) []any {
	raw := strings.TrimSpace(os.Getenv(name))
	if raw == "" {
		return nil
	}

	var payload []any
	if err := json.Unmarshal([]byte(raw), &payload); err != nil {
		panic(err)
	}

	return payload
}

func BoolEnv(name string, fallback bool) bool {
	raw := strings.TrimSpace(strings.ToLower(os.Getenv(name)))
	if raw == "" {
		return fallback
	}
	return raw == "1" || raw == "true" || raw == "yes"
}
