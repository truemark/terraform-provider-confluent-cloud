package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// A TerraForm Provider that supportsthe following operations:
// 
//    - Environment
//    - Kafka Clusters
//    - Kafka ACLs
//    - Kafka Topics
//
func Provider() *schema.Provider {
	Println("Into Provider()")

	p := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"environment": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Debug indicates whether or not Helm is running in Debug mode.",
				DefaultFunc: schema.EnvDefaultFunc("HELM_DEBUG", false),
			}
		}
	}
	return p
}


////
// Supports Confluence Cloud Environment Operations. Text from 'ccloud environment' is as follows:
//    Manage and select ccloud environments.
//
//    Usage:
//      ccloud environment [command]
//
//    Available Commands:
//      create      Create a new Confluent Cloud environment.
//      delete      Delete a Confluent Cloud environment and all its resources.
//      list        List Confluent Cloud environments.
//      update      Update an existing Confluent Cloud environment.
//      use         Switch to the specified Confluent Cloud environment.
//
//    Global Flags:
//      -h, --help            Show help for this command.
//      -v, --verbose count   Increase verbosity (-v for warn, -vv for info, -vvv for debug, -vvvv for trace).
//
//    Use "ccloud environment [command] --help" for more information about a command.
func resourceEnvironment() {
	p := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"create": {

			}, 
			"delete": {

			},
			"list": {

			},
			"update": {

			},
			"use": {

			},
		}
}