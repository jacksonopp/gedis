.PHONY: server client cli

server:
	@go run ./cmd/gedis/gedis.go

client:
	@go run ./cmd/client/client.go

cli:
	@go run ./cmd/cli/cli.go

demo: 
	@go run ./cmd/demo/demo.go