HORUSEC ?= horusec

security:
    ifeq (, $(shell which $(HORUSEC)))
		    curl -fsSL https://horusec.io/bin/install.sh | bash
		    $(HORUSEC) start -p="./" -e="true"
    else
		    $(HORUSEC) start -p="./" -e="true"
    endif
