# Telemetry Service

## Как запустить локально с docker

Запускаем PostgreSQL и Kafka

```shell
docker compose up -d
```

Инициализируем PostgreSQL, понадобится установить goose

```shell
go install github.com/pressly/goose/v3/cmd/goose@latest
```

```shell
chmod +x ./scripts/postgresql-init.sh
./scripts/postgresql-init.sh
```

Создаем топики в kafka

```shell
chmod +x ./scripts/kafka-init.sh
./scripts/kafka-init.sh
```

Запускаем сервис

```shell
go run ./cmd/main.go
```

## Проверяем работоспособность

Получение телеметрии устройства по id:

```bash
curl -X GET "http://localhost:8080/devices/1/telemetry" -H "Accept: application/json"
```

Получение последней телеметрии устройства по id:

```bash
curl -X GET "http://localhost:8080/devices/1/telemetry/latest" -H "Accept: application/json"
```
 
Отпавка новой телеметрии устройства в очередь 'telemetry' в Kafka:

```
"{\"device_id\": 1, \"temperature\": 14.5, \"timestamp\": \"2023-10-01T12:00:00Z\"}"        
```