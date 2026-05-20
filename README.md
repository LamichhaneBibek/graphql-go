# GraphQL User Authentication API

A clean and simple GraphQL API for user authentication and management, built with Go, gqlgen, and PostgreSQL.

## Features

- **User Registration (Sign Up)**: Create new user accounts with email and password
- **User Login (Sign In)**: Authenticate and receive JWT tokens
- **JWT Authentication**: Secure token-based authentication
- **User Profile**: Retrieve current user profile (`me` query)
- **User Management**: List all users and fetch specific user details

## Architecture

- **Models**: Simple User model with bcrypt password hashing
- **Services**: Auth and User services with clear separation of concerns
- **Repository Pattern**: Abstracted database access layer
- **GraphQL Resolvers**: Clean resolver implementations
- **JWT Middleware**: Automatic token verification on authenticated endpoints

## Getting Started

### Prerequisites

- Go 1.26+
- PostgreSQL
- Docker & Docker Compose (optional)

### Environment Setup

Configure your `.env` file:

```env
PORT=8080
DATABASE_URL=postgresql://user:password@localhost:5432/dbname
JWT_SECRET=your-secret-key
```

### Installation

```bash
# Install dependencies
go mod download
go mod tidy

# Generate GraphQL code
go run github.com/99designs/gqlgen generate

# Build the project
go build -v ./cmd/server/

# Run the server
./server
```

## GraphQL API Usage

### 1. Register (Sign Up)

```graphql
mutation {
  register(input: {
    name: "John Doe"
    email: "john@example.com"
    password: "securepassword123"
  }) {
    token
    user {
      id
      name
      email
      createdAt
    }
  }
}
```

### 2. Login (Sign In)

```graphql
mutation {
  login(input: {
    email: "john@example.com"
    password: "securepassword123"
  }) {
    token
    user {
      id
      name
      email
      createdAt
    }
  }
}
```

### 3. Get Current User Profile

```graphql
query {
  me {
    id
    name
    email
    createdAt
  }
}
```

**Note**: Include the JWT token in the Authorization header:
```
Authorization: Bearer <token>
```

### 4. Get All Users

```graphql
query {
  users {
    id
    name
    email
    createdAt
  }
}
```

**Note**: Requires authentication

### 5. Get Specific User

```graphql
query {
  user(id: "1") {
    id
    name
    email
    createdAt
  }
}
```

**Note**: Requires authentication

## Commands

```bash
# Start with Docker
docker compose up --build

# Run container shell
docker exec -it go_graphql_api sh

# Stop containers
docker compose down

# Regenerate GraphQL code after schema changes
go run github.com/99designs/gqlgen generate

# Format code
go fmt ./...

# Build
go build -v ./cmd/server/
```

## Project Structure

```
.
├── cmd/server/           # Server entry point
├── config/              # Configuration (DB, env)
├── graph/               # GraphQL schema, resolvers, mappers
├── internal/
│   ├── auth/            # JWT token generation and middleware
│   ├── models/          # Database models (User)
│   ├── repository/      # Data access layer
│   ├── seed/            # Database seeding (if needed)
│   └── service/         # Business logic (Auth, User)
├── docker-compose.yml   # Docker compose configuration
├── go.mod               # Go module definition
└── gqlgen.yml           # GraphQL codegen config
```

## Cleanup Summary

This version has been cleaned up by removing:
- ✅ Role and Permission system (unnecessary complexity)
- ✅ Post management features (out of scope)
- ✅ Role-based access control mutations
- ✅ Unused service layer complexity

Kept and streamlined:
- ✅ Clean User model
- ✅ Secure authentication (bcrypt + JWT)
- ✅ Simple and focused API
- ✅ Repository pattern for data access
- ✅ Service layer for business logic