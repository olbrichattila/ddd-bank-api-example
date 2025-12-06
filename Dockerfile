# Stage 1: Builder
FROM golang:1.25 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY internal/ ./internal
COPY cmd/ ./cmd

RUN go install github.com/olbrichattila/godbmigrator_cmd/cmd/migrator@latest
RUN go build -o bankapi ./cmd/eaglebank

# Stage 2: Runtime
FROM debian:bookworm-slim


WORKDIR /app
COPY --from=builder /app/bankapi /app/bankapi
COPY --from=builder /go/bin/migrator /app/migrator

COPY migrations/ /app/migrations/
COPY .env.migrator /app/.env.migrator

COPY start.sh /app/start.sh
RUN chmod +x /app/start.sh \
    && mkdir /app/session \
    && mkdir /app/data


CMD ["/app/start.sh"]
