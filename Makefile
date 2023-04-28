BUILD_DEST_DIR ?= build
GIT_TAG=$(shell git rev-parse --abbrev-ref HEAD | grep HEAD >> /dev/null && git describe --tags || git rev-parse --abbrev-ref HEAD)
GIT_HASH=$(shell git rev-parse --short HEAD)
GIT_COMMIT_COUNT=$(shell git log --oneline|wc -l|sed s/[[:space:]]//g)
OS=$(shell uname -s)
ifeq ($(OS), Darwin)
	TIME=$(shell date +"%Y%m%d%H%M%S")
else
	TIME=$(shell date -d today +"%Y%m%d%H%M%S")
endif
VERSION=$(shell git rev-parse --abbrev-ref HEAD | grep HEAD >> /dev/null && git describe --tags || echo "`git rev-parse --abbrev-ref HEAD`.${GIT_HASH}.${GIT_COMMIT_COUNT}.${TIME}")

.PHONY: model
model:
	goctl model mysql ddl --database starbucks -d "sql/model" --src "sql/*.sql" -c

.PHONY: gen-go
gen-go:
	goctl api go -api starbucks.api -dir .

.PHONY: build run push
build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build

run:
	go run starbucks.go

.PHONY: swagger
swagger:
	goctl api plugin -plugin goctl-swagger="swagger -filename docs/starbucks-api.swagger.json" -api ./starbucks/starbucks.api -dir .

.PHONY: api-build
api-build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${BUILD_DEST_DIR}/starbucks-api starbucks.go