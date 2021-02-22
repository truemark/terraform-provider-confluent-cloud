package confluent_cloud

import (
	"context"
	"fmt"
)

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
func resourceEnvironment() *schema.Resource {
	fmt.Println("Into resourceEnvironment()")
	return &schema.Resource{
		CreateContext: environmentCreate,
		ReadContext:   environmentRead,
		UpdateContext: environmentUpdate,
		DeleteContext: environmentDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    false,
				Description: "The name of the environment",
			},
		},
	}
}

func environmentCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
}

func environmentUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

}

func environmentRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
}

func environmentDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
}
