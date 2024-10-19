FROM golang:1.22.4 AS builder
RUN go version

WORKDIR /app
COPY . .
RUN set -Eeux && \
    go mod download && \
    go mod verify

RUN GOOS=linux GOARCH=amd64 \
    go build \
    -o device-service cmd/device-service/main.go

FROM scratch
WORKDIR /root/

COPY --from=builder /app/device-service device-service

ENTRYPOINT ["./device-service"]