# Device Service 

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

Получение данных устройства по id:

```bash
curl -X GET "http://localhost:8080/devices/1" -H "Accept: application/json"
```

Обовление статуса устройства по id:

```bash
curl -X PUT "http://localhost:8080/devices/1/status" -H "Content-Type: application/json" -d '{
  "status": "Inactive"
}'
```

Отправка команды устройству:

```bash
curl -X POST "http://localhost:8080/devices/1/commands" -H "Content-Type: application/json" -d '{
    "command": "turn_off",
    "value": 0
}'
```

## Как запустить c minikube и terraform

Запуск с terraform пока не работает - таймауты при выполнении `terraform apply`