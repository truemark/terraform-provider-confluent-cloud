// Main Provider file. 
package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	provider "github.com/truemark/terraform-provider-confluent-cloud/confluent-cloud"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{ProviderFunc: provider.Provider})
}