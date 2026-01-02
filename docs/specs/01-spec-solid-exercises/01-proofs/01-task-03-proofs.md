# Task 3.0 Proof Artifacts: Dependency Inversion Principle (DIP) Project

## Summary

Exercise 3 (DIP) has been successfully created with all required components:
- Project structure with `pyproject.toml`, `.gitignore`, and tests directory
- `NotificationService` class demonstrating DIP violation (direct instantiation)
- Mock `SMTPClient` and `TwilioClient` classes for demonstration
- `main.py` demo script showing notifications through both channels
- Behavior-based test suite (10 tests)
- Complete README with setup instructions and dependency injection guidance

## CLI Output: main.py Execution

```
$ cd examples/ch11/solid-exercises/exercise-3-dip && uv run main.py
============================================================
Notification Service Demo
============================================================

User: Alice Smith
Email: alice@example.com
Phone: +1-555-123-4567

============================================================
Creating NotificationService (dependencies created internally)
============================================================
[SMTPClient] Initialized with smtp.example.com:587
[TwilioClient] Initialized with account: ACXXXXXXXXXX

============================================================
Sending notification through all channels
============================================================

--- Sending Email Notification ---
[SMTPClient] Sending email to: alice@example.com
[SMTPClient] Message: Your order #12345 has been shipped!
[SMTPClient] Email sent successfully!

--- Sending SMS Notification ---
[TwilioClient] Sending SMS to: +1-555-123-4567
[TwilioClient] Message: Your order #12345 has been shipped!
[TwilioClient] SMS sent successfully!

============================================================
Results
============================================================
  EMAIL: SUCCESS
  SMS: SUCCESS

============================================================
Why this violates the Dependency Inversion Principle
============================================================

The NotificationService class:
  - Creates SMTPClient() directly in __init__
  - Creates TwilioClient() directly in __init__

This means:
  1. We can't substitute different implementations
  2. We can't easily test without real clients
  3. Adding a new channel (Slack, Push) requires modifying the class
  4. High-level module depends on low-level implementations

To fix this:
  - Define abstract interfaces (MessageSender protocol)
  - Inject dependencies through the constructor
  - NotificationService should depend on abstractions, not concretions
```

## CLI Output: Test Suite Execution

```
$ cd examples/ch11/solid-exercises/exercise-3-dip && uv run --extra dev pytest
============================= test session starts ==============================
platform darwin -- Python 3.13.5, pytest-9.0.2, pluggy-1.6.0
rootdir: /Users/jburns/git/devops-bootcamp/examples/ch11/solid-exercises/exercise-3-dip
configfile: pyproject.toml
collected 10 items

tests/test_notification_service.py ..........                            [100%]

============================== 10 passed in 0.01s ==============================
```

## CLI Output: Type Checking (mypy)

```
$ cd examples/ch11/solid-exercises/exercise-3-dip && uv run --extra dev mypy .
Success: no issues found in 5 source files
```

## CLI Output: Linting (ruff)

```
$ cd examples/ch11/solid-exercises/exercise-3-dip && uv run --extra dev ruff check .
All checks passed!
```

## File Verification: README.md

The README exists at `examples/ch11/solid-exercises/exercise-3-dip/README.md` and contains:
- Exercise overview explaining DIP violation via direct instantiation
- Setup instructions (uv sync, uv run main.py, uv run pytest)
- Learning objectives about dependency injection and abstractions
- Suggested refactoring approach with Protocol-based design
- Code smells to watch for
- Types of dependency injection explained

## Project Structure

```
examples/ch11/solid-exercises/exercise-3-dip/
├── .gitignore
├── README.md
├── main.py
├── mock_clients.py
├── notification_service.py
├── pyproject.toml
├── tests/
│   ├── __init__.py
│   └── test_notification_service.py
└── uv.lock
```

## Verification Summary

| Proof Artifact | Status |
|----------------|--------|
| `uv run main.py` sends notifications | PASS |
| `uv run pytest` all tests pass (10 tests) | PASS |
| README.md exists with setup instructions | PASS |
| `uv run mypy .` type checking passes | PASS |
| `uv run ruff check .` linting passes | PASS |
