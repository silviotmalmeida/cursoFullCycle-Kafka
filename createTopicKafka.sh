#!/bin/bash

echo "Criando o tópico 'teste' no kafka.."
docker exec -it kafka kafka-topics --create --bootstrap-server=localhost:9092 --topic=teste --partitions=3
