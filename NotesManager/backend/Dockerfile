FROM golang:1.23-alpine AS builder

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o server ./cmd/main.go

FROM alpine:3.17

WORKDIR /app

COPY --from=builder /src/server /app/server
COPY configs /app/configs
COPY .env /app/.env

EXPOSE 8000

ENTRYPOINT ["./server"]
