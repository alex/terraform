package openstack

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/config"
	"github.com/hashicorp/terraform/helper/diff"
	"github.com/hashicorp/terraform/terraform"

	"github.com/rackspace/gophercloud/openstack/storage/v1/containers"
)

func resource_openstack_swift_container_validation() *config.Validator {
	return &config.Validator{
		Required: []string{
			"name",
		},
	}
}

func resource_openstack_swift_container_create(
	s *terraform.ResourceState,
	d *terraform.ResourceDiff,
	meta interface{}) (*terraform.ResourceState, error) {

	p := meta.(*ResourceProvider)
	storageClient := p.storageClient

	rs := s.MergeDiff(d)

	name := rs.Attributes["name"]

	log.Printf("[DEBUG] Swift container create: %s", name)
	_, err := containers.Create(storageClient, containers.CreateOpts{
		Name: name,
	})
	if err != nil {
		return nil, fmt.Errorf("Error create Swift container: %s", err)
	}

	rs.ID = name
	return rs, nil
}

func resource_openstack_swift_container_destroy(
	s *terraform.ResourceState,
	meta interface{}) error {

	p := meta.(*ResourceProvider)
	storageClient := p.storageClient

	name := s.Attributes["name"]
	log.Printf("[DEBUG] Swift Delete Bucket: %s", name)
	return containers.Delete(storageClient, containers.DeleteOpts{
		Name: name,
	})
}

func resource_openstack_swift_container_refresh(
	s *terraform.ResourceState,
	meta interface{}) (*terraform.ResourceState, error) {

	p := meta.(*ResourceProvider)
	storageClient := p.storageClient

	name := s.Attributes["name"]
	resp, err := containers.Get(storageClient, containers.GetOpts{
		Name: name,
	})
	if err != nil {
		return s, err
	}
	defer resp.Body.Close()
	return s, nil
}

func resource_openstack_swift_container_diff(
	s *terraform.ResourceState,
	c *terraform.ResourceConfig,
	meta interface{}) (*terraform.ResourceDiff, error) {

	b := &diff.ResourceBuilder{
		Attrs: map[string]diff.AttrType{
			"name": diff.AttrTypeCreate,
		},
	}

	return b.Diff(s, c)
}
