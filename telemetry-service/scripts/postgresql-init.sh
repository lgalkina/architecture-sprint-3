#!/bin/bash

goose -allow-missing -dir ./migrations postgres "user=postgres password=postgres dbname=telemetry_db host=localhost port=5432 sslmode=disable" up
