.PHONY: lint build

DOCKER_FLAGS := --rm --volume ${PWD}:/go/src --workdir /go/src
ifeq ($(CI),)
	DOCKER_FLAGS += -it
endif

.PHONY: lint
lint:
	docker run $(DOCKER_FLAGS) \
		golangci/golangci-lint:v2.1.2 \
		golangci-lint run

.PHONY: lint-fix
lint-fix:
	docker run $(DOCKER_FLAGS) \
    		golangci/golangci-lint:v2.1.2 \
    		golangci-lint run --fix

.PHONY: build
build:
	docker run $(DOCKER_FLAGS) \
		--volume /var/run/docker.sock:/var/run/docker.sock \
		ghcr.io/viqueen/docker-images-golang:main \
		goreleaser build --clean --snapshot
