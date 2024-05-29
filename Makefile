# Install dependencies
dependencies:
	@echo Installing dependencies
	go mod download
	go mod vendor

# Generate gRPC files
generate-grpc:
	@echo Generating grpc files
	protoc --go_out=. --go_opt=paths=source_relative proto/*.proto \
    	--go-grpc_out=. --go-grpc_opt=paths=source_relative proto/*.proto

# Run server and client locally
run-local-server:
	@go run cmd/server/main.go

run-local-client:
	@go run cmd/client/main.go


# Build user service and run locally
run-build:
	@echo Building user service
	@go build -o build/user-service  ./cmd/server/main.go
	@./build/user-service

# Run tests
test:
	@echo Running tests
	@go test ./... -coverprofile cover.out

# Build user service for Docker
build:
	@echo Building user service for Docker
	@go build -o build/user-service ./cmd/server/main.go