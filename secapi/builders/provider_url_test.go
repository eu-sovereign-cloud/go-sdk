package builders

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProviderUrlBuilder(t *testing.T) {
	builder := NewProviderUrlBuilder()
	assert.NotNil(t, builder)
}

func TestProviderUrlBuilder_BuildPathUrl(t *testing.T) {
	builder := NewProviderUrlBuilder().
		BaseDomain("demo.secapi.cloud").
		PathProvider("seca.compute").
		Version("v1")

	pathUrl := builder.BuildPathUrl()

	assert.NotEmpty(t, pathUrl)
	assert.Equal(t, "https://demo.secapi.cloud/providers/seca.compute/v1/", pathUrl)
}

func TestProviderUrlBuilder_BuildDomainUrl(t *testing.T) {
	builder := NewProviderUrlBuilder().
		BaseDomain("demo.secapi.cloud").
		DomainProvider("compute.seca").
		Version("v1")

	domainUrl := builder.BuildDomainUrl()

	assert.NotEmpty(t, domainUrl)
	assert.Equal(t, "https://compute.seca.demo.secapi.cloud/v1/", domainUrl)
}
