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