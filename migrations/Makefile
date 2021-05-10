HORUSEC ?= horusec
IMAGE_NAME ?= horuszup/horusec-:local

security:
    ifeq (, $(shell which $(HORUSEC)))
		curl -fsSL https://horusec.io/bin/install.sh | bash
		$(HORUSEC) start -p="./" -e="true"
    else
		$(HORUSEC) start -p="./" -e="true"
    endif

build:
	docker build --no-cache -t "$(IMAGE_NAME)" -f ./deployments/dockerfiles/Dockerfile .
