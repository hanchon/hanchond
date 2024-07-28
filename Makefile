.phony: build docs

dev-docs:
	@source /opt/homebrew/opt/nvm/nvm.sh && nvm use && npm run docs:dev

build-docs:
	@source /opt/homebrew/opt/nvm/nvm.sh && nvm use && npm run docs:build

generate:
	@sqlc generate

install-deps:
	@go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

lint:
	golangci-lint run --fix --out-format=line-number --issues-exit-code=0 --config .golangci.yml --color always ./...

