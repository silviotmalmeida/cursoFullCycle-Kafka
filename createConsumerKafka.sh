#!/bin/bash

echo "Iniciado consumer..."
docker exec -it kafka kafka-console-consumer --bootstrap-server=localhost:9092 --topic=teste