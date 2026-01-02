# 01-spec-solid-exercises.md

## Introduction/Overview

This specification defines the creation of three standalone Python projects for the SOLID Principles exercises in the DevOps Bootcamp. Each project will provide participants with a working "bad" code implementation that violates a specific SOLID principle, along with a test suite that validates behavior. Participants will refactor the code to follow SOLID principles while ensuring all tests continue to pass, demonstrating that good design maintains functionality while improving maintainability.

## Goals

1. Create three runnable Python projects (SRP, OCP, DIP) with uv configuration that participants can execute and test locally
2. Provide test suites that validate behavior for both the initial "bad" implementation and participant-refactored "good" implementations
3. Update exercise documentation to reference the new project structure, replacing inline code examples with project references
4. Demonstrate modern Python development practices including type checking (mypy) and linting (ruff)
5. Ensure projects are self-contained, well-documented, and align with the existing bootcamp structure

## User Stories

1. **As a bootcamp participant**, I want to work with runnable code examples so that I can see SOLID violations in action and test my refactored solutions against the same test suite.

2. **As an instructor**, I want self-contained exercise projects so that I can easily distribute exercises and verify that participant solutions maintain correct behavior.

3. **As a developer learning SOLID principles**, I want tests that pass for both bad and good implementations so that I can verify my refactoring preserves functionality while improving design.

## Demoable Units of Work

### Unit 1: Exercise 1 - Single Responsibility Principle (SRP)

**Purpose:** Create a working ReportGenerator project that demonstrates SRP violations through a class handling multiple concerns (validation, calculation, formatting, file I/O), with participants refactoring into separate single-responsibility classes.

**Functional Requirements:**
- The system shall provide a `ReportGenerator` class that combines 3-4 distinct responsibilities (validation, calculation, formatting, and file operations)
- The system shall include a `main.py` that demonstrates the report generation functionality when executed
- The system shall provide a test suite with abstract tests that validate behavior regardless of implementation structure
- The user shall be able to run `uv run main.py` to execute the bad implementation
- The user shall be able to run `uv run pytest` to verify tests pass with the bad implementation
- The system shall include a `pyproject.toml` configured for Python 3.11+ with dependencies: pytest, mypy, ruff
- The system shall include a README.md explaining the exercise goal, setup instructions, and expected refactoring outcome

**Proof Artifacts:**
- CLI output: Running `uv run main.py` in `examples/ch11/solid-exercises/exercise-1-srp/` successfully generates a report, demonstrating the bad implementation works
- CLI output: Running `uv run pytest` shows all tests passing with the bad implementation, demonstrating test validity

### Unit 2: Exercise 2 - Open/Closed Principle (OCP)

**Purpose:** Create a DiscountCalculator project that demonstrates OCP violations through string-based type checking and if/elif chains, with participants refactoring to use abstract base classes and polymorphism.

**Functional Requirements:**
- The system shall provide a `DiscountCalculator` class that uses string-based customer type checking with if/elif chains
- The system shall include a `main.py` that demonstrates discount calculation for multiple customer types
- The system shall provide a test suite that validates discount calculations work correctly regardless of whether implementation uses conditionals or polymorphism
- The user shall be able to run `uv run main.py` to execute the bad implementation
- The user shall be able to run `uv run pytest` to verify tests pass
- The system shall include a `pyproject.toml` configured for Python 3.11+ with dependencies: pytest, mypy, ruff
- The system shall include a README.md explaining the exercise goal and OCP principle

**Proof Artifacts:**
- CLI output: Running `uv run main.py` in `examples/ch11/solid-exercises/exercise-2-ocp/` calculates discounts correctly for different customer types, demonstrating the bad implementation works
- CLI output: Running `uv run pytest` shows all tests passing, demonstrating behavior validation

### Unit 3: Exercise 3 - Dependency Inversion Principle (DIP)

**Purpose:** Create a NotificationService project that demonstrates DIP violations through direct instantiation of concrete dependencies (SMTP email client and SMS client), with participants refactoring to use dependency injection and abstractions.

**Functional Requirements:**
- The system shall provide a `NotificationService` class that directly instantiates concrete `SMTPClient` and `TwilioClient` classes in its constructor
- The system shall include mock implementations of email and SMS clients that print to console instead of sending real messages
- The system shall include a `main.py` that demonstrates sending notifications via both channels
- The system shall provide a test suite that validates notification behavior works with both tightly-coupled and dependency-injected implementations
- The user shall be able to run `uv run main.py` to execute the bad implementation
- The user shall be able to run `uv run pytest` to verify tests pass
- The system shall include a `pyproject.toml` configured for Python 3.11+ with dependencies: pytest, mypy, ruff
- The system shall include a README.md explaining the exercise goal and DIP principle

**Proof Artifacts:**
- CLI output: Running `uv run main.py` in `examples/ch11/solid-exercises/exercise-3-dip/` sends notifications through both email and SMS channels, demonstrating the bad implementation works
- CLI output: Running `uv run pytest` shows all tests passing, demonstrating test coverage

### Unit 4: Documentation Updates

**Purpose:** Update the SOLID principles documentation to reference the new standalone projects, replacing inline code examples with project references.

**Functional Requirements:**
- The system shall update `docs/11-application-development/11.2.1-solid-principles.md` Exercise 1 section to replace inline code with a reference to `examples/ch11/solid-exercises/exercise-1-srp/`
- The system shall update Exercise 2 section to replace inline code with a reference to `examples/ch11/solid-exercises/exercise-2-ocp/`
- The system shall update Exercise 3 section to replace inline code with a reference to `examples/ch11/solid-exercises/exercise-3-dip/`
- The documentation shall maintain the educational explanation of SOLID principles while directing participants to the runnable projects
- The documentation shall include instructions on how to navigate to and run the projects

**Proof Artifacts:**
- File content: Updated `docs/11-application-development/11.2.1-solid-principles.md` shows project references instead of inline code in exercise sections, demonstrating documentation integration

## Non-Goals (Out of Scope)

1. **Solution implementations**: Projects will NOT include "good" refactored code or solution directories - participants must create refactored versions themselves
2. **Automated solution validation**: Will NOT include automated checking to verify participant refactorings follow SOLID principles correctly (beyond tests passing)
3. **Interactive tutorials or guided refactoring**: Will NOT include step-by-step refactoring guides or interactive prompts
4. **Real external service integration**: Will NOT use actual SMTP servers, SMS APIs, or databases - all external services will be mocked/simulated
5. **Multiple language implementations**: Will NOT create versions in other programming languages beyond Python
6. **CI/CD integration**: Will NOT set up GitHub Actions or other CI/CD pipelines for the exercise projects
7. **Additional SOLID exercises**: Will NOT create projects for Liskov Substitution Principle (LSP) or Interface Segregation Principle (ISP) beyond the three specified

## Design Considerations

No specific design requirements identified. These are command-line Python projects focused on code structure rather than UI/UX.

## Repository Standards

The implementation should follow these established bootcamp patterns:

- **Directory Organization**: Place projects under `examples/ch11/` following the chapter-based structure
- **Front-matter Metadata**: Existing exercises already have metadata in `docs/11-application-development/11.2.1-solid-principles.md` that should be preserved
- **Documentation Style**: Follow the bootcamp's markdown style guide in `STYLE.md`
- **Project Structure**: Use uv for Python dependency management as this is the chosen tool for the bootcamp
- **Code Quality**: README files should follow the bootcamp's documentation patterns with setup instructions and clear learning objectives

## Technical Considerations

- **Python Version**: Target Python 3.11+ to leverage modern features including pattern matching and improved type hints
- **UV Configuration**: Each `pyproject.toml` must include:
  - Project name, version, and description
  - Python version requirement (>=3.11)
  - Dependencies: pytest, mypy, ruff
  - Build system configuration for uv
- **Test Design Pattern**: Tests must be implementation-agnostic, testing behavior/outcomes rather than internal structure, allowing the same tests to pass whether code uses the "bad" pattern or refactored "good" pattern
- **Import Strategy**: Test files should use configurable imports or module paths to test different implementations without modifying test code
- **Mock Objects**: For Exercise 3, email and SMS clients should be simple mock implementations that print to stdout rather than requiring actual service credentials

## Security Considerations

- **No Real Credentials**: Projects will NOT include or require actual API keys, SMTP credentials, or SMS service tokens
- **Mock Services Only**: All external service interactions (email, SMS, database) will be simulated/mocked
- **No Sensitive Data**: Test data and examples will use fictional names, emails, and phone numbers
- **Proof Artifacts**: CLI output demonstrations will NOT contain any real credentials or sensitive information

## Success Metrics

1. **Functional Completeness**: All three exercise projects successfully run with `uv run main.py` and pass tests with `uv run pytest`
2. **Test Resilience**: The same test suite passes for both the initial "bad" implementation and a correctly refactored "good" implementation
3. **Documentation Clarity**: Exercise documentation successfully references projects with clear instructions on how to access and run them
4. **Participant Usability**: Projects include sufficient README documentation for participants to understand goals and run exercises independently

## Open Questions

No open questions at this time.
