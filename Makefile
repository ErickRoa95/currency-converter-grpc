.PHONY: server client

protoc:
	@echo "Running Protoc compiler : `protoc --version` " 
	@mkdir -p currencygrpc
	protoc --go_out=currencygrpc --go_opt=paths=source_relative \
    --go-grpc_out=currencygrpc --go-grpc_opt=paths=source_relative \
		./currencygrpc/currency.proto

server:
	@go run server/server.go

client: 
	@go run client/client.go 

mocks:
	@echo "== Generating Mocks for Interface | Mockgen `mockgen --version` =="
	@go generate ./...
