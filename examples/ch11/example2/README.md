# Example 2: Layered Flask Application

This example demonstrates a Flask application organized using layered architecture. The code is separated into distinct layers: presentation (routes), business logic (services), and data access (repositories). This structure makes the code more maintainable, testable, and flexible.

## Structure

```
example2/
├── run.py                      # Application entry point
├── requirements.txt            # Python dependencies
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
# Create and activate a virtual environment
python -m venv venv
source venv/bin/activate  # On Windows: venv\Scripts\activate

# Install dependencies
pip install -r requirements.txt
```

### Run the Server

```bash
python run.py
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
PYTHONPATH=. pytest -q
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
- SQLite implementation (`sqlite_repo.py`)

## Exercise: Switch to SQLite Storage

This example makes it trivial to change the storage backend. Here's how:

### Step 1: Edit the Repository Interface

Open `app/data/__init__.py` and change the import:

```python
# Comment out the in-memory import
# from .in_memory_repo import get_all, add

# Uncomment the SQLite import
from .sqlite_repo import get_all, add
```

### Step 2: Run the Application

```bash
python run.py
```

That's it! The application now uses SQLite for storage. A `data.sqlite` file will be created in the example2 directory.

### What Happened?

- The presentation layer (routes) didn't change
- The business logic layer (services) didn't change
- Only the data layer changed (one import statement)

This demonstrates the power of layered architecture and dependency inversion.

## Key Observations

- **Clear Organization**: Easy to find where things are
- **Single Responsibility**: Each file has one job
- **Easy to Test**: You can test business logic without Flask or a database
- **Easy to Change**: Swapping implementations is trivial
- **Team Friendly**: Different developers can work on different layers

## Comparing with Example 1

Go back and try the same exercise in Example 1:

1. Example 1 requires editing route handlers directly
2. Example 1 mixes HTTP code with database code
3. Example 1 is harder to test in isolation
4. Example 2 required changing only one import statement

This is the benefit of layered architecture in practice.

## Extension Ideas

Try these exercises to deepen your understanding:

1. Add a `get_user_by_id(user_id)` function to all three layers
2. Add a `delete_user(user_id)` function
3. Create a third repository implementation that uses a JSON file
4. Add more business logic (e.g., name must be at least 2 characters)
5. Write unit tests for the service layer that mock the repository

## Best Practices Demonstrated

1. **Application Factory Pattern**: `create_app()` avoids circular imports
2. **Blueprint Registration**: Routes are organized in blueprints
3. **Repository Pattern**: Data access is abstracted behind an interface
4. **Explicit Dependencies**: Imports show the dependency flow
5. **Error Handling**: Business layer raises exceptions, presentation layer handles them
