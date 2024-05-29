# Golang + gRPC

This repository contains a Go gRPC service for managing user details. The service includes functionalities to fetch user details, retrieve a list of user details based on a list of user IDs, and search for user details based on specific criteria.
**Assignment:** Developing a Golang gRPC User Service with Search

**Objective:** 
Your task in this assessment is to create a Golang gRPC service that provides specific functionalities for managing user details and includes a search capability. The primary objectives are as follows:
- Simulate a database by maintaining a list of user details within a variable.
- Create gRPC endpoints for fetching user details based on a user ID and retrieving a list of 
   user details based on a list of user IDs.
- Implement a search functionality to find user details based on specific criteria.

## Prerequisites

- Go (>= 1.16)
- Docker
- Docker Compose
- Protocol Buffers Compiler (`protoc`)

## Makefile Commands

### Install Dependencies

Install the required Go dependencies and vendor them.

```sh
make dependencies
```

### Generate gRPC Files

Generate the gRPC Go files from the `.proto` definitions.

```sh
make generate-grpc
```

### Run Server and Client Locally

Run the gRPC server locally.

```sh
make run-local-server
```

Run the gRPC client locally.

```sh
make run-local-client
```

### Build User Service and Run Locally

Build the user service binary and run it locally.

```sh
make run-build
```

### Run Tests

Run the tests for the project.

```sh
make test
```

### Build User Service for Docker

Build the user service binary for Docker.

```sh
make build
```

## Docker Commands

### Build Docker Image

Build the Docker image for the user service.

```sh
docker build -t user-service .
```

## Deployment

### Build and Run with Docker Compose

Build and run the user service using Docker Compose.

```sh
docker compose -f docker/docker-compose.yml up -d --build
```

### Stop Docker Compose Services

Stop the services started by Docker Compose.

```sh
docker compose -f docker/docker-compose.yml down
```

## Project Structure

```
go-grpc/
├── cmd/
│   └── client/
│       └── main.go
│   └── server/
│       └── main.go
├── docker/
├── proto/
├── src/
├── vendor/
├── .env
├── .gitignore
├── Dockerfile
├── go.mod
├── go.sum
├── Makefile
└── README.md
```

## Proto Definitions

The `.proto` files defining the gRPC services and messages are located in the `proto` directory.

## License

This project is licensed under the MIT License. See the LICENSE file for details.
