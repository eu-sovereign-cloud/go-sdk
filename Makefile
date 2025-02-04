GO := go
OPENAPI_GENERATOR := github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen
TOOLS_GOMOD := -modfile=./tools/go.mod
GO_TOOL := $(GO) run $(TOOLS_GOMOD)

PKG := pkg
SPEC_SRC := spec/spec
SPEC_DIST := spec/dist/specs
SPEC_SOURCES := $(shell ls $(SPEC_SRC)/*.yaml)
SCHEMAS_SOURCES := $(SPEC_SOURCES:$(SPEC_SRC)/%.yaml=$(SPEC_DIST)/%.yaml)
SCHEMAS_FINAL = $(SCHEMAS_SOURCES:$(SPEC_DIST)/%.yaml=$(PKG)/%/api.go)

all: $(SCHEMAS_SOURCES) $(SCHEMAS_FINAL)

$(SPEC_DIST)/%.yaml: $(SPEC_SRC)/%.yaml
	make -C spec dist/specs/$(shell basename $@)

$(PKG)/%/api.go: $(SPEC_DIST)/%.yaml
	@-mkdir -p $(shell dirname $@)
	-$(GO_TOOL) $(OPENAPI_GENERATOR) -alias-types -generate "types,client,std-http" \
		-package $(patsubst %.v1,%,$(shell basename $(shell dirname $@))) -o $@ $<

.PHONY: update	
update:
	git pull --recurse-submodules
	git submodule update --remote --recursive

.PHONY: test
test:
	$(GO) test -count=1 -cover -v ./...

.PHONY: mock
mock: all
	$(GO_TOOL) github.com/vektra/mockery/v2

.PHONY: fmt
fmt:
	$(GO_TOOL) mvdan.cc/gofumpt -w .

.PHONY: vet
vet:
	$(GO) vet ./...

.PHONY: clean
clean:
	rm -rf $(SCHEMAS_FINAL) $(SPEC_DIST) mock
	
