#!/bin/bash

source runGoModTidy.sh

echo "Executando o consumer..."
docker exec -it goapp go run cmd/consumer/main.go
