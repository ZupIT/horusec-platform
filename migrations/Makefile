HORUSEC ?= horusec
IMAGE_NAME ?= horuszup/horusec-migrations:local

security:
    ifeq (, $(shell which $(HORUSEC)))
		curl -fsSL https://raw.githubusercontent.com/ZupIT/horusec/master/deployments/scripts/install.sh | bash -s latest
		$(HORUSEC) start -p="./" -e="true"
    else
		$(HORUSEC) start -p="./" -e="true"
    endif

build:
	docker build --no-cache -t "$(IMAGE_NAME)" -f ./deployments/dockerfiles/Dockerfile .
