#!/bin/bash
docker-compose down
docker volume rm grpc-hydra-bench_data-app grpc-hydra-bench_data-hydra
docker-compose up -d
docker-compose logs -f