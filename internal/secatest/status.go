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

func buildResponseStatus[T any](state schema.ResourceState, ctor func(*schema.ResourceState, []schema.StatusCondition) *T) *T {
	return ctor(ptr.To(state), buildResponseStatusConditions(string(state)))
}

// Authorization

func NewRoleStatus(state schema.ResourceState) *schema.RoleStatus {
	return buildResponseStatus(state, func(s *schema.ResourceState, c []schema.StatusCondition) *schema.RoleStatus {
		return &schema.RoleStatus{
			State:      s,
			Conditions: c,
		}
	})
}

func NewRoleAssignmentStatus(state schema.ResourceState) *schema.RoleAssignmentStatus {
	return buildResponseStatus(state, func(s *schema.ResourceState, c []schema.StatusCondition) *schema.RoleAssignmentStatus {
		return &schema.RoleAssignmentStatus{
			State:      s,
			Conditions: c,
		}
	})
}

// Compute

func NewInstanceStatus(state schema.ResourceState) *schema.InstanceStatus {
	return buildResponseStatus(state, func(s *schema.ResourceState, c []schema.StatusCondition) *schema.InstanceStatus {
		return &schema.InstanceStatus{
			State:      s,
			Conditions: c,
		}
	})
}

// Storage

func NewBlockStorageStatus(state schema.ResourceState) *schema.BlockStorageStatus {
	return buildResponseStatus(state, func(s *schema.ResourceState, c []schema.StatusCondition) *schema.BlockStorageStatus {
		return &schema.BlockStorageStatus{
			State:      s,
			Conditions: c,
		}
	})
}

func NewImageStatus(state schema.ResourceState) *schema.ImageStatus {
	return buildResponseStatus(state, func(s *schema.ResourceState, c []schema.StatusCondition) *schema.ImageStatus {
		return &schema.ImageStatus{
			State:      s,
			Conditions: c,
		}
	})
}

// Network

func NewNetworkStatus(state schema.ResourceState) *schema.NetworkStatus {
	return buildResponseStatus(state, func(s *schema.ResourceState, c []schema.StatusCondition) *schema.NetworkStatus {
		return &schema.NetworkStatus{
			State:      s,
			Conditions: c,
		}
	})
}

func NewSubnetStatus(state schema.ResourceState) *schema.SubnetStatus {
	return buildResponseStatus(state, func(s *schema.ResourceState, c []schema.StatusCondition) *schema.SubnetStatus {
		return &schema.SubnetStatus{
			State:      s,
			Conditions: c,
		}
	})
}

func NewRouteTableStatus(state schema.ResourceState) *schema.RouteTableStatus {
	return buildResponseStatus(state, func(s *schema.ResourceState, c []schema.StatusCondition) *schema.RouteTableStatus {
		return &schema.RouteTableStatus{
			State:      s,
			Conditions: c,
		}
	})
}

func NewInternetGatewayStatus(state schema.ResourceState) *schema.InternetGatewayStatus {
	return buildResponseStatus(state, func(s *schema.ResourceState, c []schema.StatusCondition) *schema.InternetGatewayStatus {
		return &schema.InternetGatewayStatus{
			State:      s,
			Conditions: c,
		}
	})
}

func NewSecurityGroupStatus(state schema.ResourceState) *schema.SecurityGroupStatus {
	return buildResponseStatus(state, func(s *schema.ResourceState, c []schema.StatusCondition) *schema.SecurityGroupStatus {
		return &schema.SecurityGroupStatus{
			State:      s,
			Conditions: c,
		}
	})
}

func NewNicStatus(state schema.ResourceState) *schema.NicStatus {
	return buildResponseStatus(state, func(s *schema.ResourceState, c []schema.StatusCondition) *schema.NicStatus {
		return &schema.NicStatus{
			State:      s,
			Conditions: c,
		}
	})
}

func NewPublicIpStatus(state schema.ResourceState) *schema.PublicIpStatus {
	return buildResponseStatus(state, func(s *schema.ResourceState, c []schema.StatusCondition) *schema.PublicIpStatus {
		return &schema.PublicIpStatus{
			State:      s,
			Conditions: c,
		}
	})
}

// Workspace

func NewWorkspaceStatus(state schema.ResourceState) *schema.WorkspaceStatus {
	return buildResponseStatus(state, func(s *schema.ResourceState, c []schema.StatusCondition) *schema.WorkspaceStatus {
		return &schema.WorkspaceStatus{
			State:      s,
			Conditions: c,
		}
	})
}
