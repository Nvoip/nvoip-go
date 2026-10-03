package nvoip

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"
)

type recordingTransport struct {
	request *http.Request
	body    string
}

func (t *recordingTransport) RoundTrip(request *http.Request) (*http.Response, error) {
	t.request = request
	body, _ := io.ReadAll(request.Body)
	t.body = string(body)
	return &http.Response{StatusCode: 200, Body: io.NopCloser(strings.NewReader(`{"access_token":"token"}`)), Header: make(http.Header)}, nil
}

func TestClientCredentialsUsesAuthServerAndBasicAuth(t *testing.T) {
	transport := &recordingTransport{}
	client := NewClient("", "client", "secret")
	client.HTTPClient.Transport = transport

	if _, err := client.CreateClientCredentialsToken(context.Background()); err != nil {
		t.Fatal(err)
	}
	if got := transport.request.URL.String(); got != "https://api.nvoip.com.br/auth/oauth2/token" {
		t.Fatalf("URL = %s", got)
	}
	if got := transport.request.Header.Get("Authorization"); got != "Basic Y2xpZW50OnNlY3JldA==" {
		t.Fatalf("Authorization = %s", got)
	}
	if transport.body != "grant_type=client_credentials" {
		t.Fatalf("body = %s", transport.body)
	}
}
