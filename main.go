// Main Provider file. 
package main

import (
	confluentcloud "github.com/truemark/terraform-provider-confluent-cloud/confluent-cloud"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{ProviderFunc: confluentcloud.provider.Provider})
}