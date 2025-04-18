.PHONY: mocks lint build

.PHONY: mocks
mocks:
	docker run --rm \
		--volume "${PWD}":/go/src \
		--workdir /go/src \
		vektra/mockery:v2.46 \
		--all

.PHONY: lint
lint:
	docker run --rm \
		--volume "${PWD}":/go/src \
		--workdir /go/src \
		golangci/golangci-lint:v2.1.2 \
		golangci-lint run

.PHONY: lint-fix
lint-fix:
	docker run --rm \
    		--volume "${PWD}":/go/src \
    		--workdir /go/src \
    		golangci/golangci-lint:v2.1.2 \
    		golangci-lint run --fix

.PHONY: build
build:
	docker run --rm \
		--volume "${PWD}":/go/src \
		--volume /var/run/docker.sock:/var/run/docker.sock \
		--workdir /go/src \
		goreleaser/goreleaser:v2.4.4 \
		build --clean --snapshot
