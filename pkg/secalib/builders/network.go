package builders

import (
	"github.com/eu-sovereign-cloud/go-sdk/pkg/secalib/generators"
	"github.com/eu-sovereign-cloud/go-sdk/pkg/spec/schema"
)

// Network

type NetworkMetadataBuilder struct {
	*RegionalWorkspaceResourceMetadataBuilder
}

func NewNetworkMetadataBuilder() *NetworkMetadataBuilder {
	builder := &NetworkMetadataBuilder{
		RegionalWorkspaceResourceMetadataBuilder: newRegionalWorkspaceResourceMetadataBuilder(),
	}

	return builder
}

func (builder *NetworkMetadataBuilder) BuildResponse() (*schema.RegionalWorkspaceResourceMetadata, error) {

	medatata, err := builder.kind(schema.RegionalWorkspaceResourceMetadataKindResourceKindNetwork).buildResponse()
	if err != nil {
		return nil, err
	}

	resource := generators.GenerateNetworkResource(builder.metadata.Tenant, builder.metadata.Workspace, builder.metadata.Name)
	medatata.Resource = resource

	return medatata, nil
}

type NetworkBuilder struct {
	*regionalWorkspaceResourceBuilder[NetworkBuilder, schema.NetworkSpec]
	metadata *NetworkMetadataBuilder
	labels   schema.Labels
	spec     *schema.NetworkSpec
}

func NewNetworkBuilder() *NetworkBuilder {
	builder := &NetworkBuilder{
		metadata: NewNetworkMetadataBuilder(),
		spec:     &schema.NetworkSpec{},
	}

	builder.regionalWorkspaceResourceBuilder = newRegionalWorkspaceResourceBuilder(newRegionalWorkspaceResourceBuilderParams[NetworkBuilder, schema.NetworkSpec]{
		newGlobalResourceBuilderParams: &newGlobalResourceBuilderParams[NetworkBuilder, schema.NetworkSpec]{
			parent:        builder,
			setName:       func(name string) { builder.metadata.setName(name) },
			setProvider:   func(provider string) { builder.metadata.setProvider(provider) },
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

	metadata, err := builder.metadata.buildResponse()
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

type InternetGatewayMetadataBuilder struct {
	*RegionalWorkspaceResourceMetadataBuilder
}

func NewInternetGatewayMetadataBuilder() *InternetGatewayMetadataBuilder {
	builder := &InternetGatewayMetadataBuilder{
		RegionalWorkspaceResourceMetadataBuilder: newRegionalWorkspaceResourceMetadataBuilder(),
	}

	return builder
}

func (builder *InternetGatewayMetadataBuilder) BuildResponse() (*schema.RegionalWorkspaceResourceMetadata, error) {

	medatata, err := builder.kind(schema.RegionalWorkspaceResourceMetadataKindResourceKindInternetGateway).buildResponse()
	if err != nil {
		return nil, err
	}

	resource := generators.GenerateInternetGatewayResource(builder.metadata.Tenant, builder.metadata.Workspace, builder.metadata.Name)
	medatata.Resource = resource

	return medatata, nil
}

type InternetGatewayBuilder struct {
	*regionalWorkspaceResourceBuilder[InternetGatewayBuilder, schema.InternetGatewaySpec]
	metadata *InternetGatewayMetadataBuilder
	labels   schema.Labels
	spec     *schema.InternetGatewaySpec
}

func NewInternetGatewayBuilder() *InternetGatewayBuilder {
	builder := &InternetGatewayBuilder{
		metadata: NewInternetGatewayMetadataBuilder(),
		spec:     &schema.InternetGatewaySpec{},
	}

	builder.regionalWorkspaceResourceBuilder = newRegionalWorkspaceResourceBuilder(newRegionalWorkspaceResourceBuilderParams[InternetGatewayBuilder, schema.InternetGatewaySpec]{
		newGlobalResourceBuilderParams: &newGlobalResourceBuilderParams[InternetGatewayBuilder, schema.InternetGatewaySpec]{
			parent:        builder,
			setName:       func(name string) { builder.metadata.setName(name) },
			setProvider:   func(provider string) { builder.metadata.setProvider(provider) },
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

	metadata, err := builder.metadata.buildResponse()
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

type RouteTableMetadataBuilder struct {
	*RegionalNetworkResourceMetadataBuilder
}

func NewRouteTableMetadataBuilder() *RouteTableMetadataBuilder {
	builder := &RouteTableMetadataBuilder{
		RegionalNetworkResourceMetadataBuilder: newRegionalNetworkResourceMetadataBuilder(),
	}

	return builder
}

func (builder *RouteTableMetadataBuilder) BuildResponse() (*schema.RegionalNetworkResourceMetadata, error) {

	medatata, err := builder.kind(schema.RegionalNetworkResourceMetadataKindResourceKindRoutingTable).buildResponse()
	if err != nil {
		return nil, err
	}

	resource := generators.GenerateRouteTableResource(builder.metadata.Tenant, builder.metadata.Workspace, builder.metadata.Network, builder.metadata.Name)
	medatata.Resource = resource

	return medatata, nil
}

type RouteTableBuilder struct {
	*regionalNetworkResourceBuilder[RouteTableBuilder, schema.RouteTableSpec]
	metadata *RouteTableMetadataBuilder
	labels   schema.Labels
	spec     *schema.RouteTableSpec
}

func NewRouteTableBuilder() *RouteTableBuilder {
	builder := &RouteTableBuilder{
		metadata: NewRouteTableMetadataBuilder(),
		spec:     &schema.RouteTableSpec{},
	}

	builder.regionalNetworkResourceBuilder = newRegionalNetworkResourceBuilder(newRegionalNetworkResourceBuilderParams[RouteTableBuilder, schema.RouteTableSpec]{
		newGlobalResourceBuilderParams: &newGlobalResourceBuilderParams[RouteTableBuilder, schema.RouteTableSpec]{
			parent:        builder,
			setName:       func(name string) { builder.metadata.setName(name) },
			setProvider:   func(provider string) { builder.metadata.setProvider(provider) },
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

	metadata, err := builder.metadata.buildResponse()
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

type SubnetMetadataBuilder struct {
	*RegionalNetworkResourceMetadataBuilder
}

func NewSubnetMetadataBuilder() *SubnetMetadataBuilder {
	builder := &SubnetMetadataBuilder{
		RegionalNetworkResourceMetadataBuilder: newRegionalNetworkResourceMetadataBuilder(),
	}

	return builder
}

func (builder *SubnetMetadataBuilder) BuildResponse() (*schema.RegionalNetworkResourceMetadata, error) {

	medatata, err := builder.kind(schema.RegionalNetworkResourceMetadataKindResourceKindSubnet).buildResponse()
	if err != nil {
		return nil, err
	}

	resource := generators.GenerateSubnetResource(builder.metadata.Tenant, builder.metadata.Workspace, builder.metadata.Network, builder.metadata.Name)
	medatata.Resource = resource

	return medatata, nil
}

type SubnetBuilder struct {
	*regionalNetworkResourceBuilder[SubnetBuilder, schema.SubnetSpec]
	metadata *SubnetMetadataBuilder
	labels   schema.Labels
	spec     *schema.SubnetSpec
}

func NewSubnetBuilder() *SubnetBuilder {
	builder := &SubnetBuilder{
		metadata: NewSubnetMetadataBuilder(),
		spec:     &schema.SubnetSpec{},
	}

	builder.regionalNetworkResourceBuilder = newRegionalNetworkResourceBuilder(newRegionalNetworkResourceBuilderParams[SubnetBuilder, schema.SubnetSpec]{
		newGlobalResourceBuilderParams: &newGlobalResourceBuilderParams[SubnetBuilder, schema.SubnetSpec]{
			parent:        builder,
			setName:       func(name string) { builder.metadata.setName(name) },
			setProvider:   func(provider string) { builder.metadata.setProvider(provider) },
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

	metadata, err := builder.metadata.buildResponse()
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

type PublicIpMetadataBuilder struct {
	*RegionalWorkspaceResourceMetadataBuilder
}

func NewPublicIpMetadataBuilder() *PublicIpMetadataBuilder {
	builder := &PublicIpMetadataBuilder{
		RegionalWorkspaceResourceMetadataBuilder: newRegionalWorkspaceResourceMetadataBuilder(),
	}

	return builder
}

func (builder *PublicIpMetadataBuilder) BuildResponse() (*schema.RegionalWorkspaceResourceMetadata, error) {

	medatata, err := builder.kind(schema.RegionalWorkspaceResourceMetadataKindResourceKindPublicIP).buildResponse()
	if err != nil {
		return nil, err
	}

	resource := generators.GeneratePublicIpResource(builder.metadata.Tenant, builder.metadata.Workspace, builder.metadata.Name)
	medatata.Resource = resource

	return medatata, nil
}

type PublicIpBuilder struct {
	*regionalWorkspaceResourceBuilder[PublicIpBuilder, schema.PublicIpSpec]
	metadata *PublicIpMetadataBuilder
	labels   schema.Labels
	spec     *schema.PublicIpSpec
}

func NewPublicIpBuilder() *PublicIpBuilder {
	builder := &PublicIpBuilder{
		metadata: NewPublicIpMetadataBuilder(),
		spec:     &schema.PublicIpSpec{},
	}

	builder.regionalWorkspaceResourceBuilder = newRegionalWorkspaceResourceBuilder(newRegionalWorkspaceResourceBuilderParams[PublicIpBuilder, schema.PublicIpSpec]{
		newGlobalResourceBuilderParams: &newGlobalResourceBuilderParams[PublicIpBuilder, schema.PublicIpSpec]{
			parent:        builder,
			setName:       func(name string) { builder.metadata.setName(name) },
			setProvider:   func(provider string) { builder.metadata.setProvider(provider) },
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

	metadata, err := builder.metadata.buildResponse()
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

type NicMetadataBuilder struct {
	*RegionalWorkspaceResourceMetadataBuilder
}

func NewNicMetadataBuilder() *NicMetadataBuilder {
	builder := &NicMetadataBuilder{
		RegionalWorkspaceResourceMetadataBuilder: newRegionalWorkspaceResourceMetadataBuilder(),
	}

	return builder
}

func (builder *NicMetadataBuilder) BuildResponse() (*schema.RegionalWorkspaceResourceMetadata, error) {

	medatata, err := builder.kind(schema.RegionalWorkspaceResourceMetadataKindResourceKindNic).buildResponse()
	if err != nil {
		return nil, err
	}

	resource := generators.GenerateNicResource(builder.metadata.Tenant, builder.metadata.Workspace, builder.metadata.Name)
	medatata.Resource = resource

	return medatata, nil
}

type NicBuilder struct {
	*regionalWorkspaceResourceBuilder[NicBuilder, schema.NicSpec]
	metadata *NicMetadataBuilder
	labels   schema.Labels
	spec     *schema.NicSpec
}

func NewNicBuilder() *NicBuilder {
	builder := &NicBuilder{
		metadata: NewNicMetadataBuilder(),
		spec:     &schema.NicSpec{},
	}

	builder.regionalWorkspaceResourceBuilder = newRegionalWorkspaceResourceBuilder(newRegionalWorkspaceResourceBuilderParams[NicBuilder, schema.NicSpec]{
		newGlobalResourceBuilderParams: &newGlobalResourceBuilderParams[NicBuilder, schema.NicSpec]{
			parent:        builder,
			setName:       func(name string) { builder.metadata.setName(name) },
			setProvider:   func(provider string) { builder.metadata.setProvider(provider) },
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

	metadata, err := builder.metadata.buildResponse()
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

type SecurityGroupMetadataBuilder struct {
	*RegionalWorkspaceResourceMetadataBuilder
}

func NewSecurityGroupMetadataBuilder() *SecurityGroupMetadataBuilder {
	builder := &SecurityGroupMetadataBuilder{
		RegionalWorkspaceResourceMetadataBuilder: newRegionalWorkspaceResourceMetadataBuilder(),
	}

	return builder
}

func (builder *SecurityGroupMetadataBuilder) BuildResponse() (*schema.RegionalWorkspaceResourceMetadata, error) {

	medatata, err := builder.kind(schema.RegionalWorkspaceResourceMetadataKindResourceKindSecurityGroup).buildResponse()
	if err != nil {
		return nil, err
	}

	resource := generators.GenerateSecurityGroupResource(builder.metadata.Tenant, builder.metadata.Workspace, builder.metadata.Name)
	medatata.Resource = resource

	return medatata, nil
}

type SecurityGroupBuilder struct {
	*regionalWorkspaceResourceBuilder[SecurityGroupBuilder, schema.SecurityGroupSpec]
	metadata *SecurityGroupMetadataBuilder
	labels   schema.Labels
	spec     *schema.SecurityGroupSpec
}

func NewSecurityGroupBuilder() *SecurityGroupBuilder {
	builder := &SecurityGroupBuilder{
		metadata: NewSecurityGroupMetadataBuilder(),
		spec:     &schema.SecurityGroupSpec{},
	}

	builder.regionalWorkspaceResourceBuilder = newRegionalWorkspaceResourceBuilder(newRegionalWorkspaceResourceBuilderParams[SecurityGroupBuilder, schema.SecurityGroupSpec]{
		newGlobalResourceBuilderParams: &newGlobalResourceBuilderParams[SecurityGroupBuilder, schema.SecurityGroupSpec]{
			parent:        builder,
			setName:       func(name string) { builder.metadata.setName(name) },
			setProvider:   func(provider string) { builder.metadata.setProvider(provider) },
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

	metadata, err := builder.metadata.buildResponse()
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
