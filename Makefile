GO := go
OPENAPI_GENERATOR := github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen
TOOLS_GOMOD := -modfile=./tools/go.mod
GO_TOOL := $(GO) run $(TOOLS_GOMOD)

PKG := pkg/spec
SPEC_SRC := spec/spec
SPEC_SCHEMAS_SRC := spec/spec/schemas
SPEC_SOURCES := $(shell ls $(SPEC_SRC)/*.yaml)
SPEC_SCHEMAS_SOURCES := $(shell ls $(SPEC_SCHEMAS_SRC)/*.yaml)

GEN_TARGETS := $(SPEC_SCHEMAS_SOURCES:$(SPEC_SCHEMAS_SRC)/%.yaml=$(PKG)/schema/%.go)
GEN_TARGETS += $(SPEC_SOURCES:$(SPEC_SRC)/%.yaml=$(PKG)/%/api.go)

.PHONY: update	
update:
	@echo "Updating submodules..."
	git pull --recurse-submodules
	git submodule update --remote --recursive

.PHONY: spec
spec:
	@echo "Generating spec..."
	sh -c "cd spec && make build"

.PHONY: schemas
schemas: $(GEN_TARGETS)
	@echo "Generating schemas..."

# generate types
$(PKG)/schema/%.go: $(SPEC_SCHEMAS_SRC)/%.yaml
	@-mkdir -p $(shell dirname $@)
	$(GO_TOOL) $(OPENAPI_GENERATOR) -config config/schema.yaml -o $@ $<

# generate api spec
$(PKG)/%/api.go: $(SPEC_SRC)/%.yaml
	@-mkdir -p $(shell dirname $@)
	$(GO_TOOL) $(OPENAPI_GENERATOR) -config config/api.yaml \
		-package $(shell basename $(shell dirname $@) | cut -d '.' -f 2) -o $@ $<

.PHONY: mock
mock: schemas
	@echo "Generating mocks..."
	$(GO_TOOL) github.com/vektra/mockery/v2

.PHONY: build
build:
	@echo "Building..."
	$(GO) build ./...

.PHONY: test
test:
	@echo "Running tests..."
	$(GO) test -count=1 -cover -coverprofile=coverage.out -v ./...
	$(GO) tool cover -html=coverage.out -o coverage.html
	rm coverage.out

.PHONY: fmt
fmt:
	@echo "Formating code..."
	$(GO_TOOL) mvdan.cc/gofumpt -w .

.PHONY: lint
lint:
	@echo "Linting code..."
	$(GO_TOOL) github.com/golangci/golangci-lint/cmd/golangci-lint run --verbose -c .golangci.yml

.PHONY: clean
clean:
	@echo "Cleaning..."
	rm -rf $(SCHEMAS_FINAL) $(SPEC_DIST) mock

.PHONY: generate
generate: clean spec schemas mock

.PHONY: tag
tag:
ifndef VERSION
    $(error VERSION is required. Usage: make tag VERSION=v0.3.20)
endif
	@echo "Tagging $(VERSION)..."
	git tag $(VERSION)
	git push origin $(VERSION)