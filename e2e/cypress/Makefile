DOCKER_COMPOSE ?= docker-compose
NPM ?= npm

migrate-up:
	cd ../.. && make migrate-up && cd ./e2e/cypress

test-e2e-install:
    $(NPM) install

test-e2e-auth-horusec-without-application-admin: compose-e2e-auth-horusec-without-application-admin test-e2e-install
	$(NPM) run test::auth-horusec::without-application-admin
test-e2e-auth-keycloak-without-application-admin: compose-e2e-auth-keycloak-without-application-admin test-e2e-install
	$(NPM) run test::auth-keycloak::without-application-admin

compose-e2e-auth-horusec-without-application-admin:
	$(DOCKER_COMPOSE) -f ./deployments/docker-compose.auth-horusec.without-application-admin.yaml down -v
	$(DOCKER_COMPOSE) -f ./deployments/docker-compose.auth-horusec.without-application-admin.yaml up -d --build --force-recreate
    make migrate-up
compose-e2e-auth-keycloak-without-application-admin:
	$(DOCKER_COMPOSE) -f ./deployments/docker-compose.auth-keycloak.without-application-admin.yaml down -v
	$(DOCKER_COMPOSE) -f ./deployments/docker-compose.auth-keycloak.without-application-admin.yaml up -d --build --force-recreate
    make migrate-up