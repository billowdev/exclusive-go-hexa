.PHONY: run init tidy build

# Dependencies
MODULES := $(wildcard cmd/**/*.go internal/**/*.go pkg/**/*.go)

# Initialize project by tidying up dependencies and running the application
init: tidy run

# Tidy up Go modules
tidy:
	go mod tidy

# Build the Go application
build:
	go build -o bin/main ./cmd/main.go

# Run the Go application
run: $(MODULES)
	go run ./cmd/main.go

test:
	go test ./... -cover

gql-gen:
	go get github.com/99designs/gqlgen
	go run github.com/99designs/gqlgen generate


grpc-server-gen:
	mkdir -p internal/adapters/grpc_server/pb
	protoc \
		--proto_path=internal/adapters/grpc_server/proto  \
		--go_out=internal/adapters/grpc_server/pb \
		--go_opt=paths=source_relative \
		--go-grpc_out=internal/adapters/grpc_server/pb \
		--go-grpc_opt=paths=source_relative \
		internal/adapters/grpc_server/proto/*.proto

grpc-client-gen:
	mkdir -p internal/adapters/grpc_server/pb
	protoc \
		--proto_path=internal/adapters/grpc_server/proto  \
		--go_out=internal/adapters/grpc_server/pb \
		--go_opt=paths=source_relative \
		--go-grpc_out=internal/adapters/grpc_server/pb \
		--go-grpc_opt=paths=source_relative \
		internal/adapters/grpc_server/proto/*.proto
