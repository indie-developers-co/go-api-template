SHELL := /bin/bash
PROJECT_NAME := "go-api-echo-template"
PKG_LIST := $(shell go list ./... | grep -v /vendor/)
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/ | grep -v _test.go)
TEST_COVERAGE_THRESHOLD = 100

.PHONY: build, test, test_coverage, test_coverage_html, test_coverage_xml, run, tidy, fmt, vet, race, dep, imports, infra, infra-down, runner-test

build: dep
	@go build -o main cmd/app/main.go

test:
	@go test ./... -short ${PKG_LIST}

test_coverage:
	@echo "Current coverage threshold 	: ${TEST_COVERAGE_THRESHOLD} %"
	@go test ./... -coverprofile=coverage.out
	@coverage=$$(go tool cover -func=coverage.out | grep total | grep -Eo '[0-9]+\.[0-9]+') ;\
	 if [ $$(bc <<< "$$coverage < $(TEST_COVERAGE_THRESHOLD)") -eq 1 ]; then \
		echo "Current test coverage is below threshold: $$coverage < $(TEST_COVERAGE_THRESHOLD)" ;\
		echo "Failed";\
		exit 1 ;\
	else \
		echo "Current test coverage		: $$coverage %" ;\
	  	echo "OK";\
	fi

test_coverage_html:
	@go test ./... -coverprofile=coverage.out && go tool cover -html=coverage.out
	@rm coverage.out

test_coverage_xml: test_coverage xml

xml:
	@if [ ! -f "coverage.out" ]; then \
  		echo "coverage.out file does not found, please generate this file first";\
  		exit 1;\
  	fi
	@gocover-cobertura < coverage.out > coverage.xml
	@rm coverage.out ;\

run:
	@go run cmd/app/main.go

tidy:
	@go mod tidy

fmt: tidy
	@go fmt ./... ${PKG_LIST}

lint: tidy
	@staticcheck ./... ${PKG_LIST}
	@go vet ./... ${PKG_LIST}
	@gocritic check ./... ${PKG_LIST}

race: tidy
	@go test ./... -race -short ${PKG_LIST}

dep: tidy
	go mod download
	go install golang.org/x/tools/cmd/goimports@latest
	go install honnef.co/go/tools/cmd/staticcheck@latest
	go install github.com/boumenot/gocover-cobertura@latest
	go install github.com/go-critic/go-critic/cmd/gocritic@latest
	go install github.com/vektra/mockery/v2@v2.20.0

imports:
	@goimports -l ./..

infra:
	@docker-compose up -d

infra-down:
	@docker-compose down --remove-orphans

runner-test:
	@gitlab-runner exec docker test

mocks:
	mockery --all --dir internal --output ./tests/mocks --case underscore