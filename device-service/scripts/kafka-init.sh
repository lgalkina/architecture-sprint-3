#!/bin/bash

docker-compose exec kafka kafka-topics.sh --create --topic commands --bootstrap-server kafka:9092 --partitions 1 --replication-factor 1
