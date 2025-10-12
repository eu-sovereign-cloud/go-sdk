package secatest

type MetadataResponseV1 struct {
	Name      string
	Tenant    string
	Workspace *string
	Network   *string
}
type StatusResponseV1 struct {
	State string
}

// Network
type NetworkSkuResponseV1 struct {
	Metadata  MetadataResponseV1
	Tier      string
	Bandwidth int
	Packets   int
}
type NetworkResponseV1 struct {
	Metadata      MetadataResponseV1
	RouteTableRef string
	Status        StatusResponseV1
}
type SubnetResponseV1 struct {
	Metadata MetadataResponseV1
	SkuRef   string
	Status   StatusResponseV1
}
type RouteTableResponseV1 struct {
	Metadata       MetadataResponseV1
	RouteCidrBlock string
	RouteTargetRef string
	Status         StatusResponseV1
}
type InternetGatewayResponseV1 struct {
	Metadata   MetadataResponseV1
	EgressOnly bool
	Status     StatusResponseV1
}
type SecurityGroupResponseV1 struct {
	Metadata      MetadataResponseV1
	RuleDirection string
	Status        StatusResponseV1
}
type NicResponseV1 struct {
	Metadata  MetadataResponseV1
	SubnetRef string
	Status    StatusResponseV1
}
type PublicIpResponseV1 struct {
	Metadata MetadataResponseV1
	Address  string
	Status   StatusResponseV1
}
