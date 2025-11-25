package builders

import (
	"github.com/eu-sovereign-cloud/go-sdk/pkg/spec/schema"
)

// Network

type NetworkBuilder struct {
	*resourceBuilder[NetworkBuilder, schema.NetworkSpec]
	metadata *RegionalWorkspaceResourceMetadataBuilder
	spec     *schema.NetworkSpec
}

func NewNetworkBuilder() *NetworkBuilder {
	builder := &NetworkBuilder{
		metadata: NewRegionalWorkspaceResourceMetadataBuilder(),
		spec:     &schema.NetworkSpec{},
	}

	builder.resourceBuilder = newResourceBuilder(newResourceBuilderParams[NetworkBuilder, schema.NetworkSpec]{
		parent:        builder,
		setName:       func(name string) { builder.metadata.setName(name) },
		setProvider:   func(provider string) { builder.metadata.setProvider(provider) },
		setResource:   func(resource string) { builder.metadata.setResource(resource) },
		setApiVersion: func(apiVersion string) { builder.metadata.setApiVersion(apiVersion) },
		setSpec:       func(spec *schema.NetworkSpec) { builder.spec = spec },
	})

	return builder
}

func (builder *NetworkBuilder) Tenant(tenant string) *NetworkBuilder {
	builder.metadata.Tenant(tenant)
	return builder
}

func (builder *NetworkBuilder) Workspace(workspace string) *NetworkBuilder {
	builder.metadata.Workspace(workspace)
	return builder
}

func (builder *NetworkBuilder) Region(region string) *NetworkBuilder {
	builder.metadata.Region(region)
	return builder
}

func (builder *NetworkBuilder) BuildResponse() (*schema.Network, error) {
	metadata, err := builder.metadata.Kind(schema.RegionalWorkspaceResourceMetadataKindResourceKindNetwork).BuildResponse()
	if err != nil {
		return nil, err
	}

	// Validate the spec
	if err := validateRequired(builder.validator,
		builder.spec,
		builder.spec.Cidr,
		builder.spec.SkuRef,
		builder.spec.RouteTableRef,
	); err != nil {
		return nil, err
	}
	if err := validateOneRequired(builder.validator,
		builder.spec.Cidr.Ipv4,
		builder.spec.Cidr.Ipv6,
	); err != nil {
		return nil, err
	}

	return &schema.Network{
		Metadata: metadata,
		Spec:     *builder.spec,
		Status:   &schema.NetworkStatus{},
	}, nil
}

// Internet Gateway

type InternetGatewayBuilder struct {
	*resourceBuilder[InternetGatewayBuilder, schema.InternetGatewaySpec]
	metadata *RegionalWorkspaceResourceMetadataBuilder
	spec     *schema.InternetGatewaySpec
}

func NewInternetGatewayBuilder() *InternetGatewayBuilder {
	builder := &InternetGatewayBuilder{
		metadata: NewRegionalWorkspaceResourceMetadataBuilder(),
		spec:     &schema.InternetGatewaySpec{},
	}

	builder.resourceBuilder = newResourceBuilder(newResourceBuilderParams[InternetGatewayBuilder, schema.InternetGatewaySpec]{
		parent:        builder,
		setName:       func(name string) { builder.metadata.setName(name) },
		setProvider:   func(provider string) { builder.metadata.setProvider(provider) },
		setResource:   func(resource string) { builder.metadata.setResource(resource) },
		setApiVersion: func(apiVersion string) { builder.metadata.setApiVersion(apiVersion) },
		setSpec:       func(spec *schema.InternetGatewaySpec) { builder.spec = spec },
	})

	return builder
}

func (builder *InternetGatewayBuilder) Tenant(tenant string) *InternetGatewayBuilder {
	builder.metadata.Tenant(tenant)
	return builder
}

func (builder *InternetGatewayBuilder) Workspace(workspace string) *InternetGatewayBuilder {
	builder.metadata.Workspace(workspace)
	return builder
}

func (builder *InternetGatewayBuilder) Region(region string) *InternetGatewayBuilder {
	builder.metadata.Region(region)
	return builder
}

func (builder *InternetGatewayBuilder) BuildResponse() (*schema.InternetGateway, error) {
	metadata, err := builder.metadata.Kind(schema.RegionalWorkspaceResourceMetadataKindResourceKindInternetGateway).BuildResponse()
	if err != nil {
		return nil, err
	}

	// Validate the spec
	if err := validateRequired(builder.validator,
		builder.spec,
		builder.spec.EgressOnly,
	); err != nil {
		return nil, err
	}

	return &schema.InternetGateway{
		Metadata: metadata,
		Spec:     *builder.spec,
		Status:   &schema.InternetGatewayStatus{},
	}, nil
}

// Route Table

type RouteTableBuilder struct {
	*resourceBuilder[RouteTableBuilder, schema.RouteTableSpec]
	metadata *RegionalNetworkResourceMetadataBuilder
	spec     *schema.RouteTableSpec
}

func NewRouteTableBuilder() *RouteTableBuilder {
	builder := &RouteTableBuilder{
		metadata: NewRegionalNetworkResourceMetadataBuilder(),
		spec:     &schema.RouteTableSpec{},
	}

	builder.resourceBuilder = newResourceBuilder(newResourceBuilderParams[RouteTableBuilder, schema.RouteTableSpec]{
		parent:        builder,
		setName:       func(name string) { builder.metadata.setName(name) },
		setProvider:   func(provider string) { builder.metadata.setProvider(provider) },
		setResource:   func(resource string) { builder.metadata.setResource(resource) },
		setApiVersion: func(apiVersion string) { builder.metadata.setApiVersion(apiVersion) },
		setSpec:       func(spec *schema.RouteTableSpec) { builder.spec = spec },
	})

	return builder
}

func (builder *RouteTableBuilder) Tenant(tenant string) *RouteTableBuilder {
	builder.metadata.Tenant(tenant)
	return builder
}

func (builder *RouteTableBuilder) Workspace(workspace string) *RouteTableBuilder {
	builder.metadata.Workspace(workspace)
	return builder
}

func (builder *RouteTableBuilder) Network(network string) *RouteTableBuilder {
	builder.metadata.Network(network)
	return builder
}

func (builder *RouteTableBuilder) Region(region string) *RouteTableBuilder {
	builder.metadata.Region(region)
	return builder
}

func (builder *RouteTableBuilder) BuildResponse() (*schema.RouteTable, error) {
	metadata, err := builder.metadata.Kind(schema.RegionalNetworkResourceMetadataKindResourceKindRoutingTable).BuildResponse()
	if err != nil {
		return nil, err
	}

	// Validate the spec
	if err := validateRequired(builder.validator,
		builder.spec.Routes,
	); err != nil {
		return nil, err
	}
	// Validate each route
	for _, route := range builder.spec.Routes {
		if err := validateRequired(builder.validator,
			route.DestinationCidrBlock,
			route.TargetRef,
		); err != nil {
			return nil, err
		}
	}

	return &schema.RouteTable{
		Metadata: metadata,
		Spec:     *builder.spec,
		Status:   &schema.RouteTableStatus{},
	}, nil
}

// Subnet

type SubnetBuilder struct {
	*resourceBuilder[SubnetBuilder, schema.SubnetSpec]
	metadata *RegionalNetworkResourceMetadataBuilder
	spec     *schema.SubnetSpec
}

func NewSubnetBuilder() *SubnetBuilder {
	builder := &SubnetBuilder{
		metadata: NewRegionalNetworkResourceMetadataBuilder(),
		spec:     &schema.SubnetSpec{},
	}

	builder.resourceBuilder = newResourceBuilder(newResourceBuilderParams[SubnetBuilder, schema.SubnetSpec]{
		parent:        builder,
		setName:       func(name string) { builder.metadata.setName(name) },
		setProvider:   func(provider string) { builder.metadata.setProvider(provider) },
		setResource:   func(resource string) { builder.metadata.setResource(resource) },
		setApiVersion: func(apiVersion string) { builder.metadata.setApiVersion(apiVersion) },
		setSpec:       func(spec *schema.SubnetSpec) { builder.spec = spec },
	})

	return builder
}

func (builder *SubnetBuilder) Tenant(tenant string) *SubnetBuilder {
	builder.metadata.Tenant(tenant)
	return builder
}

func (builder *SubnetBuilder) Workspace(workspace string) *SubnetBuilder {
	builder.metadata.Workspace(workspace)
	return builder
}

func (builder *SubnetBuilder) Network(network string) *SubnetBuilder {
	builder.metadata.Network(network)
	return builder
}

func (builder *SubnetBuilder) Region(region string) *SubnetBuilder {
	builder.metadata.Region(region)
	return builder
}

func (builder *SubnetBuilder) BuildResponse() (*schema.Subnet, error) {
	metadata, err := builder.metadata.Kind(schema.RegionalNetworkResourceMetadataKindResourceKindSubnet).BuildResponse()
	if err != nil {
		return nil, err
	}

	// Validate the spec
	if err := validateRequired(builder.validator,
		builder.spec,
		builder.spec.Cidr,
		builder.spec.Zone,
	); err != nil {
		return nil, err
	}

	return &schema.Subnet{
		Metadata: metadata,
		Spec:     *builder.spec,
		Status:   &schema.SubnetStatus{},
	}, nil
}

// Public Ip

type PublicIpBuilder struct {
	*resourceBuilder[PublicIpBuilder, schema.PublicIpSpec]
	metadata *RegionalWorkspaceResourceMetadataBuilder
	spec     *schema.PublicIpSpec
}

func NewPublicIpBuilder() *PublicIpBuilder {
	builder := &PublicIpBuilder{
		metadata: NewRegionalWorkspaceResourceMetadataBuilder(),
		spec:     &schema.PublicIpSpec{},
	}

	builder.resourceBuilder = newResourceBuilder(newResourceBuilderParams[PublicIpBuilder, schema.PublicIpSpec]{
		parent:        builder,
		setName:       func(name string) { builder.metadata.setName(name) },
		setProvider:   func(provider string) { builder.metadata.setProvider(provider) },
		setResource:   func(resource string) { builder.metadata.setResource(resource) },
		setApiVersion: func(apiVersion string) { builder.metadata.setApiVersion(apiVersion) },
		setSpec:       func(spec *schema.PublicIpSpec) { builder.spec = spec },
	})

	return builder
}

func (builder *PublicIpBuilder) Tenant(tenant string) *PublicIpBuilder {
	builder.metadata.Tenant(tenant)
	return builder
}

func (builder *PublicIpBuilder) Workspace(workspace string) *PublicIpBuilder {
	builder.metadata.Workspace(workspace)
	return builder
}

func (builder *PublicIpBuilder) Region(region string) *PublicIpBuilder {
	builder.metadata.Region(region)
	return builder
}

func (builder *PublicIpBuilder) BuildResponse() (*schema.PublicIp, error) {
	metadata, err := builder.metadata.Kind(schema.RegionalWorkspaceResourceMetadataKindResourceKindPublicIP).BuildResponse()
	if err != nil {
		return nil, err
	}

	// Validate the spec
	if err := validateRequired(builder.validator,
		builder.spec,
		builder.spec.Version,
	); err != nil {
		return nil, err
	}

	return &schema.PublicIp{
		Metadata: metadata,
		Spec:     *builder.spec,
		Status:   &schema.PublicIpStatus{},
	}, nil
}

// Nic

type NicBuilder struct {
	*resourceBuilder[NicBuilder, schema.NicSpec]
	metadata *RegionalWorkspaceResourceMetadataBuilder
	spec     *schema.NicSpec
}

func NewNicBuilder() *NicBuilder {
	builder := &NicBuilder{
		metadata: NewRegionalWorkspaceResourceMetadataBuilder(),
		spec:     &schema.NicSpec{},
	}

	builder.resourceBuilder = newResourceBuilder(newResourceBuilderParams[NicBuilder, schema.NicSpec]{
		parent:        builder,
		setName:       func(name string) { builder.metadata.setName(name) },
		setProvider:   func(provider string) { builder.metadata.setProvider(provider) },
		setResource:   func(resource string) { builder.metadata.setResource(resource) },
		setApiVersion: func(apiVersion string) { builder.metadata.setApiVersion(apiVersion) },
		setSpec:       func(spec *schema.NicSpec) { builder.spec = spec },
	})

	return builder
}

func (builder *NicBuilder) Tenant(tenant string) *NicBuilder {
	builder.metadata.Tenant(tenant)
	return builder
}

func (builder *NicBuilder) Workspace(workspace string) *NicBuilder {
	builder.metadata.Workspace(workspace)
	return builder
}

func (builder *NicBuilder) Region(region string) *NicBuilder {
	builder.metadata.Region(region)
	return builder
}

func (builder *NicBuilder) BuildResponse() (*schema.Nic, error) {
	metadata, err := builder.metadata.Kind(schema.RegionalWorkspaceResourceMetadataKindResourceKindNic).BuildResponse()
	if err != nil {
		return nil, err
	}

	// Validate the spec
	if err := validateRequired(builder.validator,
		builder.spec,
		builder.spec.Addresses,
		builder.spec.SubnetRef,
	); err != nil {
		return nil, err
	}

	return &schema.Nic{
		Metadata: metadata,
		Spec:     *builder.spec,
		Status:   &schema.NicStatus{},
	}, nil
}

// Security Group

type SecurityGroupBuilder struct {
	*resourceBuilder[SecurityGroupBuilder, schema.SecurityGroupSpec]
	metadata *RegionalWorkspaceResourceMetadataBuilder
	spec     *schema.SecurityGroupSpec
}

func NewSecurityGroupBuilder() *SecurityGroupBuilder {
	builder := &SecurityGroupBuilder{
		metadata: NewRegionalWorkspaceResourceMetadataBuilder(),
		spec:     &schema.SecurityGroupSpec{},
	}

	builder.resourceBuilder = newResourceBuilder(newResourceBuilderParams[SecurityGroupBuilder, schema.SecurityGroupSpec]{
		parent:        builder,
		setName:       func(name string) { builder.metadata.setName(name) },
		setProvider:   func(provider string) { builder.metadata.setProvider(provider) },
		setResource:   func(resource string) { builder.metadata.setResource(resource) },
		setApiVersion: func(apiVersion string) { builder.metadata.setApiVersion(apiVersion) },
		setSpec:       func(spec *schema.SecurityGroupSpec) { builder.spec = spec },
	})

	return builder
}

func (builder *SecurityGroupBuilder) Tenant(tenant string) *SecurityGroupBuilder {
	builder.metadata.Tenant(tenant)
	return builder
}

func (builder *SecurityGroupBuilder) Workspace(workspace string) *SecurityGroupBuilder {
	builder.metadata.Workspace(workspace)
	return builder
}

func (builder *SecurityGroupBuilder) Region(region string) *SecurityGroupBuilder {
	builder.metadata.Region(region)
	return builder
}

func (builder *SecurityGroupBuilder) BuildResponse() (*schema.SecurityGroup, error) {
	metadata, err := builder.metadata.Kind(schema.RegionalWorkspaceResourceMetadataKindResourceKindSecurityGroup).BuildResponse()
	if err != nil {
		return nil, err
	}

	// Validate the spec
	if err := validateRequired(builder.validator,
		builder.spec,
		builder.spec.Rules,
	); err != nil {
		return nil, err
	}
	// Validate each rule
	for _, rule := range builder.spec.Rules {
		if err := validateRequired(builder.validator,
			rule.Direction,
		); err != nil {
			return nil, err
		}
	}

	return &schema.SecurityGroup{
		Metadata: metadata,
		Spec:     *builder.spec,
		Status:   &schema.SecurityGroupStatus{},
	}, nil
}
