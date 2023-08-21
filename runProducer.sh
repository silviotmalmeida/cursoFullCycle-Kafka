#!/bin/bash

source runGoModTidy.sh

echo "Executando o producer..."
docker exec -it goapp go run cmd/producer/main.go
