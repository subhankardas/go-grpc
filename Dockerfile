# Base image for golang build
FROM golang:1.22-alpine AS builder

# Set working directory
WORKDIR /app

# Copy and download dependency using go mod
COPY go.mod go.sum ./
RUN go mod download

# Copy the code into the build folder
COPY . .
RUN go mod vendor

# Build the service binary
RUN apk update
RUN apk add make
RUN make test build

# Final image with minimal dependencies
FROM alpine:latest

# Set the root folder
WORKDIR /root/

# Copy the binary and .env files
COPY --from=builder /app/build/user-service .
COPY --from=builder /app/.env .

# Expose port
EXPOSE 8080

# Run the binary
CMD ["./user-service"]