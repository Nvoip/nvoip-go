# nvoip-go

[![CI](https://github.com/Nvoip/nvoip-go/actions/workflows/ci.yml/badge.svg)](https://github.com/Nvoip/nvoip-go/actions/workflows/ci.yml) [![Go Reference](https://pkg.go.dev/badge/github.com/Nvoip/nvoip-go.svg)](https://pkg.go.dev/github.com/Nvoip/nvoip-go) [![GitHub tag](https://img.shields.io/github/v/tag/Nvoip/nvoip-go?style=flat-square)](https://github.com/Nvoip/nvoip-go/releases) [![Nvoip](https://img.shields.io/badge/Nvoip-site-00A3E0?style=flat-square)](https://www.nvoip.com.br/) [![API v2](https://img.shields.io/badge/API-v2-1F6FEB?style=flat-square)](https://www.nvoip.com.br/api/) [![Docs](https://img.shields.io/badge/docs-Apiary-6A737D?style=flat-square)](https://nvoip.docs.apiary.io/) [![Postman](https://img.shields.io/badge/Postman-workspace-FF6C37?style=flat-square)](https://nvoip-api.postman.co/workspace/e671d01f-168a-4c38-8d0e-c217229dd61a/team-quickstart) [![Stack](https://img.shields.io/badge/stack-Go-00ADD8?style=flat-square)](https://github.com/Nvoip/nvoip-api-examples) [![License: GPL-3.0](https://img.shields.io/badge/license-GPL--3.0-blue?style=flat-square)](LICENSE)

SDK e exemplos oficiais da [Nvoip](https://www.nvoip.com.br/) para integrar a API v2 com OAuth, chamadas, OTP, WhatsApp, SMS e saldo em Go.

## Requisitos

- Go 1.21+

## Instalacao

```bash
go get github.com/Nvoip/nvoip-go@v0.1.0
```

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

## Links oficiais

- [Site da Nvoip](https://www.nvoip.com.br/)
- [Documentação da API](https://nvoip.docs.apiary.io/)
- [Página da API](https://www.nvoip.com.br/api/)
- [Workspace Postman](https://nvoip-api.postman.co/workspace/e671d01f-168a-4c38-8d0e-c217229dd61a/team-quickstart)
- [Hub de exemplos](https://github.com/Nvoip/nvoip-api-examples)
