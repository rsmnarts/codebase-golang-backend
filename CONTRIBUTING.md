# Contributing to Golang Fiber Backend

Thank you for your interest in contributing to this project! This document provides guidelines for contributing.

## Getting Started

1. Fork the repository
2. Clone your fork: `git clone https://github.com/your-username/codebase-golang-backend.git`
3. Create a new branch: `git checkout -b feature/your-feature-name`
4. Make your changes
5. Test your changes
6. Commit your changes: `git commit -m "Add your feature"`
7. Push to your fork: `git push origin feature/your-feature-name`
8. Create a Pull Request

## Development Setup

### Prerequisites

- Go 1.21 or higher
- Docker (optional)
- Make (optional, but recommended)

### Installation

```bash
# Clone the repository
git clone https://github.com/rsmnarts/codebase-golang-backend.git
cd codebase-golang-backend

# Install dependencies
go mod download

# Run the application
make run
```

## Code Style

- Follow standard Go conventions
- Run `make lint` before committing
- Use meaningful variable and function names
- Add comments for complex logic
- Keep functions small and focused

## Testing

```bash
# Run all tests
make test

# Run tests with coverage
go test -v -cover ./...
```

## Project Structure

```
.
├── cmd/api/              # Application entry point
├── internal/
│   ├── domain/           # Business entities and interfaces
│   ├── usecase/          # Business logic
│   ├── delivery/http/    # HTTP handlers
│   └── infrastructure/   # External implementations
└── pkg/                  # Shared packages
```

## Clean Architecture Guidelines

When contributing, please follow these Clean Architecture principles:

1. **Domain Layer**: 
   - No dependencies on other layers
   - Pure business logic
   - Define entities and repository interfaces

2. **Use Case Layer**:
   - Depends only on Domain layer
   - Implements business logic
   - Orchestrates data flow

3. **Delivery Layer**:
   - Handles HTTP requests/responses
   - Converts data formats
   - Depends on Use Case layer

4. **Infrastructure Layer**:
   - Implements repository interfaces
   - Handles external services
   - Depends on Domain layer

## Adding New Features

### Adding a New Entity

1. Create entity in `internal/domain/`
2. Define repository interface in domain layer
3. Create use case in `internal/usecase/`
4. Create handler in `internal/delivery/http/`
5. Implement repository in `internal/infrastructure/persistence/`
6. Add routes in `internal/delivery/http/routes.go`
7. Write tests for all layers

### Example: Adding a Product Entity

```go
// 1. internal/domain/product.go
type Product struct {
    ID    string
    Name  string
    Price float64
}

type ProductRepository interface {
    Create(product *Product) error
    GetByID(id string) (*Product, error)
    // ...
}

// 2. internal/usecase/product_usecase.go
type ProductUseCase struct {
    productRepo domain.ProductRepository
}

// 3. internal/delivery/http/product_handler.go
type ProductHandler struct {
    productUseCase *usecase.ProductUseCase
}

// 4. internal/infrastructure/persistence/product_repository.go
type InMemoryProductRepository struct {
    // implementation
}
```

## Pull Request Guidelines

- Keep PRs focused on a single feature or fix
- Write clear, descriptive commit messages
- Update documentation if needed
- Add tests for new features
- Ensure all tests pass
- Follow the existing code style

## Commit Message Format

```
<type>: <subject>

<body>

<footer>
```

Types:
- `feat`: New feature
- `fix`: Bug fix
- `docs`: Documentation changes
- `style`: Code style changes (formatting, etc.)
- `refactor`: Code refactoring
- `test`: Adding or updating tests
- `chore`: Maintenance tasks

Example:
```
feat: Add product management endpoints

- Add Product entity and repository interface
- Implement ProductUseCase with CRUD operations
- Add ProductHandler with HTTP endpoints
- Add tests for all layers

Closes #123
```

## Questions?

If you have questions, please open an issue or reach out to the maintainers.

## License

By contributing, you agree that your contributions will be licensed under the MIT License.
