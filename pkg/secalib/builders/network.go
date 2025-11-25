package builders

import (
	"github.com/eu-sovereign-cloud/go-sdk/pkg/spec/schema"
)

// Network

type NetworkBuilder struct {
	*regionalWorkspaceResourceBuilder[NetworkBuilder, schema.NetworkSpec]
	metadata *RegionalWorkspaceResourceMetadataBuilder
	spec     *schema.NetworkSpec
}

func NewNetworkBuilder() *NetworkBuilder {
	builder := &NetworkBuilder{
		metadata: NewRegionalWorkspaceResourceMetadataBuilder(),
		spec:     &schema.NetworkSpec{},
	}

	builder.regionalWorkspaceResourceBuilder = newRegionalWorkspaceResourceBuilder(newRegionalWorkspaceResourceBuilderParams[NetworkBuilder, schema.NetworkSpec]{
		newGlobalResourceBuilderParams: &newGlobalResourceBuilderParams[NetworkBuilder, schema.NetworkSpec]{
			parent:        builder,
			setName:       func(name string) { builder.metadata.setName(name) },
			setProvider:   func(provider string) { builder.metadata.setProvider(provider) },
			setResource:   func(resource string) { builder.metadata.setResource(resource) },
			setApiVersion: func(apiVersion string) { builder.metadata.setApiVersion(apiVersion) },
			setSpec:       func(spec *schema.NetworkSpec) { builder.spec = spec },
		},
		setTenant:    func(tenant string) { builder.metadata.Tenant(tenant) },
		setWorkspace: func(workspace string) { builder.metadata.Workspace(workspace) },
		setRegion:    func(region string) { builder.metadata.Region(region) },
	})

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
	*regionalWorkspaceResourceBuilder[InternetGatewayBuilder, schema.InternetGatewaySpec]
	metadata *RegionalWorkspaceResourceMetadataBuilder
	spec     *schema.InternetGatewaySpec
}

func NewInternetGatewayBuilder() *InternetGatewayBuilder {
	builder := &InternetGatewayBuilder{
		metadata: NewRegionalWorkspaceResourceMetadataBuilder(),
		spec:     &schema.InternetGatewaySpec{},
	}

	builder.regionalWorkspaceResourceBuilder = newRegionalWorkspaceResourceBuilder(newRegionalWorkspaceResourceBuilderParams[InternetGatewayBuilder, schema.InternetGatewaySpec]{
		newGlobalResourceBuilderParams: &newGlobalResourceBuilderParams[InternetGatewayBuilder, schema.InternetGatewaySpec]{
			parent:        builder,
			setName:       func(name string) { builder.metadata.setName(name) },
			setProvider:   func(provider string) { builder.metadata.setProvider(provider) },
			setResource:   func(resource string) { builder.metadata.setResource(resource) },
			setApiVersion: func(apiVersion string) { builder.metadata.setApiVersion(apiVersion) },
			setSpec:       func(spec *schema.InternetGatewaySpec) { builder.spec = spec },
		},
		setTenant:    func(tenant string) { builder.metadata.Tenant(tenant) },
		setWorkspace: func(workspace string) { builder.metadata.Workspace(workspace) },
		setRegion:    func(region string) { builder.metadata.Region(region) },
	})

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
	*regionalNetworkResourceBuilder[RouteTableBuilder, schema.RouteTableSpec]
	metadata *RegionalNetworkResourceMetadataBuilder
	spec     *schema.RouteTableSpec
}

func NewRouteTableBuilder() *RouteTableBuilder {
	builder := &RouteTableBuilder{
		metadata: NewRegionalNetworkResourceMetadataBuilder(),
		spec:     &schema.RouteTableSpec{},
	}

	builder.regionalNetworkResourceBuilder = newRegionalNetworkResourceBuilder(newRegionalNetworkResourceBuilderParams[RouteTableBuilder, schema.RouteTableSpec]{
		newGlobalResourceBuilderParams: &newGlobalResourceBuilderParams[RouteTableBuilder, schema.RouteTableSpec]{
			parent:        builder,
			setName:       func(name string) { builder.metadata.setName(name) },
			setProvider:   func(provider string) { builder.metadata.setProvider(provider) },
			setResource:   func(resource string) { builder.metadata.setResource(resource) },
			setApiVersion: func(apiVersion string) { builder.metadata.setApiVersion(apiVersion) },
		},
		setTenant:    func(tenant string) { builder.metadata.Tenant(tenant) },
		setWorkspace: func(workspace string) { builder.metadata.Workspace(workspace) },
		setNetwork:   func(network string) { builder.metadata.Network(network) },
		setRegion:    func(region string) { builder.metadata.Region(region) },
	})

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
	*regionalNetworkResourceBuilder[SubnetBuilder, schema.SubnetSpec]
	metadata *RegionalNetworkResourceMetadataBuilder
	spec     *schema.SubnetSpec
}

func NewSubnetBuilder() *SubnetBuilder {
	builder := &SubnetBuilder{
		metadata: NewRegionalNetworkResourceMetadataBuilder(),
		spec:     &schema.SubnetSpec{},
	}

	builder.regionalNetworkResourceBuilder = newRegionalNetworkResourceBuilder(newRegionalNetworkResourceBuilderParams[SubnetBuilder, schema.SubnetSpec]{
		newGlobalResourceBuilderParams: &newGlobalResourceBuilderParams[SubnetBuilder, schema.SubnetSpec]{
			parent:        builder,
			setName:       func(name string) { builder.metadata.setName(name) },
			setProvider:   func(provider string) { builder.metadata.setProvider(provider) },
			setResource:   func(resource string) { builder.metadata.setResource(resource) },
			setApiVersion: func(apiVersion string) { builder.metadata.setApiVersion(apiVersion) },
		},
		setTenant:    func(tenant string) { builder.metadata.Tenant(tenant) },
		setWorkspace: func(workspace string) { builder.metadata.Workspace(workspace) },
		setNetwork:   func(network string) { builder.metadata.Network(network) },
		setRegion:    func(region string) { builder.metadata.Region(region) },
	})

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
	*regionalWorkspaceResourceBuilder[PublicIpBuilder, schema.PublicIpSpec]
	metadata *RegionalWorkspaceResourceMetadataBuilder
	spec     *schema.PublicIpSpec
}

func NewPublicIpBuilder() *PublicIpBuilder {
	builder := &PublicIpBuilder{
		metadata: NewRegionalWorkspaceResourceMetadataBuilder(),
		spec:     &schema.PublicIpSpec{},
	}

	builder.regionalWorkspaceResourceBuilder = newRegionalWorkspaceResourceBuilder(newRegionalWorkspaceResourceBuilderParams[PublicIpBuilder, schema.PublicIpSpec]{
		newGlobalResourceBuilderParams: &newGlobalResourceBuilderParams[PublicIpBuilder, schema.PublicIpSpec]{
			parent:        builder,
			setName:       func(name string) { builder.metadata.setName(name) },
			setProvider:   func(provider string) { builder.metadata.setProvider(provider) },
			setResource:   func(resource string) { builder.metadata.setResource(resource) },
			setApiVersion: func(apiVersion string) { builder.metadata.setApiVersion(apiVersion) },
			setSpec:       func(spec *schema.PublicIpSpec) { builder.spec = spec },
		},
		setTenant:    func(tenant string) { builder.metadata.Tenant(tenant) },
		setWorkspace: func(workspace string) { builder.metadata.Workspace(workspace) },
		setRegion:    func(region string) { builder.metadata.Region(region) },
	})

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
	*regionalWorkspaceResourceBuilder[NicBuilder, schema.NicSpec]
	metadata *RegionalWorkspaceResourceMetadataBuilder
	spec     *schema.NicSpec
}

func NewNicBuilder() *NicBuilder {
	builder := &NicBuilder{
		metadata: NewRegionalWorkspaceResourceMetadataBuilder(),
		spec:     &schema.NicSpec{},
	}

	builder.regionalWorkspaceResourceBuilder = newRegionalWorkspaceResourceBuilder(newRegionalWorkspaceResourceBuilderParams[NicBuilder, schema.NicSpec]{
		newGlobalResourceBuilderParams: &newGlobalResourceBuilderParams[NicBuilder, schema.NicSpec]{
			parent:        builder,
			setName:       func(name string) { builder.metadata.setName(name) },
			setProvider:   func(provider string) { builder.metadata.setProvider(provider) },
			setResource:   func(resource string) { builder.metadata.setResource(resource) },
			setApiVersion: func(apiVersion string) { builder.metadata.setApiVersion(apiVersion) },
			setSpec:       func(spec *schema.NicSpec) { builder.spec = spec },
		},
		setTenant:    func(tenant string) { builder.metadata.Tenant(tenant) },
		setWorkspace: func(workspace string) { builder.metadata.Workspace(workspace) },
		setRegion:    func(region string) { builder.metadata.Region(region) },
	})

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
	*regionalWorkspaceResourceBuilder[SecurityGroupBuilder, schema.SecurityGroupSpec]
	metadata *RegionalWorkspaceResourceMetadataBuilder
	spec     *schema.SecurityGroupSpec
}

func NewSecurityGroupBuilder() *SecurityGroupBuilder {
	builder := &SecurityGroupBuilder{
		metadata: NewRegionalWorkspaceResourceMetadataBuilder(),
		spec:     &schema.SecurityGroupSpec{},
	}

	builder.regionalWorkspaceResourceBuilder = newRegionalWorkspaceResourceBuilder(newRegionalWorkspaceResourceBuilderParams[SecurityGroupBuilder, schema.SecurityGroupSpec]{
		newGlobalResourceBuilderParams: &newGlobalResourceBuilderParams[SecurityGroupBuilder, schema.SecurityGroupSpec]{
			parent:        builder,
			setName:       func(name string) { builder.metadata.setName(name) },
			setProvider:   func(provider string) { builder.metadata.setProvider(provider) },
			setResource:   func(resource string) { builder.metadata.setResource(resource) },
			setApiVersion: func(apiVersion string) { builder.metadata.setApiVersion(apiVersion) },
			setSpec:       func(spec *schema.SecurityGroupSpec) { builder.spec = spec },
		},
		setTenant:    func(tenant string) { builder.metadata.Tenant(tenant) },
		setWorkspace: func(workspace string) { builder.metadata.Workspace(workspace) },
		setRegion:    func(region string) { builder.metadata.Region(region) },
	})

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
