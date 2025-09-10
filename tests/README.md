# Test Configuration
This directory contains configuration files and scripts for testing the gotodo application.

## Test Structure

### Unit Tests
- `internal/storage/storage_test.go` - Tests for storage layer operations
- `cmd/cmd_test.go` - Tests for command layer operations
- `cmd/test_helper.go` - Helper functions for command tests

### Test Coverage
The test suite covers:
- Storage operations (CRUD, error handling)
- Command execution (add, list, done, delete, clear)
- Configuration management
- File permissions and edge cases

### Running Tests

#### Run all tests
```bash
go test -v ./...
```

#### Run tests with coverage
```bash
go test -v -coverprofile=coverage.out ./...
go tool cover -html=coverage.out -o coverage.html
```

#### Run specific test package
```bash
go test -v ./internal/storage
go test -v ./cmd
```

#### Run specific test function
```bash
go test -v -run="TestStorageOperations" ./internal/storage
go test -v -run="TestAddCommand" ./cmd
```

#### Run tests with race detection
```bash
go test -race -v ./...
```

### Using Makefile

```bash
# Run all tests
make test

# Run unit tests only
make test-unit

# Run tests with coverage
make test-coverage

# Run benchmark tests
make test-benchmark

# Clean test artifacts
make clean
```

### Test Environment
- Tests use temporary directories that are automatically cleaned up
- Each test runs in isolation to prevent interference
- File operations are tested with various permission scenarios
- Error handling is tested for edge cases

### Continuous Integration
GitHub Actions workflow automatically runs:
- Unit tests on multiple Go versions (1.22, 1.23, 1.24)
- Integration tests
- Code coverage reporting
- Lint checks with golangci-lint