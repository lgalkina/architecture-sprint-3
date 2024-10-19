FROM golang:1.22-alpine AS builder
RUN go version

WORKDIR /app
COPY . .
RUN set -Eeux && \
    go mod download && \
    go mod verify

RUN GOOS=linux GOARCH=amd64 \
    go build \
    -o telemetry-service cmd/telemetry-service/main.go

FROM scratch
WORKDIR /root/

COPY --from=builder /app/telemetry-service telemetry-service

ENTRYPOINT ["./telemetry-service"]