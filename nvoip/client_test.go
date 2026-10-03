package nvoip

import (
	"context"
	"encoding/base64"
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

type recordingTransport struct {
	request *http.Request
	body    string
	status  int
}

func (t *recordingTransport) RoundTrip(request *http.Request) (*http.Response, error) {
	t.request = request
	t.body = ""
	if request.Body != nil {
		body, _ := io.ReadAll(request.Body)
		t.body = string(body)
	}
	code := t.status
	if code == 0 {
		code = 200
	}
	return &http.Response{StatusCode: code, Body: io.NopCloser(strings.NewReader(`{"ok":true}`)), Header: make(http.Header)}, nil
}

func TestTokenFormAndSpecialBasicCredentials(t *testing.T) {
	transport := &recordingTransport{}
	client := NewClient("", "client: &+á", "dummy: &+é")
	client.HTTPClient.Transport = transport
	if _, err := client.CreateClientCredentialsToken(context.Background()); err != nil {
		t.Fatal(err)
	}
	if transport.request.URL.String() != "https://api.nvoip.com.br/auth/oauth2/token" {
		t.Fatal(transport.request.URL)
	}
	if transport.request.Header.Get("Content-Type") != "application/x-www-form-urlencoded" {
		t.Fatal("form header missing")
	}
	got, _ := base64.StdEncoding.DecodeString(strings.TrimPrefix(transport.request.Header.Get("Authorization"), "Basic "))
	want := url.QueryEscape("client: &+á") + ":" + url.QueryEscape("dummy: &+é")
	if string(got) != want {
		t.Fatalf("Basic was not form encoded: %q", got)
	}
	form, _ := url.ParseQuery(transport.body)
	if form.Get("grant_type") != "client_credentials" || form.Has("password") {
		t.Fatal(form)
	}
	if _, err := client.RefreshAccessToken(context.Background(), "dummy&+á"); err != nil {
		t.Fatal(err)
	}
	form, _ = url.ParseQuery(transport.body)
	fixture := "dummy&+á"
	if form.Get("refresh_token") != fixture {
		t.Fatal("refresh encoding")
	}
}

func TestBalanceAndOtpBearerQueryAndError(t *testing.T) {
	transport := &recordingTransport{}
	client := NewClient("", "client", "dummy")
	client.HTTPClient.Transport = transport
	if _, err := client.GetBalance(context.Background(), "dummy"); err != nil {
		t.Fatal(err)
	}
	if transport.request.URL.String() != "https://api.nvoip.com.br/v3/balance" {
		t.Fatal("wrong balance endpoint")
	}
	if transport.request.Header.Get("Authorization") != "Bearer dummy" {
		t.Fatal("missing Bearer")
	}
	if _, err := client.CheckOTP(context.Background(), "dummy", "001122", "key&+á"); err != nil {
		t.Fatal(err)
	}
	if transport.request.Header.Get("Authorization") != "Bearer dummy" {
		t.Fatal("missing OTP Bearer")
	}
	if transport.request.URL.Query().Get("key") != "key&+á" || transport.request.URL.Query().Has("napikey") {
		t.Fatal("wrong OTP query")
	}
	transport.status = 401
	if _, err := client.GetBalance(context.Background(), "dummy"); err == nil {
		t.Fatal("HTTP error treated as success")
	}
}

func TestMissingClientCredentials(t *testing.T) {
	transport := &recordingTransport{}
	client := NewClient("", "", "")
	client.HTTPClient.Transport = transport
	if _, err := client.CreateClientCredentialsToken(context.Background()); err == nil {
		t.Fatal("missing credentials accepted")
	}
	if transport.request != nil {
		t.Fatal("unexpected network request")
	}
}
