DOCKER_COMPOSE ?= docker-compose
NPM ?= npm
CD ?= cd
MAKE ?= make

restart-services:
	docker restart horusec-auth
	echo "Waiting grpc connection..." && sleep 5
	docker restart horusec-vulnerability
	docker restart horusec-webhook
	docker restart horusec-core
	docker restart horusec-api
	docker restart horusec-analytic
	sleep 10

migrate-horusec-postgresql:
	$(CD) ../.. && $(MAKE) migrate-drop && $(MAKE) migrate-up && $(CD) ./e2e/cypress

test-e2e-install:
	$(NPM) install

test-e2e-auth-horusec-without-application-admin: test-e2e-install compose-e2e-auth-horusec-without-application-admin
	$(NPM) run test::auth-horusec::without-application-admin

test-e2e-auth-keycloak-without-application-admin: test-e2e-install compose-e2e-auth-keycloak-without-application-admin
	$(NPM) run test::auth-keycloak::without-application-admin

compose-e2e-auth-horusec-without-application-admin:
	$(DOCKER_COMPOSE) -f ./deployments/docker-compose.auth-horusec.without-application-admin.yaml down -v
	$(DOCKER_COMPOSE) -f ./deployments/docker-compose.auth-horusec.without-application-admin.yaml up -d --build --force-recreate
	make migrate-horusec-postgresql
	make restart-services

compose-e2e-auth-keycloak-without-application-admin:
	$(DOCKER_COMPOSE) -f ./deployments/docker-compose.auth-keycloak.without-application-admin.yaml down -v
	$(DOCKER_COMPOSE) -f ./deployments/docker-compose.auth-keycloak.without-application-admin.yaml up -d --build --force-recreate
	make migrate-horusec-postgresql
	make restart-services
