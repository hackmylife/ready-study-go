.PHONY: help list check lint diff progress next hint fmt fmt-check vet verify verify-starters test db-up db-down

EX ?= 01-language/01-variables
HINT ?= 1

help:
	@go run ./cmd/koans

list:
	go run ./cmd/koans list

check:
	go run ./cmd/koans check "$(EX)"

lint:
	go run ./cmd/koans lint "$(EX)"

diff:
	go run ./cmd/koans diff "$(EX)"

test: check

progress:
	go run ./cmd/koans progress

next:
	go run ./cmd/koans next

hint:
	go run ./cmd/koans hint "$(EX)" "$(HINT)"

fmt:
	gofmt -w .

fmt-check:
	@files=$$(gofmt -l .) && { test -z "$$files" || { echo "$$files"; exit 1; }; }

vet:
	go vet ./...

verify:
	go run ./cmd/koans verify

verify-starters:
	go run ./cmd/koans verify --starters

db-up:
	docker compose -p ready-study-go-koans up -d --wait

db-down:
	docker compose -p ready-study-go-koans down
