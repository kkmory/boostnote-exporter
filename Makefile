# ANSI color
RED=\033[31m
GREEN=\033[32m
RESET=\033[0m

COLORIZE_PASS=sed ''/PASS/s//$$(printf "$(GREEN)PASS$(RESET)")/''
COLORIZE_FAIL=sed ''/FAIL/s//$$(printf "$(RED)FAIL$(RESET)")/''

# Const for project
GO_MOD_PROJECT='github.com/kkmory/boostnote-exporter'
TEST_FOLDER='test'

.PHONY: up
up: ## ローカルで実行
	go run cmd/exporter/main.go -folder $(TEST_FOLDER)

.PHONY: prepare_env
prepare_env: ## 開発に必要なパッケージ類をダウンロードします。
	go install github.com/tommy-sho/grouper@latest
	go mod download

.PHONY: fmt
fmt: ## ソースコードをフォーマットします | grouper, go fmt, go vet を実行
	find `pwd`/ -name "*.go" | grep -v mock | xargs -L 1 grouper -local $(GO_MOD_PROJECT) -w
	go vet ./... && \
	go fmt ./...

.PHONY: test
test: ## テストを実行する
	go test -v ./... | $(COLORIZE_PASS) | $(COLORIZE_FAIL)

.PHONY: clean_cache
clean_cache: ## テストキャッシュを消去する
	go clean -testcache

.PHONY: help
help: ## ヘルプ
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-16s\033[0m %s\n", $$1, $$2}'

.PHONY: tidy
tidy: ## go mod tidy
	go mod tidy
