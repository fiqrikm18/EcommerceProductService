include .env

GO_LINT=$(shell which golangci-lint 2> /dev/null || echo '')
GO_LINT_URI=github.com/golangci/golangci-lint/cmd/golangci-lint@latest

GO_SEC=$(shell which gosec 2> /dev/null || echo '')
GO_SEC_URI=github.com/securego/gosec/v2/cmd/gosec@latest

GO_VULNCHECK=$(shell which govulncheck 2> /dev/null || echo '')
GO_VULNCHECK_URI=golang.org/x/vuln/cmd/govulncheck@latest

GO_SWAGGO=$(shell which swag 2> /dev/null || echo '')
GO_SWAGGO_URI=github.com/swaggo/swag/cmd/swag@latest

.PHONY:
deps:
	@echo "##### Installing dependencies"
	$(if $(GO_LINT), ,go install $(GO_LINT_URI))
	$(if $(GO_SEC), ,go install $(GO_SEC_URI))
	$(if $(GO_VULNCHECK), ,go install $(GO_VULNCHECK_URI))
	$(if $(GO_SWAGGO), ,go install $(GO_SWAGGO_URI))
	@bash scripts/install-go-migrate.sh
	go mod download
	go mod tidy

.PHONY: golangci-lint
golangci-lint:
	$(if $(GO_LINT), ,go install $(GO_LINT_URI))
	@echo "##### Running golangci-lint"
	golangci-lint run -v --exclude 'docs/*'

.PHONY: gosec
gosec:
	$(if $(GO_SEC), ,go install $(GO_SEC_URI))
	@echo "##### Running gosec"
	gosec ./...

.PHONY: govulncheck
govulncheck:
	$(if $(GO_VULNCHECK), ,go install $(GO_VULNCHECK_URI))
	@echo "##### Running govulncheck"
	govulncheck ./...

.PHONY: verify
verify: deps generate-docs generate-mock golangci-lint gosec govulncheck dev

generate-docs:
	swag init -g cmd/app/main.go

.PHONY: test
test:
	@echo "##### Running tests"
	go test -race -cover -coverprofile=coverage.coverprofile -covermode=atomic -v ./...

.PHONY: dev
dev:
	go run ${APPLICATION_ROOT_CMD} --port ${APPLICATION_PORT}

generate-mock:
	@echo "##### Generating mock"
	go generate ./...

create-migration:
	migrate create -ext sql -dir db/migrations -seq $(seq)

run-migration:
	migrate -database ${POSTGRES_MIGRATION_DSN} -path db/migrations up

rollback-migration:
	migrate -database ${POSTGRES_MIGRATION_DSN} -path db/migrations down

docker-run:
	docker compose up -d

docker-down:
	docker compose down
