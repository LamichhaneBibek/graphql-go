# Build stage
FROM golang:1.26-alpine AS builder

WORKDIR /app

# Set GOPATH and ensure /go/bin is in PATH
ENV GOPATH /go
ENV PATH $PATH:/go/bin

# Install gqlgen binary
RUN go install github.com/99designs/gqlgen@latest

# Copy go.mod and go.sum first for caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Ensure all project dependencies are downloaded
RUN go mod tidy

# Generate GraphQL code using the gqlgen binary
RUN gqlgen generate

# Build your server
RUN go build -o server ./cmd/server

# Final stage
FROM alpine:3.21

WORKDIR /app

COPY --from=builder /app/server .
COPY --from=builder /app/graph/ ./graph/

EXPOSE 8080

CMD ["./server"]