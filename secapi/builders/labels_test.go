package builders

import (
	"testing"

	"github.com/eu-sovereign-cloud/go-sdk/internal/secatest"

	"github.com/stretchr/testify/assert"
)

func TestNewLabelsBuilder(t *testing.T) {
	builder := NewLabelsBuilder()
	assert.NotNil(t, builder)
	assert.Empty(t, builder.items)
}

func TestLabelsBuilder_Equals(t *testing.T) {
	builder := NewLabelsBuilder().Equals(secatest.LabelEnvKey, "prod")

	assert.Len(t, builder.items, 1)
	assert.Equal(t, secatest.LabelEnvKey+"=prod", builder.items[0])
	assert.Equal(t, secatest.LabelEnvKey+"=prod", builder.Build())
}

func TestLabelsBuilder_NsEquals(t *testing.T) {
	builder := NewLabelsBuilder().NsEquals(secatest.LabelMonitoringValue, secatest.LabelAlertLevelValue, secatest.LabelHightValue)

	assert.Len(t, builder.items, 1)
	assert.Equal(t, secatest.LabelMonitoringValue+":"+secatest.LabelAlertLevelValue+"="+secatest.LabelHightValue, builder.items[0])
	assert.Equal(t, secatest.LabelMonitoringValue+":"+secatest.LabelAlertLevelValue+"="+secatest.LabelHightValue, builder.Build())
}

func TestLabelsBuilder_Neq(t *testing.T) {
	builder := NewLabelsBuilder().Neq(secatest.LabelTierKey, "free")

	assert.Len(t, builder.items, 1)
	assert.Equal(t, secatest.LabelTierKey+"!=free", builder.items[0])
	assert.Equal(t, secatest.LabelTierKey+"!=free", builder.Build())
}

func TestLabelsBuilder_Gt(t *testing.T) {
	builder := NewLabelsBuilder().Gt(secatest.LabelVersion, 1)

	assert.Len(t, builder.items, 1)
	assert.Equal(t, secatest.LabelVersion+">1", builder.items[0])
	assert.Equal(t, secatest.LabelVersion+">1", builder.Build())
}

func TestLabelsBuilder_Lt(t *testing.T) {
	builder := NewLabelsBuilder().Lt(secatest.LabelVersion, 3)

	assert.Len(t, builder.items, 1)
	assert.Equal(t, secatest.LabelVersion+"<3", builder.items[0])
	assert.Equal(t, secatest.LabelVersion+"<3", builder.Build())
}

func TestLabelsBuilder_Gte(t *testing.T) {
	builder := NewLabelsBuilder().Gte(secatest.LabelUptime, 99)

	assert.Len(t, builder.items, 1)
	assert.Equal(t, secatest.LabelUptime+">=99", builder.items[0])
	assert.Equal(t, secatest.LabelUptime+">=99", builder.Build())
}

func TestLabelsBuilder_Lte(t *testing.T) {
	builder := NewLabelsBuilder().Lte(secatest.LabelLoad, 75)

	assert.Len(t, builder.items, 1)
	assert.Equal(t, secatest.LabelLoad+"<=75", builder.items[0])
	assert.Equal(t, secatest.LabelLoad+"<=75", builder.Build())
}

func TestLabelsBuilder_Build_Empty(t *testing.T) {
	builder := NewLabelsBuilder()

	assert.Equal(t, "", builder.Build())
}

func TestLabelsBuilder_BuildPtr_Empty(t *testing.T) {
	builder := NewLabelsBuilder()

	result := builder.BuildPtr()
	assert.Nil(t, result)
}

func TestLabelsBuilder_BuildPtr_NonEmpty(t *testing.T) {
	builder := NewLabelsBuilder().Equals(secatest.LabelEnvKey, "prod")

	result := builder.BuildPtr()
	assert.NotNil(t, result)
	assert.Equal(t, secatest.LabelEnvKey+"=prod", *result)
}

func TestLabelsBuilder_BuildPtr_MultipleFilters(t *testing.T) {
	builder := NewLabelsBuilder().
		Equals(secatest.LabelEnvKey, "prod").
		Gt(secatest.LabelVersion, 1)

	result := builder.BuildPtr()
	assert.NotNil(t, result)
	assert.Equal(t, secatest.LabelEnvKey+"=prod,"+secatest.LabelVersion+">1", *result)
}
