PROJECT_NAME := "mbox"
PKG := "github.com/twgcode/$(PROJECT_NAME)"
.PHONY:  dep fmt gen help

dep: ## Get the dependencies
	@go mod tidy

fmt: ## Format the entire project
	@go fmt ./...

gen: ## Generate code
	@protoc -I=. -I=/usr/local/include --go_out=. --go_opt=module=${PKG} pb/*/*.proto
	@go fmt ./...
	@protoc-go-inject-tag -input=pb/*/*.pb.go

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'