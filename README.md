# Shortly

A lightning‑fast, minimalist Go URL shortener with support for custom aliases and built‑in click analytics.

---

## Features

- 📊 Click analytics per link (timestamps, count)  
- 🔒 Zero external dependencies—pure Go + PostgreSQL  
- 🐳 Docker & Docker Compose for one‑command setup  

---

## Tech Stack

- **Language**: Go 1.24 
- **Database**: PostgreSQL  
- **Router**: `net/http` & `http.ServeMux`  
- **Migrations**: raw SQL (in `migrations/`)  
- **Containerization**: Docker & Docker Compose  

## Running with Docker

```
docker-compose up --build
```

