# Build stage
FROM golang:1.26-alpine AS builder

WORKDIR /app

# Set GOPATH and PATH with modern syntax
ENV GOPATH=/go
ENV PATH=/go/bin:/usr/local/go/bin:$PATH

# Install gqlgen binary
RUN go install github.com/99designs/gqlgen@latest

# Copy go.mod and go.sum first for caching
COPY go.mod go.sum ./
RUN go mod download && go mod tidy

# Copy the rest of the source code
COPY . .

# Generate GraphQL code
RUN gqlgen generate

# Build the server
RUN go build -o server ./cmd/server

# Final stage
FROM alpine:3.21

WORKDIR /app

# Install runtime dependencies
RUN apk add --no-cache ca-certificates tzdata

# Copy the server binary and GraphQL files
COPY --from=builder /app/server .
COPY --from=builder /app/graph/ ./graph/

# Runtime command
CMD ["./server"]
