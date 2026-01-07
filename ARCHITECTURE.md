# Clean Architecture Diagram

```
┌─────────────────────────────────────────────────────────────┐
│                      External Interface                      │
│                    (HTTP Requests/Responses)                 │
└─────────────────────────┬───────────────────────────────────┘
                          │
┌─────────────────────────▼───────────────────────────────────┐
│                    Delivery Layer                            │
│                (internal/delivery/http)                      │
│                                                               │
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐         │
│  │   Routes    │  │  Handlers   │  │ Middleware  │         │
│  └─────────────┘  └─────────────┘  └─────────────┘         │
└─────────────────────────┬───────────────────────────────────┘
                          │
┌─────────────────────────▼───────────────────────────────────┐
│                     Use Case Layer                           │
│                  (internal/usecase)                          │
│                                                               │
│  ┌──────────────────────────────────────────────────┐       │
│  │          Business Logic / Use Cases              │       │
│  │  - CreateUser, GetUser, UpdateUser, etc.        │       │
│  └──────────────────────────────────────────────────┘       │
└─────────────────────────┬───────────────────────────────────┘
                          │
┌─────────────────────────▼───────────────────────────────────┐
│                      Domain Layer                            │
│                   (internal/domain)                          │
│                                                               │
│  ┌─────────────┐  ┌─────────────────────────────┐          │
│  │  Entities   │  │  Repository Interfaces      │          │
│  │   (User)    │  │    (UserRepository)         │          │
│  └─────────────┘  └─────────────────────────────┘          │
└──────────────────────────────────┬──────────────────────────┘
                                   │
┌──────────────────────────────────▼──────────────────────────┐
│                  Infrastructure Layer                        │
│            (internal/infrastructure/persistence)             │
│                                                               │
│  ┌────────────────────────────────────────────────┐         │
│  │      Repository Implementations                │         │
│  │    (InMemoryUserRepository, etc.)              │         │
│  └────────────────────────────────────────────────┘         │
└─────────────────────────┬───────────────────────────────────┘
                          │
┌─────────────────────────▼───────────────────────────────────┐
│                    External Services                         │
│              (Database, Cache, APIs, etc.)                   │
└─────────────────────────────────────────────────────────────┘

Dependency Flow: Delivery → Use Case → Domain ← Infrastructure
```

## Key Principles

1. **Dependency Rule**: Dependencies point inward. Inner layers know nothing about outer layers.

2. **Entities (Domain)**: Enterprise business rules. The most stable layer.

3. **Use Cases**: Application-specific business rules. Orchestrate data flow.

4. **Interface Adapters (Delivery)**: Convert data between use cases and external world.

5. **Frameworks & Drivers (Infrastructure)**: External tools and frameworks.

## Benefits

- **Testability**: Each layer can be tested independently
- **Independence**: Business logic doesn't depend on frameworks
- **Flexibility**: Easy to swap implementations (e.g., database changes)
- **Maintainability**: Clear separation of concerns
- **Scalability**: Easy to add new features without affecting existing code
