package secatest

import (
	"time"

	"github.com/eu-sovereign-cloud/go-sdk/pkg/spec/schema"
	"k8s.io/utils/ptr"
)

func buildResponseStatusConditions(state string) []schema.StatusCondition {
	return []schema.StatusCondition{
		{
			LastTransitionAt: time.Now(),
			State:            schema.ResourceState(state),
		},
	}
}

func buildResponseStatus[T any](state string, ctor func(*schema.ResourceState, []schema.StatusCondition) *T) *T {
	return ctor(ptr.To(schema.ResourceState(state)), buildResponseStatusConditions(state))
}

// Authorization

func NewRoleStatus(state string) *schema.RoleStatus {
	return buildResponseStatus(state, func(s *schema.ResourceState, c []schema.StatusCondition) *schema.RoleStatus {
		return &schema.RoleStatus{
			State:      s,
			Conditions: c,
		}
	})
}

func NewRoleAssignmentStatus(state string) *schema.RoleAssignmentStatus {
	return buildResponseStatus(state, func(s *schema.ResourceState, c []schema.StatusCondition) *schema.RoleAssignmentStatus {
		return &schema.RoleAssignmentStatus{
			State:      s,
			Conditions: c,
		}
	})
}

// Compute

func NewInstanceStatus(state string) *schema.InstanceStatus {
	return buildResponseStatus(state, func(s *schema.ResourceState, c []schema.StatusCondition) *schema.InstanceStatus {
		return &schema.InstanceStatus{
			State:      s,
			Conditions: c,
		}
	})
}

// Storage

func NewBlockStorageStatus(state string) *schema.BlockStorageStatus {
	return buildResponseStatus(state, func(s *schema.ResourceState, c []schema.StatusCondition) *schema.BlockStorageStatus {
		return &schema.BlockStorageStatus{
			State:      s,
			Conditions: c,
		}
	})
}

func NewImageStatus(state string) *schema.ImageStatus {
	return buildResponseStatus(state, func(s *schema.ResourceState, c []schema.StatusCondition) *schema.ImageStatus {
		return &schema.ImageStatus{
			State:      s,
			Conditions: c,
		}
	})
}

// Network

func NewNetworkStatus(state string) *schema.NetworkStatus {
	return buildResponseStatus(state, func(s *schema.ResourceState, c []schema.StatusCondition) *schema.NetworkStatus {
		return &schema.NetworkStatus{
			State:      s,
			Conditions: c,
		}
	})
}

func NewSubnetStatus(state string) *schema.SubnetStatus {
	return buildResponseStatus(state, func(s *schema.ResourceState, c []schema.StatusCondition) *schema.SubnetStatus {
		return &schema.SubnetStatus{
			State:      s,
			Conditions: c,
		}
	})
}

func NewRouteTableStatus(state string) *schema.RouteTableStatus {
	return buildResponseStatus(state, func(s *schema.ResourceState, c []schema.StatusCondition) *schema.RouteTableStatus {
		return &schema.RouteTableStatus{
			State:      s,
			Conditions: c,
		}
	})
}

func NewInternetGatewayStatus(state string) *schema.InternetGatewayStatus {
	return buildResponseStatus(state, func(s *schema.ResourceState, c []schema.StatusCondition) *schema.InternetGatewayStatus {
		return &schema.InternetGatewayStatus{
			State:      s,
			Conditions: c,
		}
	})
}

func NewSecurityGroupStatus(state string) *schema.SecurityGroupStatus {
	return buildResponseStatus(state, func(s *schema.ResourceState, c []schema.StatusCondition) *schema.SecurityGroupStatus {
		return &schema.SecurityGroupStatus{
			State:      s,
			Conditions: c,
		}
	})
}

func NewNicStatus(state string) *schema.NicStatus {
	return buildResponseStatus(state, func(s *schema.ResourceState, c []schema.StatusCondition) *schema.NicStatus {
		return &schema.NicStatus{
			State:      s,
			Conditions: c,
		}
	})
}

func NewPublicIpStatus(state string) *schema.PublicIpStatus {
	return buildResponseStatus(state, func(s *schema.ResourceState, c []schema.StatusCondition) *schema.PublicIpStatus {
		return &schema.PublicIpStatus{
			State:      s,
			Conditions: c,
		}
	})
}

// Workspace

func NewWorkspaceStatus(state string) *schema.WorkspaceStatus {
	return buildResponseStatus(state, func(s *schema.ResourceState, c []schema.StatusCondition) *schema.WorkspaceStatus {
		return &schema.WorkspaceStatus{
			State:      s,
			Conditions: c,
		}
	})
}
