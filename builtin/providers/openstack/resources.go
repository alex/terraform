package openstack

import (
	"github.com/hashicorp/terraform/helper/resource"
)

// resourceMap is the mapping of resources we support to their basic
// operations. This makes it easy to implement new resource types.
var resourceMap *resource.Map

func init() {
	resourceMap = &resource.Map{
		Mapping: map[string]resource.Resource{
			"openstack_swift_container": resource.Resource{
				ConfigValidator: resource_openstack_swift_container_validation(),
				Create:          resource_openstack_swift_container_create,
				Destroy:         resource_openstack_swift_container_destroy,
				Diff:            resource_openstack_swift_container_diff,
				Refresh:         resource_openstack_swift_container_refresh,
			},
		},
	}
}
