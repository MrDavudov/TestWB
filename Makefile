.PHONY: build
build:
	go build -v ./cmd/main.go
	./main

.PHONY: test
test:
	go test -v -race -cover -coverprofile=cover.out -timeout 30s ./...
	go tool cover -html=cover.out

.PHONY: gen
gen:
	mockgen -source=pkg/service/service.go \
	-destination=pkg/service/mocks/mock_service.go
	mockgen -source=pkg/repository/repository.go \
	-destination=pkg/repository/mocks/mock_repository.go

.DEFAULT_GOAL := build