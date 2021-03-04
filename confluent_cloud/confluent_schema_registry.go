package confluent_cloud

import (
	"context"

	clientapi "github.com/cgroschupp/go-client-confluent-cloud/confluentcloud"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func SchemaRegistryCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*clientapi.Client)

	environment := d.Get("environment_id").(string)
	region := d.Get("region").(string)
	serviceProvider := d.Get("service_provider").(string)

	// TODO: log.Printf("[INFO] Creating Schema Registry %s", environment)

	reg, err := c.CreateSchemaRegistry(environment, region, serviceProvider)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(reg.ID)
	err = d.Set("endpoint", reg.Endpoint)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func SchemaRegistryRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*clientapi.Client)

	environment := d.Get("environment_id").(string)
	// TODO: log.Printf("[INFO] Reading Schema Registry %s", environment)

	env, err := c.GetSchemaRegistry(environment)
	if err != nil {
		return diag.FromErr(err)
	}

	err = d.Set("environment_id", environment)
	if err != nil {
		err = d.Set("endpoint", env.Endpoint)
	}

	return diag.FromErr(err)
}

func SchemaRegistryDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}
