package confluent_cloud

import (
	"context"

	clientapi "github.com/cgroschupp/go-client-confluent-cloud/confluentcloud"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

////
// Supports Confluence Cloud Environment Operations.
//
func EnvironmentCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*clientapi.Client)
	name := d.Get("name").(string)

	// TODO: log.Printf("[INFO] Creating Environment %s", name)

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

////
// Updates the name of an existing Confluent.Cloud Environment
//
func EnvironmentUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*clientapi.Client)
	newName := d.Get("name").(string)

	// TODO: log.Printf("[INFO] Updating Environment %s", d.Id())
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

///
// Performs a read operation on the
func EnvironmentRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*clientapi.Client)

	// TODO: log.Printf("[INFO] Reading Environment %s", d.Id())
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

func EnvironmentDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*clientapi.Client)

	// TODO: log.Printf("[INFO] Deleting Environment %s", d.Id())
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
