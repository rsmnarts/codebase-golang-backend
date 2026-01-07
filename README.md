# Golang Fiber Backend with Clean Architecture

A simple and scalable REST API backend built with Go (Golang) using the Fiber framework, following Clean Architecture principles.

## 🏗️ Architecture

This project follows **Clean Architecture** principles with clear separation of concerns:

```
.
├── cmd/
│   └── api/                    # Application entry point
│       └── main.go
├── internal/
│   ├── domain/                 # Business entities and repository interfaces
│   │   ├── user.go
│   │   └── errors.go
│   ├── usecase/                # Business logic / use cases
│   │   └── user_usecase.go
│   ├── delivery/               # Delivery mechanisms (HTTP, gRPC, etc.)
│   │   └── http/
│   │       ├── user_handler.go
│   │       └── routes.go
│   └── infrastructure/         # External implementations
│       └── persistence/
│           └── user_repository.go
└── pkg/
    ├── config/                 # Configuration management
    │   └── config.go
    └── middleware/             # HTTP middleware
        └── middleware.go
```

### Architecture Layers

1. **Domain Layer** (`internal/domain`): Contains business entities and repository interfaces. This is the core of the application and has no dependencies on other layers.

2. **Use Case Layer** (`internal/usecase`): Contains business logic and orchestrates data flow. Depends only on the domain layer.

3. **Delivery Layer** (`internal/delivery`): Contains HTTP handlers (or other delivery mechanisms). Translates HTTP requests into use case calls.

4. **Infrastructure Layer** (`internal/infrastructure`): Contains implementations of repository interfaces and external service integrations.

## 🚀 Features

- Clean Architecture implementation
- RESTful API with Fiber framework
- In-memory data storage (easily replaceable with database)
- CRUD operations for User entity
- Middleware support (CORS, Logger, Recovery)
- Docker support
- Environment-based configuration

## 📋 Prerequisites

- Go 1.21 or higher
- Docker and Docker Compose (optional)

## 🔧 Installation

### Local Setup

1. Clone the repository:
```bash
git clone https://github.com/rsmnarts/codebase-golang-backend.git
cd codebase-golang-backend
```

2. Install dependencies:
```bash
go mod download
```

3. Run the application:
```bash
go run cmd/api/main.go
```

Or using Make:
```bash
make run
```

The server will start on `http://localhost:8080`

### Docker Setup

1. Build and run with Docker Compose:
```bash
docker-compose up --build
```

2. Or build and run manually:
```bash
docker build -t golang-fiber-backend .
docker run -p 8080:8080 golang-fiber-backend
```

## 📚 API Endpoints

### Health Check
- `GET /health` - Check if the API is running

### User Management

- `POST /api/v1/users` - Create a new user
  ```json
  {
    "name": "John Doe",
    "email": "john@example.com"
  }
  ```

- `GET /api/v1/users` - Get all users

- `GET /api/v1/users/:id` - Get a user by ID

- `PUT /api/v1/users/:id` - Update a user
  ```json
  {
    "name": "Jane Doe",
    "email": "jane@example.com"
  }
  ```

- `DELETE /api/v1/users/:id` - Delete a user

## 🧪 Example Usage

### Create a user
```bash
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{"name":"John Doe","email":"john@example.com"}'
```

### Get all users
```bash
curl http://localhost:8080/api/v1/users
```

### Get a specific user
```bash
curl http://localhost:8080/api/v1/users/{user-id}
```

### Update a user
```bash
curl -X PUT http://localhost:8080/api/v1/users/{user-id} \
  -H "Content-Type: application/json" \
  -d '{"name":"Jane Doe","email":"jane@example.com"}'
```

### Delete a user
```bash
curl -X DELETE http://localhost:8080/api/v1/users/{user-id}
```

## ⚙️ Configuration

The application can be configured using environment variables:

- `PORT` - Server port (default: 8080)
- `APP_NAME` - Application name (default: "Golang Fiber Backend")

Example:
```bash
export PORT=3000
export APP_NAME="My Custom API"
go run cmd/api/main.go
```

## 🏗️ Project Structure Explained

### Clean Architecture Benefits

1. **Independence of Frameworks**: The business logic doesn't depend on Fiber or any other framework.
2. **Testability**: Business logic can be tested without UI, database, or external services.
3. **Independence of UI**: The delivery mechanism can be changed without affecting business logic.
4. **Independence of Database**: The in-memory repository can be replaced with any database implementation.
5. **Maintainability**: Clear separation of concerns makes the code easier to maintain and extend.

### Dependency Rule

Dependencies flow inward: `Delivery -> Use Case -> Domain`

- Domain layer has no dependencies
- Use Case layer depends only on Domain
- Delivery and Infrastructure layers depend on Use Case and Domain

## 🛠️ Development Commands

The project includes a Makefile for common development tasks:

```bash
make help          # Show all available commands
make build         # Build the application
make run           # Run the application
make test          # Run tests
make clean         # Remove build artifacts
make docker-build  # Build Docker image
make docker-run    # Run Docker container
make lint          # Run linters (fmt + vet)
```

## 🔄 Extending the Application

### Adding a New Entity

1. Create entity in `internal/domain/`
2. Create repository interface in domain layer
3. Create use case in `internal/usecase/`
4. Create handler in `internal/delivery/http/`
5. Create repository implementation in `internal/infrastructure/persistence/`
6. Register routes in `internal/delivery/http/routes.go`

### Replacing In-Memory Storage with Database

1. Create a new repository implementation in `internal/infrastructure/persistence/`
2. Update `cmd/api/main.go` to use the new repository
3. No changes needed in domain, use case, or delivery layers!

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## 📝 License

This project is open source and available under the [MIT License](LICENSE).

## 🙏 Acknowledgments

- [Fiber Framework](https://gofiber.io/)
- [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html) by Robert C. Martin