LOCAL_BIN=$(CURDIR)/bin
PROJECT_NAME=metrics

export GO111MODULE=on
GOENV:= GO111MODULE=on

.PHONY: build
build:
	$(GOENV) CGO_ENABLED=0 go build -v -o $(LOCAL_BIN)/$(PROJECT_NAME) ./cmd/app

.PHONY: run
run:
	$(GOENV) go run cmd/main.go $(RUN_ARGS)

.PHONY: up
up:
	docker-compose up -d

.PHONY: down
down:
	docker-compose down