package confluent_cloud

import (
	"context"
	"fmt"
	"time"

	clientapi "github.com/cgroschupp/go-client-confluent-cloud/confluentcloud"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func APIKeyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*clientapi.Client)

	clusterID := d.Get("cluster_id").(string)
	logicalClusters := d.Get("logical_clusters").([]interface{})
	accountID := d.Get("environment_id").(string)
	userID := d.Get("user_id").(int)
	description := d.Get("description").(string)

	logicalClustersReq := []clientapi.LogicalCluster{}
	if len(clusterID) > 0 {
		logicalClustersReq = append(logicalClustersReq, clientapi.LogicalCluster{ID: clusterID})
	}

	for i := range logicalClusters {
		if clusterID != logicalClusters[i].(string) {
			logicalClustersReq = append(logicalClustersReq, clientapi.LogicalCluster{
				ID: logicalClusters[i].(string),
			})
		}
	}

	req := clientapi.ApiKeyCreateRequest{
		AccountID:       accountID,
		UserID:          userID,
		LogicalClusters: logicalClustersReq,
		Description:     description,
	}

	// TODO: log.Printf("[DEBUG] Creating API key")
	key, err := c.CreateAPIKey(&req)
	if err == nil {
		d.SetId(fmt.Sprintf("%d", key.ID))

		err = d.Set("key", key.Key)
		if err != nil {
			return diag.FromErr(err)
		}

		err = d.Set("secret", key.Secret)
		if err != nil {
			return diag.FromErr(err)
		}

		// TODO: log.Printf("[INFO] Created API Key, waiting for it become usable")
		stateConf := &resource.StateChangeConf{
			Pending:      []string{"Pending"},
			Target:       []string{"Ready"},
			Refresh:      clusterReady(c, clusterID, accountID, key.Key, key.Secret),
			Timeout:      300 * time.Second,
			Delay:        10 * time.Second,
			PollInterval: 5 * time.Second,
			MinTimeout:   20 * time.Second,
		}

		_, err = stateConf.WaitForStateContext(context.Background())
		if err != nil {
			return diag.FromErr(fmt.Errorf("Error waiting for API Key (%s) to be ready: %s", d.Id(), err))
		}
	} else {
		// TODO: log.Printf("[ERROR] Could not create API key: %s", err)
	}

	// TODO: log.Printf("[INFO] API Key Created successfully: %s", err)
	return diag.FromErr(err)
}

func APIKeyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*clientapi.Client)

	clusterID := d.Get("cluster_id").(string)
	logicalClusters := d.Get("logical_clusters").([]interface{})
	accountID := d.Get("environment_id").(string)

	logicalClustersReq := []clientapi.LogicalCluster{}
	if len(clusterID) > 0 {
		logicalClustersReq = append(logicalClustersReq, clientapi.LogicalCluster{ID: clusterID})
	}

	for i := range logicalClusters {
		if clusterID != logicalClusters[i].(string) {
			logicalClustersReq = append(logicalClustersReq, clientapi.LogicalCluster{
				ID: logicalClusters[i].(string),
			})
		}
	}

	id := d.Id()
	// TODO: log.Printf("[INFO] Deleting API key %s in account %s", id, accountID)
	err := c.DeleteAPIKey(id, accountID, logicalClustersReq)

	return diag.FromErr(err)
}
