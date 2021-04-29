DOCKER_COMPOSE ?= docker-compose

compose-dev:
	$(DOCKER_COMPOSE) -f deployments/compose/compose-dev.yaml up -d --build
