# 01 Questions Round 1 - SOLID Principles Exercises

Please answer each question below (select one or more options, or add your own notes). Feel free to add additional context under any question.

## 1. Project Structure and Naming

What directory structure should be used for these exercise projects in examples/?

- [x] (A) `examples/ch11/solid-exercises/exercise-1-srp/`, `exercise-2-ocp/`, `exercise-3-dip/`
- [ ] (B) `examples/ch11/srp-refactor/`, `ocp-payment/`, `dip-notification/`
- [ ] (C) `examples/ch11/solid-srp/`, `solid-ocp/`, `solid-dip/`
- [ ] (D) Other (describe)

## 2. Test Strategy

How should tests be designed to pass for both "bad" and refactored code?

- [ ] (A) Separate test files: `test_bad.py` tests the bad implementation, `test_good.py` tests the refactored version
- [x] (B) Single test file with abstract tests that work for both implementations using different imports/configurations
- [ ] (C) Single test file that tests the external behavior/interface that should remain consistent
- [ ] (D) Other (describe)

## 3. Initial State

What should participants start with when they open the project?

- [x] (A) Only the "bad" code is present; participants create the refactored version from scratch
- [ ] (B) Both "bad" and "good" code present in separate modules; participants can compare
- [ ] (C) "Bad" code in main module, "good" code in a `solution/` directory for reference
- [ ] (D) Other (describe)

## 4. Running and Testing

How should participants interact with the projects?

- [x] (A) Simple `uv run main.py` to execute, `uv run pytest` to test
- [ ] (B) CLI commands like `uv run exercise --mode=bad` or `uv run exercise --mode=good`
- [ ] (C) Just tests - no runnable application, focus on test-driven refactoring
- [ ] (D) Other (describe)

## 5. Dependencies

What external dependencies should be included?

- [ ] (A) Minimal: just pytest for testing
- [x] (B) Include type checking (mypy) and linting (ruff) to demonstrate good practices
- [ ] (C) Add additional libraries that make examples more realistic (e.g., sqlalchemy for database, smtplib for email)
- [ ] (D) Other (describe)

## 6. Documentation in Projects

What documentation should exist in each project directory?

- [x] (A) Just a README.md with setup instructions and the exercise goal
- [ ] (B) README.md plus inline code comments explaining the anti-patterns
- [ ] (C) README.md, inline comments, and a SOLUTION.md explaining the refactoring approach
- [ ] (D) Other (describe)

## 7. Exercise 1 (SRP) - ReportGenerator Scope

The document shows validation, calculation, formatting, file saving, and email notification. Should all concerns be included?

- [ ] (A) Include all 5 concerns as shown in the document
- [x] (B) Simplify to 3-4 concerns to make the exercise more focused
- [ ] (C) Make it even more realistic by adding database operations
- [ ] (D) Other (describe)

## 8. Exercise 2 (OCP) - DiscountCalculator Implementation

The bad code uses string-based type checking. What should the refactored target be?

- [x] (A) Abstract base class with concrete implementations (as shown in payment example)
- [ ] (B) Protocol/interface-based approach using typing.Protocol
- [ ] (C) Strategy pattern with dictionary dispatch
- [ ] (D) Other (describe)

## 9. Exercise 3 (DIP) - NotificationService Complexity

Should the notification service support multiple channels (email + SMS)?

- [x] (A) Yes, both email and SMS as shown in the document example
- [ ] (B) Start with just email to keep it simpler
- [ ] (C) Add a third channel (push notifications) to make it more comprehensive
- [ ] (D) Other (describe)

## 10. Exercise Documentation Updates

How should the exercises in 11.2.1-solid-principles.md reference the new projects?

- [ ] (A) Add a prominent callout box at the start of each exercise section with a link
- [x] (B) Replace the inline code with a reference to the project directory
- [ ] (C) Keep existing inline code, add a footnote or "See also" section
- [ ] (D) Other (describe)

## 11. Python Version and UV Configuration

What Python version should be targeted?

- [ ] (A) Python 3.9+ (wider compatibility)
- [ ] (B) Python 3.10+ (good balance of features and compatibility)
- [x] (C) Python 3.11+ (modern features, pattern matching)
- [ ] (D) Python 3.12+ (latest stable)
- [ ] (E) Other (describe)

## 12. Proof Artifacts

What proof artifacts should demonstrate the projects work correctly?

- [ ] (A) Screenshots of test output showing all tests passing
- [x] (B) CLI output showing both bad and good implementations running
- [ ] (C) Both test output and CLI execution examples
- [ ] (D) Other (describe)
