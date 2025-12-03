# Claude Code Working Notes

## Repository Purpose

Go client library for Transform IA API services. Provides type-safe Go
bindings for API endpoints with authentication and error handling.

## Plugin Usage

### When to use plugins

- `/go:cmd-build` - Build Go library
- `/go:cmd-test` - Run Go tests with coverage
- `/go:cmd-lint` - Lint Go code with golangci-lint
- `/go:cmd-tidy` - Update Go module dependencies
- `/github:cmd-status` - Check GitHub workflow status
- `/orchestrator:detect` - Auto-detect appropriate plugin

### Available plugins

- go, github, markdown, orchestrator

## Development Workflow

**Build Process:**

1. Modify Go code in `pkg/` or client packages
2. Run `/go:cmd-test ./...` to verify tests pass
3. Run `/go:cmd-lint .` to check code quality
4. Run `/go:cmd-tidy` to update dependencies
5. Commit changes

## Project Structure

- `pkg/client/` - Main client implementation
- `pkg/models/` - API request/response types
- `pkg/auth/` - Authentication handling
- `internal/` - Private HTTP client utilities
- `go.mod` - Go module definition
- `examples/` - Usage examples (if present)

## Client Library Features

- Type-safe API methods
- Automatic authentication (API key, OAuth, etc.)
- Request retries with exponential backoff
- Error handling and wrapping
- Context support for cancellation
- Rate limiting support

## Go Development Best Practices

- Write integration tests with API mocks
- Document all public client methods with godoc
- Use context.Context for timeouts
- Implement proper error types
- Version client library with semantic versioning

## Testing

- Unit tests: `/go:cmd-test ./...`
- Integration tests: Require API credentials or mocks
- Example code tests

## Usage Pattern

```go
import "github.com/transform-ia/client-api-go/pkg/client"

c := client.New("API_KEY")
result, err := c.GetResource(ctx, "resource-id")
```

## Deployment

Go module consumed via `go get github.com/transform-ia/client-api-go`.

## Integration

Used by:

- Go applications integrating with Transform IA services
- CLI tools for API access
- Backend services requiring API client
