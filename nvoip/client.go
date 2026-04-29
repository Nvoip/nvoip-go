package nvoip

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Client struct {
	BaseURL           string
	HTTPClient        *http.Client
	OAuthClientID     string
	OAuthClientSecret string
}

func NewClient(baseURL, oauthClientID, oauthClientSecret string) *Client {
	if baseURL == "" {
		baseURL = "https://api.nvoip.com.br/v2"
	}

	return &Client{
		BaseURL: strings.TrimRight(baseURL, "/"),
		HTTPClient: &http.Client{
			Timeout: 30 * time.Second,
		},
		OAuthClientID:     oauthClientID,
		OAuthClientSecret: oauthClientSecret,
	}
}

func EncodeBasicAuth(clientID, clientSecret string) string {
	return base64.StdEncoding.EncodeToString([]byte(clientID + ":" + clientSecret))
}

func (c *Client) CreateAccessToken(ctx context.Context, numbersip, userToken string) (map[string]any, error) {
	basicAuth, err := c.resolveBasicAuth()
	if err != nil {
		return nil, err
	}

	values := url.Values{}
	values.Set("username", numbersip)
	values.Set("password", userToken)
	values.Set("grant_type", "password")

	return c.request(ctx, http.MethodPost, "/oauth/token", map[string]string{
		"Authorization": "Basic " + basicAuth,
		"Content-Type":  "application/x-www-form-urlencoded",
	}, strings.NewReader(values.Encode()))
}

func (c *Client) RefreshAccessToken(ctx context.Context, refreshToken string) (map[string]any, error) {
	basicAuth, err := c.resolveBasicAuth()
	if err != nil {
		return nil, err
	}

	values := url.Values{}
	values.Set("grant_type", "refresh_token")
	values.Set("refresh_token", refreshToken)

	return c.request(ctx, http.MethodPost, "/oauth/token", map[string]string{
		"Authorization": "Basic " + basicAuth,
		"Content-Type":  "application/x-www-form-urlencoded",
	}, strings.NewReader(values.Encode()))
}

func (c *Client) GetBalance(ctx context.Context, accessToken string) (map[string]any, error) {
	return c.request(ctx, http.MethodGet, "/balance", map[string]string{
		"Authorization": "Bearer " + accessToken,
	}, nil)
}

func (c *Client) SendSMS(ctx context.Context, accessToken, numberPhone, message string, flashSMS bool) (map[string]any, error) {
	return c.jsonRequest(ctx, http.MethodPost, "/sms", map[string]any{
		"numberPhone": numberPhone,
		"message":     message,
		"flashSms":    flashSMS,
	}, accessToken)
}

func (c *Client) CreateCall(ctx context.Context, accessToken, caller, called string) (map[string]any, error) {
	return c.jsonRequest(ctx, http.MethodPost, "/calls/", map[string]any{
		"caller": caller,
		"called": called,
	}, accessToken)
}

func (c *Client) SendOTP(ctx context.Context, accessToken string, payload map[string]any) (map[string]any, error) {
	return c.jsonRequest(ctx, http.MethodPost, "/otp", payload, accessToken)
}

func (c *Client) CheckOTP(ctx context.Context, code, key string) (map[string]any, error) {
	return c.request(ctx, http.MethodGet, "/check/otp?code="+url.QueryEscape(code)+"&key="+url.QueryEscape(key), nil, nil)
}

func (c *Client) ListWhatsAppTemplates(ctx context.Context, accessToken string) (map[string]any, error) {
	return c.request(ctx, http.MethodGet, "/wa/listTemplates", map[string]string{
		"Authorization": "Bearer " + accessToken,
	}, nil)
}

func (c *Client) SendWhatsAppTemplate(ctx context.Context, accessToken string, payload map[string]any) (map[string]any, error) {
	return c.jsonRequest(ctx, http.MethodPost, "/wa/sendTemplates", payload, accessToken)
}

func (c *Client) resolveBasicAuth() (string, error) {
	if c.OAuthClientID != "" && c.OAuthClientSecret != "" {
		return EncodeBasicAuth(c.OAuthClientID, c.OAuthClientSecret), nil
	}
	return "", fmt.Errorf("missing OAuth client credentials: configure oauthClientID + oauthClientSecret")
}

func (c *Client) jsonRequest(ctx context.Context, method, path string, payload any, accessToken string) (map[string]any, error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	headers := map[string]string{
		"Content-Type": "application/json",
	}
	if accessToken != "" {
		headers["Authorization"] = "Bearer " + accessToken
	}

	return c.request(ctx, method, path, headers, bytes.NewReader(body))
}

func (c *Client) request(ctx context.Context, method, path string, headers map[string]string, body io.Reader) (map[string]any, error) {
	req, err := http.NewRequestWithContext(ctx, method, c.BaseURL+path, body)
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	rawBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var payload map[string]any
	if len(rawBody) == 0 {
		payload = map[string]any{}
	} else if err := json.Unmarshal(rawBody, &payload); err != nil {
		payload = map[string]any{"raw": string(rawBody)}
	}

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("nvoip request failed with status %d: %s", resp.StatusCode, string(rawBody))
	}

	return payload, nil
}
