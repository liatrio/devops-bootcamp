# Pre-Commit Flask

A simple Flask application with comprehensive Python development tools configured for UV.

## Setup

Install dependencies using UV:

```bash
# Install all dependencies including dev tools
uv sync

# Or install only dev dependencies
uv sync --all-extras
```

## Development Tools

This project uses the following tools for code quality and testing:

### Testing with pytest

Run all tests:
```bash
uv run pytest
```

Run tests with verbose output:
```bash
uv run pytest -v
```

### Code Formatting with Black

Format all Python files:
```bash
uv run black .
```

Check formatting without making changes:
```bash
uv run black --check .
```

### Import Sorting with isort

Sort imports in all files:
```bash
uv run isort .
```

Check import sorting without making changes:
```bash
uv run isort --check-only .
```

### Linting with Ruff

Run Ruff linter:
```bash
uv run ruff check .
```

Auto-fix issues where possible:
```bash
uv run ruff check --fix .
```

Format code with Ruff (alternative to Black):
```bash
uv run ruff format .
```

### Type Checking with mypy

Run type checking:
```bash
uv run mypy src/
```

Run type checking on all files:
```bash
uv run mypy .
```

## Running the Application

Start the Flask development server:

```bash
uv run python src/app.py
```

Or with Flask's CLI:
```bash
uv run flask --app src.app run
```

The application will be available at `http://localhost:8000/health`
