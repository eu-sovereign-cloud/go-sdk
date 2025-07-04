package secatest

// Sku
type NetworkSkusResponseV1 struct {
	Tenant string
	Name   string
}
type NetworkSkuResponseV1 struct {
	Tenant string
	Name   string
}

// Network
type NetworksResponseV1 struct {
	Tenant    string
	Workspace string
	Name      string
}
type NetworkResponseV1 struct {
	Tenant    string
	Workspace string
	Name      string
}

// Subnet
type SubnetsResponseV1 struct {
	Tenant    string
	Workspace string
	Name      string
}
type SubnetResponseV1 struct {
	Tenant    string
	Workspace string
	Name      string
}

// Route Table
type RouteTablesResponseV1 struct {
	Tenant    string
	Workspace string
	Name      string
}
type RouteTableResponseV1 struct {
	Tenant    string
	Workspace string
	Name      string
}

// Internet Gateway
type InternetGatewaysResponseV1 struct {
	Tenant    string
	Workspace string
	Name      string
}
type InternetGatewayResponseV1 struct {
	Tenant    string
	Workspace string
	Name      string
}

// Security Group
type SecurityGroupsResponseV1 struct {
	Tenant    string
	Workspace string
	Name      string
}
type SecurityGroupResponseV1 struct {
	Tenant    string
	Workspace string
	Name      string
}

// Nic
type NicsResponseV1 struct {
	Tenant    string
	Workspace string
	Name      string
}
type NicResponseV1 struct {
	Tenant    string
	Workspace string
	Name      string
}

// Pbulic Ips
type PublicIpsResponseV1 struct {
	Tenant    string
	Workspace string
	Name      string
}
type PublicIpResponseV1 struct {
	Tenant    string
	Workspace string
	Name      string
}
