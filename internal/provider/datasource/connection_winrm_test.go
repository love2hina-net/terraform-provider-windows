package datasource_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/love2hina-net/terraform-provider-windows/internal/provider"
)

func TestAccConnectionWinRM(t *testing.T) {
	const config = `
data "windows_connection_winrm" "test" {
	hostname = "example"
	username = "test-user"
	password = "test-password"
}
`

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { provider.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: config,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.windows_connection_winrm.test", "hostname", "example"),
					resource.TestCheckResourceAttr("data.windows_connection_winrm.test", "username", "test-user"),
					resource.TestCheckResourceAttr("data.windows_connection_winrm.test", "password", "test-password"),
				),
			},
		},
	})
}
