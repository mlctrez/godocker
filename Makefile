
help:           ## Show this help.
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

build:	        ## Builds the docker image.
	@docker build --tag godocker .

run: build      ## Builds and then run the docker image.
	@echo "starting server at http://localhost:8080"
	@docker run --rm -p 8080:8080 godocker

rmi-dangle:     ## Removes any dangling docker images. DANGER this removes stuff
	@docker images -q --filter "dangling=true" | xargs docker rmi

.PHONY: build rmi-dangle