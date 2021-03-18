package main

import (
	"github.com/marcozj/terraform-provider-centrifyvault/centrify"

	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return centrify.Provider()
		},
	})
}
