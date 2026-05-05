FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod go.sum .
RUN go mod download

COPY . .
RUN go build -o ./api ./cmd/api

FROM alpine:3

WORKDIR /app

COPY --from=builder /app .

ENTRYPOINT ["./api"]
