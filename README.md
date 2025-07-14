# Shortly

Shortly — это многокомпонентный сервис для сокращения ссылок с поддержкой CLI, REST API, Telegram‑бота и Kafka‑интеграции. Построен по принципам Clean Architecture с фокусом на расширяемость и производительность.

---

## Table of Contents

- [Features](#features)
- [Tech Stack](#tech-stack)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Configuration](#configuration)
  - [Run Locally](#run-locally)
  - [Docker Setup](#docker-setup)
- [API Reference](#api-reference)
  - [Register User](#register-user)
  - [Login User](#login-user)
  - [Create Short Link](#create-short-link)
  - [Get link](#get-link)
  - [Get all links](#get-all-links)
  - [Update link](#update-link)
  - [Delete link](#delete-link)


---

## 🚀 Features

- ⚡ **Мгновенное сокращение** длинных ссылок через CLI, HTTP или Telegram
- 🎯 **Мгновенные редиректы** с высокой производительностью
- 📊 **Реальное время** аналитики переходов по ссылкам через Kafka
- 🤖 Telegram‑бот: сокращение и разворот ссылок в чате
- 🐳 Быстрый старт через Docker & Compose
- 🔌 Kafka-интеграция для событий и логов

---

## Tech Stack

- **Language**: Go 1.24
- **Database**: PostgreSQL
- **ORM**: GORM v1.30.0
- **Routing**: `net/http` & `http.ServeMux`
- **Migrations**: GORM Migrate
- **Env Management**: `joho/godotenv`
- **Message Broker**: Kafka

---

## Getting Started

### Prerequisites

- Go 1.24+
- Docker & Docker Compose
- PostgreSQL
- Kafka

### Configuration

1. Create `.env`
2. Fill in the values:
   ```env
   DSN="host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
   SECRET="your_secret"
   ```

### Run Locally

```bash
go mod download
# start server
go run cmd/main.go
```

The server will listen on `http://localhost:8082`.

### Docker Setup

```bash
docker-compose up -d --build
```

- **shortly** service on port **8082**
- **PostgreSQL** on port **5432**
- **adminer** service on port **8081**

---

## API Reference

### Authorization

### Register User

`POST /auth/register`

**Request**

```json
{
	"email": "test@dev.com",
	"password": "your_test_password",
	"name": "Test_User"
}
```

**Response** (201 Created)

```bash
JWT_TOKEN
```

---

### Methods Using with JWT Token in Headers
```bash
  Authorization: Bareer YOUR_AUTH_TOKEN
```

### Login User

`POST /auth/login`

**Request**

```json
{
	"email": "test@dev.com",
	"password": "your_test_password"
}
```

**Response** (200 OK)

```bash
JWT_TOKEN
```

---

### Create Short Link

`POST /create`

**Request**

```json
{
  "url": "https://example.com/long/path"
}
```

**Response** (201 Created)

```json
{
	"ID": 11,
	"CreatedAt": "2025-07-02T23:33:33.263673+03:00",
	"UpdatedAt": "2025-07-02T23:33:33.263673+03:00",
	"DeletedAt": null,
	"url": "https://test.dev.com",
	"hash": "ZRjbkIMbgKirtNBTcUNgXxqxhinxPh"
}
```

---

### Get link

`GET /link/{hash}`

**Response** (200 OK)

```json
{
	"ID": 11,
	"CreatedAt": "2025-07-02T23:33:33.263673+03:00",
	"UpdatedAt": "2025-07-02T23:33:33.263673+03:00",
	"DeletedAt": null,
	"url": "https://test.dev.com",
	"hash": "ZRjbkIMbgKirtNBTcUNgXxqxhinxPh"
}
```

---

### Get all links

`GET /links?limit=5`

**Response** (200 OK)

```json
[
	{
		"ID": 7,
		"CreatedAt": "2025-06-28T19:16:31.082919+03:00",
		"UpdatedAt": "2025-06-28T20:35:03.469616+03:00",
		"DeletedAt": null,
		"url": "https://google.com",
		"hash": "PKZVdcWpDkPhomiEAKfXXKOamqeOBQ"
	},
	{
		"ID": 8,
		"CreatedAt": "2025-06-28T20:34:40.403099+03:00",
		"UpdatedAt": "2025-06-28T20:34:40.403099+03:00",
		"DeletedAt": null,
		"url": "https://go.dev.com",
		"hash": "eJMUFYTSEXtNrXlxsErlObVdkgWygq"
	}
]
```

---

### Update link

`GET /update/{id}`

**Request**

```json
{
	"url" : "https://google.com"
}
```

**Response** (200 OK)

```json
{
	"url" : "https://google.com"
}
```

---

### Delete link

`GET /delete/{id}`

**Response** (200 OK)

---