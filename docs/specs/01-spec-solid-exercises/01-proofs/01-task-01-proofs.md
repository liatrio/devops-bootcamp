# Task 1.0 Proof Artifacts: Single Responsibility Principle (SRP) Project

## Summary

Exercise 1 (SRP) has been successfully created with all required components:
- Project structure with `pyproject.toml`, `.gitignore`, and tests directory
- `ReportGenerator` class demonstrating SRP violation (4 responsibilities)
- `main.py` demo script
- Behavior-based test suite
- Complete README with setup instructions and learning objectives

## CLI Output: main.py Execution

```
$ cd examples/ch11/solid-exercises/exercise-1-srp && uv run main.py
Generating financial report...
Report generated successfully!
Output file: financial_report.html
HTML length: 1494 characters

To view the report, open 'financial_report.html' in a web browser.
```

## CLI Output: Test Suite Execution

```
$ cd examples/ch11/solid-exercises/exercise-1-srp && uv run --extra dev pytest
============================= test session starts ==============================
platform darwin -- Python 3.13.5, pytest-9.0.2, pluggy-1.6.0
rootdir: /Users/jburns/git/devops-bootcamp/examples/ch11/solid-exercises/exercise-1-srp
configfile: pyproject.toml
collected 17 items

tests/test_report_generator.py .................                         [100%]

============================== 17 passed in 0.01s ==============================
```

## CLI Output: Type Checking (mypy)

```
$ cd examples/ch11/solid-exercises/exercise-1-srp && uv run --extra dev mypy .
Success: no issues found in 4 source files
```

## CLI Output: Linting (ruff)

```
$ cd examples/ch11/solid-exercises/exercise-1-srp && uv run --extra dev ruff check .
All checks passed!
```

## File Verification: README.md

The README exists at `examples/ch11/solid-exercises/exercise-1-srp/README.md` and contains:
- Exercise overview explaining SRP violation
- Setup instructions (uv sync, uv run main.py, uv run pytest)
- Learning objectives about identifying and separating responsibilities
- Suggested refactoring approach
- Code smells to watch for
- Success criteria

## Project Structure

```
examples/ch11/solid-exercises/exercise-1-srp/
├── .gitignore
├── README.md
├── main.py
├── pyproject.toml
├── report_generator.py
├── tests/
│   ├── __init__.py
│   └── test_report_generator.py
└── uv.lock
```

## Verification Summary

| Proof Artifact | Status |
|----------------|--------|
| `uv run main.py` generates report | PASS |
| `uv run pytest` all tests pass (17 tests) | PASS |
| README.md exists with setup instructions | PASS |
| `uv run mypy .` type checking passes | PASS |
| `uv run ruff check .` linting passes | PASS |
