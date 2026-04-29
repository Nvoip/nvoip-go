# nvoip-go

SDK e exemplos oficiais da [Nvoip](https://www.nvoip.com.br/) para integrar a API v2 com OAuth, chamadas, OTP, WhatsApp, SMS e saldo em Go.

## Requisitos

- Go 1.21+

## Configuração

```bash
cp .env.example .env
```

Ou exporte:

```bash
export NVOIP_NUMBERSIP="seu_numbersip"
export NVOIP_USER_TOKEN="seu_user_token"
export NVOIP_OAUTH_CLIENT_ID="seu_client_id"
export NVOIP_OAUTH_CLIENT_SECRET="seu_client_secret"
export NVOIP_CALLER="1049"
export NVOIP_TARGET_NUMBER="11999999999"
```

## Fluxos cobertos

- gerar `access_token`
- renovar token
- consultar saldo
- enviar SMS
- realizar chamada
- enviar OTP
- validar OTP
- listar templates de WhatsApp
- enviar template de WhatsApp

## Exemplos

- `go run ./cmd/create-access-token`
- `go run ./cmd/get-balance`
- `go run ./cmd/create-call`
- `go run ./cmd/send-sms`
- `go run ./cmd/send-otp`
- `go run ./cmd/check-otp`
- `go run ./cmd/list-whatsapp-templates`
- `go run ./cmd/send-whatsapp-template`

## SDK web

Para o fluxo de popup com telefone e código, use em conjunto o repositório `nvoip-web-sdk`. Este repo cobre o consumo server-side da API.

## Documentação oficial

- https://nvoip.docs.apiary.io/
- https://www.nvoip.com.br/api
