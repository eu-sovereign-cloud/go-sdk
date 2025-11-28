package builders

import (
	"github.com/eu-sovereign-cloud/go-sdk/pkg/spec/schema"
)

// Network

type NetworkBuilder struct {
	*regionalWorkspaceResourceBuilder[NetworkBuilder, schema.NetworkSpec]
	metadata *RegionalWorkspaceResourceMetadataBuilder
	labels   schema.Labels
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
			setLabels:     func(labels schema.Labels) { builder.labels = labels },
			setSpec:       func(spec *schema.NetworkSpec) { builder.spec = spec },
		},
		setTenant:    func(tenant string) { builder.metadata.Tenant(tenant) },
		setWorkspace: func(workspace string) { builder.metadata.Workspace(workspace) },
		setRegion:    func(region string) { builder.metadata.Region(region) },
	})

	return builder
}

func (builder *NetworkBuilder) validateSpec() error {
	if err := validateRequired(builder.validator,
		builder.spec,
		builder.spec.Cidr,
		builder.spec.SkuRef,
		builder.spec.RouteTableRef,
	); err != nil {
		return err
	}
	if err := validateOneRequired(builder.validator,
		builder.spec.Cidr.Ipv4,
		builder.spec.Cidr.Ipv6,
	); err != nil {
		return err
	}

	return nil
}

func (builder *NetworkBuilder) BuildRequest() (*schema.Network, error) {
	if err := builder.validateSpec(); err != nil {
		return nil, err
	}

	return &schema.Network{
		Metadata: nil,
		Labels:   builder.labels,
		Spec:     *builder.spec,
		Status:   nil,
	}, nil
}

func (builder *NetworkBuilder) BuildResponse() (*schema.Network, error) {
	if err := builder.validateSpec(); err != nil {
		return nil, err
	}

	metadata, err := builder.metadata.Kind(schema.RegionalWorkspaceResourceMetadataKindResourceKindNetwork).BuildResponse()
	if err != nil {
		return nil, err
	}

	return &schema.Network{
		Metadata: metadata,
		Labels:   builder.labels,
		Spec:     *builder.spec,
		Status:   &schema.NetworkStatus{},
	}, nil
}

// Internet Gateway

type InternetGatewayBuilder struct {
	*regionalWorkspaceResourceBuilder[InternetGatewayBuilder, schema.InternetGatewaySpec]
	metadata *RegionalWorkspaceResourceMetadataBuilder
	labels   schema.Labels
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
			setLabels:     func(labels schema.Labels) { builder.labels = labels },
			setSpec:       func(spec *schema.InternetGatewaySpec) { builder.spec = spec },
		},
		setTenant:    func(tenant string) { builder.metadata.Tenant(tenant) },
		setWorkspace: func(workspace string) { builder.metadata.Workspace(workspace) },
		setRegion:    func(region string) { builder.metadata.Region(region) },
	})

	return builder
}

func (builder *InternetGatewayBuilder) validateSpec() error {
	if err := validateRequired(builder.validator,
		builder.spec,
		builder.spec.EgressOnly,
	); err != nil {
		return err
	}

	return nil
}

func (builder *InternetGatewayBuilder) BuildRequest() (*schema.InternetGateway, error) {
	if err := builder.validateSpec(); err != nil {
		return nil, err
	}

	return &schema.InternetGateway{
		Metadata: nil,
		Labels:   builder.labels,
		Spec:     *builder.spec,
		Status:   nil,
	}, nil
}

func (builder *InternetGatewayBuilder) BuildResponse() (*schema.InternetGateway, error) {
	if err := builder.validateSpec(); err != nil {
		return nil, err
	}

	metadata, err := builder.metadata.Kind(schema.RegionalWorkspaceResourceMetadataKindResourceKindInternetGateway).BuildResponse()
	if err != nil {
		return nil, err
	}

	return &schema.InternetGateway{
		Metadata: metadata,
		Labels:   builder.labels,
		Spec:     *builder.spec,
		Status:   &schema.InternetGatewayStatus{},
	}, nil
}

// Route Table

type RouteTableBuilder struct {
	*regionalNetworkResourceBuilder[RouteTableBuilder, schema.RouteTableSpec]
	metadata *RegionalNetworkResourceMetadataBuilder
	labels   schema.Labels
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
			setLabels:     func(labels schema.Labels) { builder.labels = labels },
			setSpec:       func(spec *schema.RouteTableSpec) { builder.spec = spec },
		},
		setTenant:    func(tenant string) { builder.metadata.Tenant(tenant) },
		setWorkspace: func(workspace string) { builder.metadata.Workspace(workspace) },
		setNetwork:   func(network string) { builder.metadata.Network(network) },
		setRegion:    func(region string) { builder.metadata.Region(region) },
	})

	return builder
}

func (builder *RouteTableBuilder) validateSpec() error {
	if err := validateRequired(builder.validator,
		builder.spec.Routes,
	); err != nil {
		return err
	}

	// Validate each route
	for _, route := range builder.spec.Routes {
		if err := validateRequired(builder.validator,
			route.DestinationCidrBlock,
			route.TargetRef,
		); err != nil {
			return err
		}
	}

	return nil
}

func (builder *RouteTableBuilder) BuildRequest() (*schema.RouteTable, error) {
	if err := builder.validateSpec(); err != nil {
		return nil, err
	}

	return &schema.RouteTable{
		Metadata: nil,
		Labels:   builder.labels,
		Spec:     *builder.spec,
		Status:   nil,
	}, nil
}

func (builder *RouteTableBuilder) BuildResponse() (*schema.RouteTable, error) {
	if err := builder.validateSpec(); err != nil {
		return nil, err
	}

	metadata, err := builder.metadata.Kind(schema.RegionalNetworkResourceMetadataKindResourceKindRoutingTable).BuildResponse()
	if err != nil {
		return nil, err
	}

	return &schema.RouteTable{
		Metadata: metadata,
		Labels:   builder.labels,
		Spec:     *builder.spec,
		Status:   &schema.RouteTableStatus{},
	}, nil
}

// Subnet

type SubnetBuilder struct {
	*regionalNetworkResourceBuilder[SubnetBuilder, schema.SubnetSpec]
	metadata *RegionalNetworkResourceMetadataBuilder
	labels   schema.Labels
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
			setLabels:     func(labels schema.Labels) { builder.labels = labels },
			setSpec:       func(spec *schema.SubnetSpec) { builder.spec = spec },
		},
		setTenant:    func(tenant string) { builder.metadata.Tenant(tenant) },
		setWorkspace: func(workspace string) { builder.metadata.Workspace(workspace) },
		setNetwork:   func(network string) { builder.metadata.Network(network) },
		setRegion:    func(region string) { builder.metadata.Region(region) },
	})

	return builder
}

func (builder *SubnetBuilder) validateSpec() error {
	if err := validateRequired(builder.validator,
		builder.spec,
		builder.spec.Cidr,
		builder.spec.Zone,
	); err != nil {
		return err
	}

	return nil
}

func (builder *SubnetBuilder) BuildRequest() (*schema.Subnet, error) {
	if err := builder.validateSpec(); err != nil {
		return nil, err
	}

	return &schema.Subnet{
		Metadata: nil,
		Labels:   builder.labels,
		Spec:     *builder.spec,
		Status:   nil,
	}, nil
}

func (builder *SubnetBuilder) BuildResponse() (*schema.Subnet, error) {
	if err := builder.validateSpec(); err != nil {
		return nil, err
	}

	metadata, err := builder.metadata.Kind(schema.RegionalNetworkResourceMetadataKindResourceKindSubnet).BuildResponse()
	if err != nil {
		return nil, err
	}

	return &schema.Subnet{
		Metadata: metadata,
		Labels:   builder.labels,
		Spec:     *builder.spec,
		Status:   &schema.SubnetStatus{},
	}, nil
}

// Public Ip

type PublicIpBuilder struct {
	*regionalWorkspaceResourceBuilder[PublicIpBuilder, schema.PublicIpSpec]
	metadata *RegionalWorkspaceResourceMetadataBuilder
	labels   schema.Labels
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
			setLabels:     func(labels schema.Labels) { builder.labels = labels },
			setSpec:       func(spec *schema.PublicIpSpec) { builder.spec = spec },
		},
		setTenant:    func(tenant string) { builder.metadata.Tenant(tenant) },
		setWorkspace: func(workspace string) { builder.metadata.Workspace(workspace) },
		setRegion:    func(region string) { builder.metadata.Region(region) },
	})

	return builder
}

func (builder *PublicIpBuilder) validateSpec() error {
	if err := validateRequired(builder.validator,
		builder.spec,
		builder.spec.Version,
	); err != nil {
		return err
	}

	return nil
}

func (builder *PublicIpBuilder) BuildRequest() (*schema.PublicIp, error) {
	if err := builder.validateSpec(); err != nil {
		return nil, err
	}

	return &schema.PublicIp{
		Metadata: nil,
		Labels:   builder.labels,
		Spec:     *builder.spec,
		Status:   nil,
	}, nil
}

func (builder *PublicIpBuilder) BuildResponse() (*schema.PublicIp, error) {
	if err := builder.validateSpec(); err != nil {
		return nil, err
	}

	metadata, err := builder.metadata.Kind(schema.RegionalWorkspaceResourceMetadataKindResourceKindPublicIP).BuildResponse()
	if err != nil {
		return nil, err
	}

	return &schema.PublicIp{
		Metadata: metadata,
		Labels:   builder.labels,
		Spec:     *builder.spec,
		Status:   &schema.PublicIpStatus{},
	}, nil
}

// Nic

type NicBuilder struct {
	*regionalWorkspaceResourceBuilder[NicBuilder, schema.NicSpec]
	metadata *RegionalWorkspaceResourceMetadataBuilder
	labels   schema.Labels
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
			setLabels:     func(labels schema.Labels) { builder.labels = labels },
			setSpec:       func(spec *schema.NicSpec) { builder.spec = spec },
		},
		setTenant:    func(tenant string) { builder.metadata.Tenant(tenant) },
		setWorkspace: func(workspace string) { builder.metadata.Workspace(workspace) },
		setRegion:    func(region string) { builder.metadata.Region(region) },
	})

	return builder
}

func (builder *NicBuilder) validateSpec() error {
	if err := validateRequired(builder.validator,
		builder.spec,
		builder.spec.Addresses,
		builder.spec.SubnetRef,
	); err != nil {
		return err
	}

	return nil
}

func (builder *NicBuilder) BuildRequest() (*schema.Nic, error) {
	if err := builder.validateSpec(); err != nil {
		return nil, err
	}

	return &schema.Nic{
		Metadata: nil,
		Labels:   builder.labels,
		Spec:     *builder.spec,
		Status:   nil,
	}, nil
}

func (builder *NicBuilder) BuildResponse() (*schema.Nic, error) {
	if err := builder.validateSpec(); err != nil {
		return nil, err
	}

	metadata, err := builder.metadata.Kind(schema.RegionalWorkspaceResourceMetadataKindResourceKindNic).BuildResponse()
	if err != nil {
		return nil, err
	}

	return &schema.Nic{
		Metadata: metadata,
		Labels:   builder.labels,
		Spec:     *builder.spec,
		Status:   &schema.NicStatus{},
	}, nil
}

// Security Group

type SecurityGroupBuilder struct {
	*regionalWorkspaceResourceBuilder[SecurityGroupBuilder, schema.SecurityGroupSpec]
	metadata *RegionalWorkspaceResourceMetadataBuilder
	labels   schema.Labels
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
			setLabels:     func(labels schema.Labels) { builder.labels = labels },
			setSpec:       func(spec *schema.SecurityGroupSpec) { builder.spec = spec },
		},
		setTenant:    func(tenant string) { builder.metadata.Tenant(tenant) },
		setWorkspace: func(workspace string) { builder.metadata.Workspace(workspace) },
		setRegion:    func(region string) { builder.metadata.Region(region) },
	})

	return builder
}

func (builder *SecurityGroupBuilder) validateSpec() error {
	if err := validateRequired(builder.validator,
		builder.spec,
		builder.spec.Rules,
	); err != nil {
		return err
	}

	// Validate each rule
	for _, rule := range builder.spec.Rules {
		if err := validateRequired(builder.validator,
			rule.Direction,
		); err != nil {
			return err
		}
	}

	return nil
}

func (builder *SecurityGroupBuilder) BuildRequest() (*schema.SecurityGroup, error) {
	if err := builder.validateSpec(); err != nil {
		return nil, err
	}

	return &schema.SecurityGroup{
		Metadata: nil,
		Labels:   builder.labels,
		Spec:     *builder.spec,
		Status:   nil,
	}, nil
}

func (builder *SecurityGroupBuilder) BuildResponse() (*schema.SecurityGroup, error) {
	if err := builder.validateSpec(); err != nil {
		return nil, err
	}

	metadata, err := builder.metadata.Kind(schema.RegionalWorkspaceResourceMetadataKindResourceKindSecurityGroup).BuildResponse()
	if err != nil {
		return nil, err
	}

	return &schema.SecurityGroup{
		Metadata: metadata,
		Labels:   builder.labels,
		Spec:     *builder.spec,
		Status:   &schema.SecurityGroupStatus{},
	}, nil
}
