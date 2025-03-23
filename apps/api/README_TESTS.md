# Vinance Backend Tests

This document provides instructions on how to run the unit tests for the Vinance backend application.

## Test Coverage

The application has unit tests for the following components:

### Handlers
- Login handler
- Register handler
- Create Account handler

### Services
- User service
- Account service

### Repositories
- User repository
- Account repository

### Utils
- Password hashing utilities

### Middleware
- Authentication middleware

### Authenticator
- JWT token generation and validation

## Running Tests

You can run all tests using the go test command:

```bash
go test ./...
```

To run tests with coverage reporting:

```bash
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out -o coverage.html
```

To run tests for a specific package:

```bash
go test ./pkg/authenticator
go test ./pkg/middleware
go test ./pkg/utils
go test ./repository/db
go test ./pkg/service
go test ./handler/server
```

## Test Structure

Each test file follows a standard structure:

1. **Setup** - Preparing test data and mocks
2. **Execute** - Calling the function or method being tested
3. **Assert** - Verifying the results or behavior

Many tests use the `testify/assert` package for cleaner assertions and the `testify/mock` package for mocking dependencies.

## Mocking Strategy

For tests that require external dependencies:

1. Mock interfaces are created to simulate the behavior of real components
2. Expectations are set on mock objects to verify interactions
3. The dependency injection pattern enables easy swapping of real implementations with mocks

## Best Practices

When adding new features:

1. Write tests before implementing the feature (Test-Driven Development)
2. Aim for high test coverage, especially for critical paths
3. Consider edge cases and error scenarios
4. Keep tests independent and avoid dependencies between test cases
5. Use descriptive test names that explain what's being tested 