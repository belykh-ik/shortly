FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN apt-get update && apt-get install -y ca-certificates librdkafka1 && rm -rf /var/lib/apt/lists/*

CMD ["go", "run", "cmd/main.go"]
