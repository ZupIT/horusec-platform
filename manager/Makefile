HORUSEC ?= horusec

security:
    ifeq (, $(shell which $(HORUSEC)))
		    curl -fsSL https://raw.githubusercontent.com/ZupIT/horusec/master/deployments/scripts/install.sh | bash -s latest
		    $(HORUSEC) start -p="./" -e="true"
    else
		    $(HORUSEC) start -p="./" -e="true"
    endif
