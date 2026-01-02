# Task 2.0 Proof Artifacts: Open/Closed Principle (OCP) Project

## Summary

Exercise 2 (OCP) has been successfully created with all required components:
- Project structure with `pyproject.toml`, `.gitignore`, and tests directory
- `DiscountCalculator` class demonstrating OCP violation (if/elif chains)
- `main.py` demo script showing discounts for all customer types
- Behavior-based test suite (16 tests)
- Complete README with setup instructions and polymorphism guidance

## CLI Output: main.py Execution

```
$ cd examples/ch11/solid-exercises/exercise-2-ocp && uv run main.py
============================================================
Discount Calculator Demo
============================================================

Customer Type: REGULAR
----------------------------------------
  Original: $  100.00
  Discount: $    5.00
  Final:    $   95.00

  Original: $  250.00
  Discount: $   12.50
  Final:    $  237.50

  Original: $  500.00
  Discount: $   25.00
  Final:    $  475.00

  Original: $ 1000.00
  Discount: $   50.00
  Final:    $  950.00


Customer Type: PREMIUM
----------------------------------------
  Original: $  100.00
  Discount: $   10.00
  Final:    $   90.00

  Original: $  250.00
  Discount: $   25.00
  Final:    $  225.00

  Original: $  500.00
  Discount: $   50.00
  Final:    $  450.00

  Original: $ 1000.00
  Discount: $  100.00
  Final:    $  900.00


Customer Type: VIP
----------------------------------------
  Original: $  100.00
  Discount: $   15.00
  Final:    $   85.00

  Original: $  250.00
  Discount: $   37.50
  Final:    $  212.50

  Original: $  500.00
  Discount: $   75.00
  Final:    $  425.00

  Original: $ 1000.00
  Discount: $  150.00
  Final:    $  850.00


What happens when we need a new customer type?
----------------------------------------
Try adding a 'platinum' customer type...
You'll need to modify the DiscountCalculator class!

This violates the Open/Closed Principle:
  - The class is NOT closed for modification
  - Every new type requires changing existing code
```

## CLI Output: Test Suite Execution

```
$ cd examples/ch11/solid-exercises/exercise-2-ocp && uv run --extra dev pytest
============================= test session starts ==============================
platform darwin -- Python 3.13.5, pytest-9.0.2, pluggy-1.6.0
rootdir: /Users/jburns/git/devops-bootcamp/examples/ch11/solid-exercises/exercise-2-ocp
configfile: pyproject.toml
collected 16 items

tests/test_discount_calculator.py ................                       [100%]

============================== 16 passed in 0.01s ==============================
```

## CLI Output: Type Checking (mypy)

```
$ cd examples/ch11/solid-exercises/exercise-2-ocp && uv run --extra dev mypy .
Success: no issues found in 4 source files
```

## CLI Output: Linting (ruff)

```
$ cd examples/ch11/solid-exercises/exercise-2-ocp && uv run --extra dev ruff check .
All checks passed!
```

## File Verification: README.md

The README exists at `examples/ch11/solid-exercises/exercise-2-ocp/README.md` and contains:
- Exercise overview explaining OCP violation via if/elif chains
- Setup instructions (uv sync, uv run main.py, uv run pytest)
- Learning objectives about polymorphism and abstract base classes
- Suggested refactoring approach with code examples
- Code smells to watch for

## Project Structure

```
examples/ch11/solid-exercises/exercise-2-ocp/
├── .gitignore
├── README.md
├── discount_calculator.py
├── main.py
├── pyproject.toml
├── tests/
│   ├── __init__.py
│   └── test_discount_calculator.py
└── uv.lock
```

## Verification Summary

| Proof Artifact | Status |
|----------------|--------|
| `uv run main.py` calculates discounts | PASS |
| `uv run pytest` all tests pass (16 tests) | PASS |
| README.md exists with setup instructions | PASS |
| `uv run mypy .` type checking passes | PASS |
| `uv run ruff check .` linting passes | PASS |
