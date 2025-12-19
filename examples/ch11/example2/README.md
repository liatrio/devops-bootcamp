# Example 2: Layered Flask Application

This example demonstrates a Flask application organized using layered architecture. The code is separated into distinct layers: presentation (routes), business logic (services), and data access (repositories). This structure makes the code more maintainable, testable, and flexible.

## Structure

```text
example2/
├── run.py                      # Application entry point
├── pyproject.toml              # Python dependencies (uv project)
├── app/
│   ├── __init__.py            # Application factory
│   ├── presentation/
│   │   └── routes.py          # HTTP routes (presentation layer)
│   ├── services/
│   │   └── user_service.py    # Business logic (service layer)
│   ├── data/
│   │   ├── __init__.py        # Repository interface (single swap point)
│   │   ├── in_memory_repo.py  # In-memory implementation
│   │   └── sqlite_repo.py     # SQLite implementation
│   └── models/
│       └── user.py            # Data models (placeholder)
├── tests/
│   └── test_routes.py         # Tests
└── README.md                  # This file
```

## What This Example Shows

- **Separation of Concerns**: Each layer has a single responsibility
- **Clear Boundaries**: Presentation doesn't know about data storage, business logic doesn't know about HTTP
- **Easy to Test**: Business logic can be tested independently
- **Easy to Change**: Switching storage backends requires changing one import

## Running the Application

### Setup

```bash
# Install dependencies (uv will automatically create and manage the virtual environment)
uv sync
```

### Run the Server

```bash
uv run python run.py
```

The server will start on `http://127.0.0.1:5001`

### Test the API

```bash
# List all users
curl http://127.0.0.1:5001/users

# Create a new user
curl -X POST http://127.0.0.1:5001/users \
  -H "Content-Type: application/json" \
  -d '{"name": "Charlie"}'
```

### Run Tests

```bash
uv run pytest -q
```

## Understanding the Layers

### Presentation Layer (`app/presentation/routes.py`)

Handles HTTP concerns:
- Parsing requests
- Calling business logic functions
- Formatting responses
- HTTP status codes

### Business Logic Layer (`app/services/user_service.py`)

Contains business rules:
- Input validation
- Business logic (none in this simple example, but this is where it would go)
- Calls data layer for persistence

### Data Access Layer (`app/data/`)

Manages data storage:
- Repository interface (`__init__.py`)
- In-memory implementation (`in_memory_repo.py`)

This is the benefit of layered architecture in practice.

## Best Practices Demonstrated

1. **Application Factory Pattern**: `create_app()` avoids circular imports
2. **Blueprint Registration**: Routes are organized in blueprints
3. **Repository Pattern**: Data access is abstracted behind an interface
4. **Explicit Dependencies**: Imports show the dependency flow
5. **Error Handling**: Business layer raises exceptions, presentation layer handles them
