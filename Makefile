DOCKER_COMPOSE ?= docker-compose

compose-dev:
	$(DOCKER_COMPOSE) -f deployments/compose/compose-dev.yaml up -d --build

install: compose-dev migrate

migrate: migrate-drop migrate-up

migrate-drop:
	chmod +x ./deployments/scripts/migration-run.sh
	./deployments/scripts/migration-run.sh drop -f

migrate-up:
	chmod +x ./deployments/scripts/migration-run.sh
	./deployments/scripts/migration-run.sh up

make run-web:
	docker run --privileged --name horusec-all-in-one -p 8000:8000 -p 8001:8001 -p 8003:8003 -p 8005:8005 \
 	-p 8006:8006 -p 8043:8080 horusec-all-in-one:latest

make stop-web:
	docker rm -f horusec-all-in-one
