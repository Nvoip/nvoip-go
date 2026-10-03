# nvoip-go

[![CI](https://github.com/Nvoip/nvoip-go/actions/workflows/ci.yml/badge.svg)](https://github.com/Nvoip/nvoip-go/actions/workflows/ci.yml) [![Go Reference](https://pkg.go.dev/badge/github.com/Nvoip/nvoip-go/v3.svg)](https://pkg.go.dev/github.com/Nvoip/nvoip-go/v3) [![GitHub tag](https://img.shields.io/github/v/tag/Nvoip/nvoip-go?style=flat-square)](https://github.com/Nvoip/nvoip-go/releases) [![Nvoip](https://img.shields.io/badge/Nvoip-site-00A3E0?style=flat-square)](https://www.nvoip.com.br/) [![API v3](https://img.shields.io/badge/API-v3-1F6FEB?style=flat-square)](https://www.nvoip.com.br/api/) [![Docs](https://img.shields.io/badge/docs-Apiary-6A737D?style=flat-square)](https://nvoip.docs.apiary.io/) [![Postman](https://img.shields.io/badge/Postman-workspace-FF6C37?style=flat-square)](https://nvoip-api.postman.co/workspace/e671d01f-168a-4c38-8d0e-c217229dd61a/team-quickstart) [![Stack](https://img.shields.io/badge/stack-Go-00ADD8?style=flat-square)](https://github.com/Nvoip/nvoip-api-examples) [![License: GPL-3.0](https://img.shields.io/badge/license-GPL--3.0-blue?style=flat-square)](LICENSE)

SDK e exemplos oficiais da [Nvoip](https://www.nvoip.com.br/) para integrar a API v3 com OAuth, chamadas, OTP, WhatsApp, SMS e saldo em Go.

## Migração para v3

Esta é uma quebra de compatibilidade: use `CreateClientCredentialsToken` e envie o access token RS256 em `Authorization: Bearer`. O SDK usa `https://api.nvoip.com.br/auth/oauth2/token`; não use `napikey`, password grant ou `/v3/oauth/token`. Para SMS de texto livre, valide antes a política e o template aprovado aplicáveis à sua conta.

## Requisitos

- Go 1.21+

## Instalacao

```bash
go get github.com/Nvoip/nvoip-go/v3@v3.0.0
```

## Configuração

```bash
cp .env.example .env
```

Ou exporte:

```bash
export NVOIP_OAUTH_CLIENT_ID="seu_numbersip"
export NVOIP_OAUTH_CLIENT_SECRET="seu_user_token"
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

- `go run ./cmd/create-client-credentials-token`
- `go run ./cmd/get-balance`
- `go run ./cmd/create-call`
- `go run ./cmd/send-sms`
- `go run ./cmd/send-otp`
- `go run ./cmd/check-otp`
- `go run ./cmd/list-whatsapp-templates`
- `go run ./cmd/send-whatsapp-template`

### Destinatário WhatsApp

O exemplo mantém `NVOIP_WA_DESTINATION` para telefone. Para o contrato tipado,
use `NVOIP_WA_RECIPIENT_TYPE=phone|bsuid|parent_bsuid` e
`NVOIP_WA_RECIPIENT_VALUE`, sem `destination`. BSUID é opaco; não use
`@username` nem o coloque em campo de telefone. Exemplos mascarados:
`US.MASKED_BSUID_001` e `PARENT.MASKED_BSUID_001`.

## SDK web

Para o fluxo de popup com telefone e código, use em conjunto o repositório `nvoip-web-sdk`. Este repo cobre o consumo server-side da API.

## Links oficiais

- [Site da Nvoip](https://www.nvoip.com.br/)
- [Documentação da API](https://nvoip.docs.apiary.io/)
- [Página da API](https://www.nvoip.com.br/api/)
- [Workspace Postman](https://nvoip-api.postman.co/workspace/e671d01f-168a-4c38-8d0e-c217229dd61a/team-quickstart)
- [Hub de exemplos](https://github.com/Nvoip/nvoip-api-examples)
