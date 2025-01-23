GO := go
OPENAPI_GENERATOR := github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen
TOOLS_GOMOD := -modfile=./tools/go.mod
GO_TOOL := $(GO) run $(TOOLS_GOMOD)

PKG := pkg
SPEC_DIST := spec/dist/specs/
SCHEMAS_SOURCES := $(shell ls $(SPEC_DIST)/*.yaml)
SCHEMAS_FINAL = $(SCHEMAS_SOURCES:$(SPEC_DIST)/%.yaml=$(PKG)/%/api.go)

all: $(SCHEMAS_FINAL)

$(PKG)/%/api.go: $(SPEC_DIST)/%.yaml
	@-mkdir -p $(shell dirname $@)
	-$(GO_TOOL) $(OPENAPI_GENERATOR) -alias-types -o $@ $<