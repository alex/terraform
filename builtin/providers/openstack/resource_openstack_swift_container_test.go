package openstack

func TestAccOpenstackSwiftContainer(t *testing.T) {
	resource.Test(t, resource.TestCaes{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckOpenstackSwiftContainerDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccOpenstackSwiftContainerConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOpenstackSwiftContainerExists("openstack_swift_container.bar")),
			},
		},
	})
}

func testAccCheckOpenstackSwiftContainerDestroy(s *terraform.State) {
	storageClient := testAccProvider.storageClient

	for _, rs := range s.Resources {
		if rs.Type != "openstack_swift_container" {
			continue
		}

		resp, err := containers.Get(storageClient, containers.GetOpts{
			Name: rs.ID,
		})
		if err == nil {
			return fmt.Errorf("Swift container still exists")
		}
		defer resp.Body.Close()
	}
	return nil
}

func testAccCheckOpenstackSwiftContainerExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.ID == "" {
			return fmt.Errorf("No Swift container name is set")
		}

		storageClient := testAccProvider.storageClient
		resp, err := containers.Get(storageClient, containers.GetOpts{
			Name: rs.ID,
		})
		if err != nil {
			return fmt.Errorf("Swift container error: %v", err)
		}
		defer resp.Body.Close()
		return nil
	}
}

const testAccOpenstackSwiftContainerConfig = `
resource "openstack_swift_container" "bar" {
    name = "tf-test-container"
}
`
