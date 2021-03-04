package confluent_cloud

import (
	"context"
	"fmt"
	"strconv"

	clientapi "github.com/cgroschupp/go-client-confluent-cloud/confluentcloud"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ServiceAccountCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*clientapi.Client)

	name := d.Get("name").(string)
	description := d.Get("description").(string)

	req := clientapi.ServiceAccountCreateRequest{
		Name:        name,
		Description: description,
	}

	serviceAccount, err := c.CreateServiceAccount(&req)
	if err == nil {
		d.SetId(fmt.Sprintf("%d", serviceAccount.ID))

		err = d.Set("name", serviceAccount.Name)
		if err != nil {
			return diag.FromErr(err)
		}

		err = d.Set("description", serviceAccount.Description)
		if err != nil {
			return diag.FromErr(err)
		}
	} else {
		// TODO: log.Printf("[ERROR] Could not create Service Account: %s", err)
	}

	return diag.FromErr(err)
}

func ServiceAccountDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*clientapi.Client)
	ID, err := strconv.Atoi(d.Id())
	if err != nil {
		// TODO: log.Printf("[ERROR] Could not parse Service Account ID %s to int", d.Id())
		return diag.FromErr(err)
	}

	err = c.DeleteServiceAccount(ID)
	if err != nil {
		// TODO: log.Printf("[ERROR] Service Account can not be deleted: %d", ID)
		return diag.FromErr(err)
	}

	// TODO: log.Printf("[INFO] Service Account deleted: %d", ID)

	return nil
}

func ServiceAccountRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}
