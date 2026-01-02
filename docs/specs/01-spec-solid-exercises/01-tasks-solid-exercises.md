# 01-tasks-solid-exercises.md

## Relevant Files

### Exercise 1: SRP Project

- `examples/ch11/solid-exercises/exercise-1-srp/pyproject.toml` - UV project configuration with Python 3.11+, pytest, mypy, ruff dependencies
- `examples/ch11/solid-exercises/exercise-1-srp/main.py` - Entry point demonstrating ReportGenerator with multiple responsibilities (validation, calculation, formatting, file I/O)
- `examples/ch11/solid-exercises/exercise-1-srp/report_generator.py` - Bad implementation of ReportGenerator class violating SRP
- `examples/ch11/solid-exercises/exercise-1-srp/tests/test_report_generator.py` - Test suite validating report generation behavior (works for both bad and refactored implementations)
- `examples/ch11/solid-exercises/exercise-1-srp/README.md` - Exercise documentation with setup instructions and learning objectives
- `examples/ch11/solid-exercises/exercise-1-srp/.gitignore` - Python gitignore (venv, __pycache__, *.pyc, report files)

### Exercise 2: OCP Project

- `examples/ch11/solid-exercises/exercise-2-ocp/pyproject.toml` - UV project configuration with Python 3.11+, pytest, mypy, ruff dependencies
- `examples/ch11/solid-exercises/exercise-2-ocp/main.py` - Entry point demonstrating DiscountCalculator with string-based type checking
- `examples/ch11/solid-exercises/exercise-2-ocp/discount_calculator.py` - Bad implementation using if/elif chains for customer types
- `examples/ch11/solid-exercises/exercise-2-ocp/tests/test_discount_calculator.py` - Test suite validating discount calculations (works for both conditional and polymorphic implementations)
- `examples/ch11/solid-exercises/exercise-2-ocp/README.md` - Exercise documentation with setup instructions and learning objectives
- `examples/ch11/solid-exercises/exercise-2-ocp/.gitignore` - Python gitignore (venv, __pycache__, *.pyc)

### Exercise 3: DIP Project

- `examples/ch11/solid-exercises/exercise-3-dip/pyproject.toml` - UV project configuration with Python 3.11+, pytest, mypy, ruff dependencies
- `examples/ch11/solid-exercises/exercise-3-dip/main.py` - Entry point demonstrating NotificationService with tightly-coupled dependencies
- `examples/ch11/solid-exercises/exercise-3-dip/notification_service.py` - Bad implementation directly instantiating concrete email/SMS clients
- `examples/ch11/solid-exercises/exercise-3-dip/mock_clients.py` - Mock implementations of SMTPClient and TwilioClient for demonstration (print to console)
- `examples/ch11/solid-exercises/exercise-3-dip/tests/test_notification_service.py` - Test suite validating notification behavior (works for both tightly-coupled and dependency-injected implementations)
- `examples/ch11/solid-exercises/exercise-3-dip/README.md` - Exercise documentation with setup instructions and learning objectives
- `examples/ch11/solid-exercises/exercise-3-dip/.gitignore` - Python gitignore (venv, __pycache__, *.pyc)

### Documentation

- `docs/11-application-development/11.2.1-solid-principles.md` - SOLID principles documentation to be updated with project references

### Notes

- All projects use `uv` for dependency management and project configuration
- Tests should be run with `uv run pytest` from the project directory
- Type checking with `uv run mypy .` and linting with `uv run ruff check .`
- Test design pattern: Tests validate behavior/outcomes rather than internal structure, allowing the same tests to pass whether code uses "bad" patterns or refactored "good" patterns
- Mock services (email, SMS) print to stdout rather than requiring actual credentials
- Python 3.11+ required for modern type hints and features
- Follow existing ch11 example patterns for project structure (pyproject.toml, README.md, tests/)

## Tasks

### [x] 1.0 Create Exercise 1: Single Responsibility Principle (SRP) Project

#### 1.0 Proof Artifact(s)

- CLI: `cd examples/ch11/solid-exercises/exercise-1-srp && uv run main.py` successfully generates a report demonstrates the bad implementation works
- CLI: `cd examples/ch11/solid-exercises/exercise-1-srp && uv run pytest` shows all tests passing demonstrates test suite validates behavior
- File: `examples/ch11/solid-exercises/exercise-1-srp/README.md` exists with setup instructions demonstrates project documentation is complete
- CLI: `cd examples/ch11/solid-exercises/exercise-1-srp && uv run mypy .` shows type checking passes demonstrates code quality tooling works
- CLI: `cd examples/ch11/solid-exercises/exercise-1-srp && uv run ruff check .` shows linting passes demonstrates code quality standards

#### 1.0 Tasks

- [x] 1.1 Create directory structure `examples/ch11/solid-exercises/exercise-1-srp/` with tests subdirectory
- [x] 1.2 Create `pyproject.toml` with project metadata, Python >=3.11 requirement, and dependencies (pytest, mypy, ruff)
- [x] 1.3 Create `.gitignore` for Python projects (include .venv, __pycache__, *.pyc, *.html report files)
- [x] 1.4 Create `report_generator.py` with ReportGenerator class containing 3-4 responsibilities: data validation, statistics calculation, HTML formatting, and file writing
- [x] 1.5 Create `main.py` that imports ReportGenerator, creates sample data, and generates a report to demonstrate the working application
- [x] 1.6 Create `tests/test_report_generator.py` with behavior-based tests that validate report generation (tests data validation, statistics accuracy, file creation) without depending on internal class structure
- [x] 1.7 Create `README.md` with exercise overview, setup instructions (`uv sync`, `uv run main.py`, `uv run pytest`), learning objectives about SRP, and guidance on what concerns to separate
- [x] 1.8 Run `uv sync` to initialize the project and install dependencies
- [x] 1.9 Verify all tests pass with `uv run pytest`
- [x] 1.10 Verify type checking passes with `uv run mypy .`
- [x] 1.11 Verify linting passes with `uv run ruff check .`

### [x] 2.0 Create Exercise 2: Open/Closed Principle (OCP) Project

#### 2.0 Proof Artifact(s)

- CLI: `cd examples/ch11/solid-exercises/exercise-2-ocp && uv run main.py` calculates discounts for different customer types demonstrates the bad implementation works
- CLI: `cd examples/ch11/solid-exercises/exercise-2-ocp && uv run pytest` shows all tests passing demonstrates test suite validates behavior
- File: `examples/ch11/solid-exercises/exercise-2-ocp/README.md` exists with setup instructions demonstrates project documentation is complete
- CLI: `cd examples/ch11/solid-exercises/exercise-2-ocp && uv run mypy .` shows type checking passes demonstrates code quality tooling works
- CLI: `cd examples/ch11/solid-exercises/exercise-2-ocp && uv run ruff check .` shows linting passes demonstrates code quality standards

#### 2.0 Tasks

- [x] 2.1 Create directory structure `examples/ch11/solid-exercises/exercise-2-ocp/` with tests subdirectory
- [x] 2.2 Create `pyproject.toml` with project metadata, Python >=3.11 requirement, and dependencies (pytest, mypy, ruff)
- [x] 2.3 Create `.gitignore` for Python projects (include .venv, __pycache__, *.pyc)
- [x] 2.4 Create `discount_calculator.py` with DiscountCalculator class using if/elif chains for customer types (regular, premium, vip) with hard-coded discount percentages (0.05, 0.10, 0.15)
- [x] 2.5 Create `main.py` that imports DiscountCalculator and demonstrates calculating discounts for different customer types with sample amounts
- [x] 2.6 Create `tests/test_discount_calculator.py` with behavior-based tests that validate discount calculations for all customer types without depending on implementation method (conditionals vs polymorphism)
- [x] 2.7 Create `README.md` with exercise overview, setup instructions, learning objectives about OCP, and guidance on using abstract base classes and polymorphism
- [x] 2.8 Run `uv sync` to initialize the project and install dependencies
- [x] 2.9 Verify all tests pass with `uv run pytest`
- [x] 2.10 Verify type checking passes with `uv run mypy .`
- [x] 2.11 Verify linting passes with `uv run ruff check .`

### [x] 3.0 Create Exercise 3: Dependency Inversion Principle (DIP) Project

#### 3.0 Proof Artifact(s)

- CLI: `cd examples/ch11/solid-exercises/exercise-3-dip && uv run main.py` sends notifications through email and SMS channels demonstrates the bad implementation works
- CLI: `cd examples/ch11/solid-exercises/exercise-3-dip && uv run pytest` shows all tests passing demonstrates test suite validates behavior
- File: `examples/ch11/solid-exercises/exercise-3-dip/README.md` exists with setup instructions demonstrates project documentation is complete
- CLI: `cd examples/ch11/solid-exercises/exercise-3-dip && uv run mypy .` shows type checking passes demonstrates code quality tooling works
- CLI: `cd examples/ch11/solid-exercises/exercise-3-dip && uv run ruff check .` shows linting passes demonstrates code quality standards

#### 3.0 Tasks

- [x] 3.1 Create directory structure `examples/ch11/solid-exercises/exercise-3-dip/` with tests subdirectory
- [x] 3.2 Create `pyproject.toml` with project metadata, Python >=3.11 requirement, and dependencies (pytest, mypy, ruff)
- [x] 3.3 Create `.gitignore` for Python projects (include .venv, __pycache__, *.pyc)
- [x] 3.4 Create `mock_clients.py` with SMTPClient and TwilioClient mock classes that print to console instead of sending real messages (include methods: SMTPClient.send(email, message), TwilioClient.send(phone, message))
- [x] 3.5 Create `notification_service.py` with NotificationService class that directly instantiates SMTPClient and TwilioClient in __init__ and has notify(user, message) method that calls both clients
- [x] 3.6 Create `main.py` that imports NotificationService, creates a sample user object (with email and phone attributes), and demonstrates sending notifications through both channels
- [x] 3.7 Create `tests/test_notification_service.py` with behavior-based tests that validate notification behavior (tests both channels are called) without depending on whether dependencies are injected or instantiated internally
- [x] 3.8 Create `README.md` with exercise overview, setup instructions, learning objectives about DIP, and guidance on creating abstractions and using dependency injection
- [x] 3.9 Run `uv sync` to initialize the project and install dependencies
- [x] 3.10 Verify all tests pass with `uv run pytest`
- [x] 3.11 Verify type checking passes with `uv run mypy .`
- [x] 3.12 Verify linting passes with `uv run ruff check .`

### [ ] 4.0 Update Exercise Documentation to Reference New Projects

#### 4.0 Proof Artifact(s)

- File: `docs/11-application-development/11.2.1-solid-principles.md` Exercise 1 section references `examples/ch11/solid-exercises/exercise-1-srp/` demonstrates documentation integration for SRP
- File: `docs/11-application-development/11.2.1-solid-principles.md` Exercise 2 section references `examples/ch11/solid-exercises/exercise-2-ocp/` demonstrates documentation integration for OCP
- File: `docs/11-application-development/11.2.1-solid-principles.md` Exercise 3 section references `examples/ch11/solid-exercises/exercise-3-dip/` demonstrates documentation integration for DIP
- Diff: Changes to `11.2.1-solid-principles.md` show inline code replaced with project references demonstrates documentation migration complete

#### 4.0 Tasks

- [ ] 4.1 Update Exercise 1 section in `docs/11-application-development/11.2.1-solid-principles.md` to replace the inline `ReportGenerator` code block with a reference and link to the project directory `examples/ch11/solid-exercises/exercise-1-srp/`
- [ ] 4.2 Update Exercise 2 section to replace the inline `DiscountCalculator` code block with a reference and link to `examples/ch11/solid-exercises/exercise-2-ocp/`
- [ ] 4.3 Update Exercise 3 section to replace the inline `NotificationService` code block with a reference and link to `examples/ch11/solid-exercises/exercise-3-dip/`
- [ ] 4.4 Add instructions in each exercise section on how to access and run the projects (navigate to directory, run `uv sync`, `uv run main.py`, `uv run pytest`)
- [ ] 4.5 Verify all markdown links are correctly formatted and point to valid directories
- [ ] 4.6 Run `npm run lint` to ensure markdown formatting is correct
