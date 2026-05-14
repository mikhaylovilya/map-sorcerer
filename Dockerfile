FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod go.sum .
RUN go mod download

COPY . .

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64 
RUN go build -o ./api ./cmd/api

FROM alpine:3

WORKDIR /app

COPY --from=builder /app .

ENTRYPOINT ["./api"]
