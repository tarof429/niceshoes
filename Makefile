# A Self-Documenting Makefile: http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html

.PHONY: all build help test

all: build ## Build all targets

help: ## Display this help
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'


# build: ## Build all targets
# 	go build -o niceshoes
# 	cp niceshoes img

build:
	go build -o niceshoes
	cp niceshoes cmd

docker: ## Build docker image
	cd img; docker build -t tarof429/niceshoes:1  .

run-docker: ## Run docker image
	docker run -td --name niceshoes -v /mnt/iso:/mnt -v ./cmd:/cmd -v ./img/samples:/samples --rm  tarof429/niceshoes:1

install:
	go install

test: ## Run all tests
	go test -v ./...

clean:
	@rm -f niceshoes