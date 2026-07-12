FROM golang:alpine AS base

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

RUN go install github.com/air-verse/air@latest

COPY . .

FROM base AS dev
CMD ["air", "--build.cmd", "go build -p 2 -o tmp/main ./cmd/api", "--build.bin", "./tmp/main"]

FROM base AS builder
RUN CGO_ENABLED=0 GOOS=linux go build -p 2 -o /app/mesenshuttle-api ./cmd/api

FROM alpine:latest AS prod
WORKDIR /app

RUN apk add --no-cache tzdata

COPY --from=builder /app/mesenshuttle-api /app/mesenshuttle-api

EXPOSE 8080

CMD ["/app/mesenshuttle-api"]
