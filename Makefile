GO := go
OPENAPI_GENERATOR := github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen
TOOLS_GOMOD := -modfile=./tools/go.mod
GO_TOOL := $(GO) run $(TOOLS_GOMOD)

PKG := pkg/spec
SPEC_SRC := spec/spec
SPEC_DIST := spec/dist/specs
SPEC_SOURCES := $(shell ls $(SPEC_SRC)/*.yaml)
SCHEMAS_SOURCES := $(SPEC_SOURCES:$(SPEC_SRC)/%.yaml=$(SPEC_DIST)/%.yaml)
SCHEMAS_FINAL = $(SCHEMAS_SOURCES:$(SPEC_DIST)/%.yaml=$(PKG)/%/api.go)

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
schemas: $(SCHEMAS_SOURCES) $(SCHEMAS_FINAL)
	@echo "Generated schemas..."

$(SPEC_DIST)/%.yaml: $(SPEC_SRC)/%.yaml
	make -C spec dist/specs/$(shell basename $@)

$(PKG)/%/api.go: $(SPEC_DIST)/%.yaml
	@-mkdir -p $(shell dirname $@)
	-$(GO_TOOL) $(OPENAPI_GENERATOR) -alias-types -generate "types,client,std-http" \
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
	@echo "Running gofumpt..."
	$(GO_TOOL) mvdan.cc/gofumpt -w .

.PHONY: golint
golint:
	@echo "Running golint..."
	$(GO_TOOL) github.com/golangci/golangci-lint/cmd/golangci-lint run --timeout 5m -c .golangci.yml

.PHONY: vet
vet:
	@echo "Running vet..."
	$(GO) vet ./...

.PHONY: sec
sec:
	@echo "Running gosec..."
	$(GO_TOOL) github.com/securego/gosec/v2/cmd/gosec ./...

.PHONY: lint
lint: fmt golint vet sec

.PHONY: clean
clean:
	@echo "Cleaning..."
	rm -rf $(SCHEMAS_FINAL) $(SPEC_DIST) mock

.PHONY: generate
generate: clean spec schemas mock