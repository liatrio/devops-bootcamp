# Example 1: Tightly Coupled Flask Application

This example demonstrates a simple Flask application where all concerns (presentation, business logic, and data access) are mixed together in a single file. This is a common pattern for quick prototypes, but it becomes difficult to maintain as applications grow.

## Structure

```
example1/
├── app.py              # Single file containing routes, logic, and data
├── requirements.txt    # Python dependencies
├── tests/
│   └── test_app.py    # Basic tests
└── README.md          # This file
```

## What This Example Shows

- **Tightly Coupled Code**: All functionality in one place
- **Direct Data Access**: Routes directly manipulate the in-memory data structure
- **Mixed Concerns**: HTTP handling, validation, and data storage all in the same functions

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
python app.py
```

The server will start on `http://127.0.0.1:5000`

### Test the API

```bash
# List all users
curl http://127.0.0.1:5000/users

# Create a new user
curl -X POST http://127.0.0.1:5000/users \
  -H "Content-Type: application/json" \
  -d '{"name": "Charlie"}'
```

### Run Tests

```bash
PYTHONPATH=. pytest -q
```

## Exercise: Add SQLite Storage

Try modifying this application to use SQLite instead of in-memory storage. You'll need to:

1. Import `sqlite3`
2. Create a database file and table
3. Modify `list_users()` to query the database
4. Modify `create_user()` to insert into the database
5. Handle connection management and errors

Notice how many places you need to touch and how database code gets mixed with HTTP handling code.

## Key Observations

- **Everything is in one file**: Easy to start with, but hard to navigate as it grows
- **No clear boundaries**: Where does the HTTP handling end and business logic begin?
- **Testing requires Flask**: You can't test validation logic without the web framework
- **Changes are risky**: Modifying data access might break HTTP handling

Compare this with Example 2 to see how layered architecture addresses these issues.
