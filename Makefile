.PHONY: lint build

.PHONY: lint
lint:
	docker run -it --rm \
		--volume "${PWD}":/go/src \
		--workdir /go/src \
		golangci/golangci-lint:v2.1.2 \
		golangci-lint run

.PHONY: lint-fix
lint-fix:
	docker run -it --rm \
    		--volume "${PWD}":/go/src \
    		--workdir /go/src \
    		golangci/golangci-lint:v2.1.2 \
    		golangci-lint run --fix

.PHONY: build
build:
	docker run -it --rm \
		--volume "${PWD}":/go/src \
		--volume /var/run/docker.sock:/var/run/docker.sock \
		--workdir /go/src \
		ghcr.io/viqueen/docker-images-golang:main \
		goreleaser build --clean --snapshot
