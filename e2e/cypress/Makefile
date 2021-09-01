DOCKER_COMPOSE ?= docker-compose
NPM ?= npm
CD ?= cd
MAKE ?= make

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
	$(MAKE) migrate-horusec-postgresql

compose-e2e-auth-keycloak-without-application-admin:
	$(DOCKER_COMPOSE) -f ./deployments/docker-compose.auth-keycloak.without-application-admin.yaml down -v
	$(DOCKER_COMPOSE) -f ./deployments/docker-compose.auth-keycloak.without-application-admin.yaml up -d --build --force-recreate
	$(MAKE) migrate-horusec-postgresql
