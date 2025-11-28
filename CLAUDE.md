# GO-SDK

## Overview

This is a Go SDK for the SECA (Sovereign European Cloud API) specification. The codebase is mostly handwritten but uses code generators for HTTP client code. The SDK will eventually use code generators for most code.

## Build and Development Commands

### Initial Setup
```bash
# Initialize submodule (contains the SECA API specification)
git submodule init

# Update all external dependencies
make update
```

### Code Generation
```bash
# Full regeneration workflow (recommended)
make clean spec generate mock

# Individual steps:
make spec      # Generate spec from YAML sources (runs in spec/ submodule)
make schemas   # Generate Go types from OpenAPI schemas
make mock      # Generate mocks using mockery
make generate  # Run all generation steps (clean + spec + schemas + mock)
```

The code generation workflow:
1. `make spec` - Builds OpenAPI specifications from YAML files in `spec/spec/` submodule
2. `make schemas` - Uses oapi-codegen to generate Go types from schemas into `pkg/spec/schema/`
3. Generated API clients go into `pkg/spec/<api-name>/api.go` (e.g., `pkg/spec/foundation.compute.v1/api.go`)
4. `make mock` - Uses mockery to generate mocks into `mock/` directory

### Testing
```bash
# Run all tests with coverage
make test

# This generates coverage.html which can be viewed in a browser
# Run a single test
go test -v ./secapi -run TestName

# Run tests in a specific package
go test -v ./secapi
```

### Code Quality
```bash
make fmt   # Format code using gofumpt
make lint  # Run golangci-lint with .golangci.yml config
make build # Build all packages
```

## Architecture

### Package Structure

**`pkg/spec/`** - Generated code from SECAPI specifications
- `schema/` - Generated Go types from SECAPI shared components (e.g., Instance, Network, etc.)
- `foundation.*/` - Generated HTTP clients for foundation APIs (compute, network, storage, etc.)
- `extensions.*/` - Generated HTTP clients for extension APIs (kubernetes, loadbalancer, etc.)

**`secapi/`** - SDK wrapper for the generated clients
- Provides higher-level, user-friendly interfaces
- Contains `GlobalClient` and `RegionalClient` for managing API access
- API-specific files: `compute_v1.go`, `network_v1.go`, `storage_v1.go`, etc.
- Common utilities: `iterator.go`, `references.go`, `list_options.go`

**`internal/secatest/`** - Testing utilities
- Mock HTTP servers for testing
- Test constants and helpers
- Mock implementations for each API version

**`mock/`** - Generated mocks (via mockery)
- Mocks for all generated client interfaces
- Used in unit tests

**`config/`** - oapi-codegen configuration files
- `api.yaml` - Config for generating API clients
- `schema.yaml` - Config for generating schema types

**`tools/`** - Go tool dependencies
- Declares tool dependencies in `tools.go` for version control

### Client Architecture

The SDK uses a two-tier client system:

**GlobalClient** (`secapi/global_client.go`)
- Entry point for the SDK
- Manages global/regional services
- Contains `RegionV1` and `AuthorizationV1` APIs
- Creates `RegionalClient` instances for specific regions

**RegionalClient** (`secapi/regional_client.go`)
- Region-specific client
- Contains `ComputeV1`, `NetworkV1`, `StorageV1`, `WorkspaceV1` APIs
- Created via `GlobalClient.NewRegionalClient(ctx, regionName)`
- Automatically discovers regional endpoints from Region API

Example flow:
1. Create `GlobalClient` with auth token and global endpoints
2. Use `RegionV1` API to discover available regions
3. Call `NewRegionalClient(ctx, "region-name")` to create regional client
4. Regional client auto-discovers provider endpoints for compute, network, storage, workspace

### Resource References

The SDK uses a reference system to identify resources hierarchically:

- `TenantReference` - For tenant-scoped resources (requires: Tenant, Name)
- `WorkspaceReference` - For workspace-scoped resources (requires: Tenant, Workspace, Name)
- `NetworkReference` - For network-scoped resources (requires: Tenant, Workspace, Network, Name)

These are defined in `secapi/references.go` and include validation methods.

### Iterators

Pagination is handled via the `Iterator[T]` generic type (`secapi/iterator.go`):
- `Next(ctx)` - Get next item, returns `io.EOF` when done
- `All(ctx)` - Fetch all items into a slice
- Handles skip tokens automatically

Most List operations return an `Iterator` to abstract pagination.

## API Versioning

APIs follow the pattern:
- Foundation APIs: `foundation.<service>.v<version>` (e.g., `foundation.compute.v1`)
- Extension APIs: `extensions.<service>.v<version>` (e.g., `extensions.kubernetes.v1beta1`)

## Testing Approach

Tests use:
1. Generated mocks from mockery (in `mock/` directory)
2. Custom test utilities in `internal/secatest/`
3. HTTP test servers that simulate API responses

Each API version has corresponding test file (e.g., `compute_v1.go` → `compute_v1_test.go`).

## Code Generation Configuration

**oapi-codegen** is configured via:
- `config/api.yaml` - Generates client interfaces with import mapping to `pkg/spec/schema`
- `config/schema.yaml` - Generates only types (models), no clients

**mockery** is configured via:
- `.mockery.yaml` - Defines which packages to mock and output locations

## Important Notes

- The `spec/` directory is a git submodule pointing to the SECA API specification repository
- Always run `make update` before regenerating code to get latest spec changes
- Generated files have consistent naming: `zz_generated.<name>.go` for mocks
- Tools are managed via `tools/go.mod` to ensure consistent versions across developers
