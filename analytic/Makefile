GO ?= go
GOFMT ?= gofmt
GOLANG_CI_LINT ?= golangci-lint
GO_FILES ?= $$(find . -name '*.go' | grep -v vendor)
GO_IMPORTS ?= goimports
GO_IMPORTS_LOCAL ?= github.com/ZupIT/horusec-platform/analytic/
HORUSEC ?= horusec
IMAGE_NAME ?= horuszup/horusec-analytic:local
GO_FUMPT ?= gofumpt
GO_GCI ?= gci
GO_LIST_TO_TEST ?= $$(go list ./...)

install-format-dependencies:
	$(GO) install golang.org/x/tools/cmd/goimports@latest
	$(GO) install mvdan.cc/gofumpt@latest
	$(GO) install github.com/daixiang0/gci@v0.2.9

format: install-format-dependencies
	$(GOFMT) -s -l -w $(GO_FILES)
	$(GO_IMPORTS) -w -local $(GO_IMPORTS_LOCAL) $(GO_FILES)
	$(GO_FUMPT) -l -w $(GO_FILES)
	$(GO_GCI) -w -local $(GO_IMPORTS_LOCAL) $(GO_FILES)

lint:
	$(GO) install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.43.0
	$(GOLANG_CI_LINT) run -v --timeout=5m -c .golangci.yml ./...

coverage:
	curl -fsSL https://raw.githubusercontent.com/ZupIT/horusec-devkit/main/scripts/coverage.sh | bash -s 100 .

test:
	$(GO) clean -testcache
	$(GO) test -v $(GO_LIST_TO_TEST) -race -timeout=5m -parallel=1 -failfast -short

security:
    ifeq (, $(shell which $(HORUSEC)))
		curl -fsSL https://raw.githubusercontent.com/ZupIT/horusec/master/deployments/scripts/install.sh | bash -s latest
		$(HORUSEC) start -p="./" -e="true"
    else
		$(HORUSEC) start -p="./" -e="true"
    endif

build:
	$(GO) build -o "./tmp/bin/analytic" ./cmd/app/main.go

update-swagger:
	curl -fsSL https://raw.githubusercontent.com/ZupIT/horusec-devkit/main/scripts/update-swagger.sh | bash -s "./cmd/app/main.go" .

pipeline: format lint test coverage build security

docker-build: ## Build docker image with the analytic.
	docker build -t ${IMAGE_NAME} -f ./deployments/dockerfiles/Dockerfile .
