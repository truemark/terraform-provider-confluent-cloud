package confluent_cloud

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	clientapi "github.com/truemark/terraform-provider-confluent-cloud/confluent_cloud/client"
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

func environmentCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*clientapi.Client)
	name := d.Get("name").(string)

	log.Printf("[INFO] Creating Environment %s", name)

	orgID, err := getOrganizationID(c)
	if err != nil {
		return diag.FromErr(err)
	}
	env, err := c.CreateEnvironment(name, orgID)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(env.ID)

	return nil
}

func environmentUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*clientapi.Client)
	newName := d.Get("name").(string)

	log.Printf("[INFO] Updating Environment %s", d.Id())
	orgID, err := getOrganizationID(c)
	if err != nil {
		return diag.FromErr(err)
	}

	env, err := c.UpdateEnvironment(d.Id(), newName, orgID)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(env.ID)
	return nil
}

func environmentRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*clientapi.Client)

	log.Printf("[INFO] Reading Environment %s", d.Id())
	env, err := c.GetEnvironment(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	err = d.Set("name", env.Name)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func environmentDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*clientapi.Client)

	log.Printf("[INFO] Deleting Environment %s", d.Id())
	err := c.DeleteEnvironment(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func getOrganizationID(clientapi *clientapi.Client) (int, error) {
	userData, err := clientapi.Me()
	if err != nil {
		return 0, err
	}

	return userData.Account.OrganizationID, nil
}
