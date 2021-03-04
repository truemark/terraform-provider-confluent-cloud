package confluent_cloud

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/Shopify/sarama"
	clientapi "github.com/cgroschupp/go-client-confluent-cloud/confluentcloud"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ClusterCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*clientapi.Client)

	name := d.Get("name").(string)
	region := d.Get("region").(string)
	serviceProvider := d.Get("service_provider").(string)
	durability := d.Get("availability").(string)
	accountID := d.Get("environment_id").(string)
	// deployment := d.Get("deployment").(map[string]interface{})
	deployment := d.Get("deployment").(*schema.Set)
	storage := d.Get("storage").(int)
	networkIngress := d.Get("network_ingress").(int)
	networkEgress := d.Get("network_egress").(int)
	cku := d.Get("cku").(int)

	// TODO: log.Printf("[DEBUG] Creating kafka_cluster")

	dep := clientapi.ClusterCreateDeploymentConfig{
		AccountID: accountID,
		Sku:       "BASIC",
	}

	// TODO: log.Printf("LETS EXPLORE THE TerraForm schema.Set object ")
	// set_list := deployment.List()

	// if val, ok := deployment["sku"]; ok {
	// 	dep.Sku = val.(string)
	// } else {
	// 	dep.Sku = "BASIC"
	// }

	req := clientapi.ClusterCreateConfig{
		Name:            name,
		Region:          region,
		ServiceProvider: serviceProvider,
		Storage:         storage,
		AccountID:       accountID,
		Durability:      durability,
		Deployment:      dep,
		NetworkIngress:  networkIngress,
		NetworkEgress:   networkEgress,
		Cku:             cku,
	}

	cluster, err := c.CreateCluster(req)
	if err != nil {
		// TODO: log.Printf("[ERROR] createCluster failed %v, %s", req, err)
		return diag.FromErr(err)
	}
	d.SetId(cluster.ID)
	// TODO: log.Printf("[DEBUG] Created kafka_cluster %s, Endpoint: %s", cluster.ID, cluster.Endpoint)

	err = d.Set("bootstrap_servers", cluster.Endpoint)
	if err != nil {
		return diag.FromErr(err)
	}

	logicalClusters := []clientapi.LogicalCluster{
		clientapi.LogicalCluster{ID: cluster.ID},
	}

	apiKeyReq := clientapi.ApiKeyCreateRequest{
		AccountID:       accountID,
		LogicalClusters: logicalClusters,
		Description:     "terraform-provider-confluentcloud cluster connection bootstrap",
	}

	// TODO: log.Printf("[DEBUG] Creating bootstrap keypair")
	key, err := c.CreateAPIKey(&apiKeyReq)
	if err != nil {
		return diag.FromErr(err)
	}

	stateConf := &resource.StateChangeConf{
		Pending:      []string{"Pending"},
		Target:       []string{"Ready"},
		Refresh:      clusterReady(c, d.Id(), accountID, key.Key, key.Secret),
		Timeout:      300 * time.Second,
		Delay:        3 * time.Second,
		PollInterval: 5 * time.Second,
		MinTimeout:   20 * time.Second,
	}

	// TODO: log.Printf("[DEBUG] Waiting for cluster to become healthy")
	_, err = stateConf.WaitForStateContext(ctx)
	if err != nil {
		return diag.FromErr(fmt.Errorf("Error waiting for cluster (%s) to be ready: %s", d.Id(), err))
	}

	// TODO: log.Printf("[DEBUG] Deleting bootstrap keypair")
	err = c.DeleteAPIKey(fmt.Sprintf("%d", key.ID), accountID, logicalClusters)
	if err != nil {
		// TODO: log.Printf("[ERROR] Unable to delete bootstrap api key %s", err)
	}

	return nil
}

func clusterReady(client *clientapi.Client, clusterID, accountID, username, password string) resource.StateRefreshFunc {
	return func() (result interface{}, s string, err error) {
		cluster, err := client.GetCluster(clusterID, accountID)
		// TODO: log.Printf("[DEBUG] Waiting for Cluster to be UP: current status %s %s:%s", cluster.Status, username, password)
		// TODO: log.Printf("[DEBUG] cluster %v", cluster)

		if err != nil {
			return cluster, "UNKNOWN", err
		}

		// TODO: log.Printf("[DEBUG] Attempting to connect to %s, created %s", cluster.Endpoint, cluster.Deployment.Created)
		if cluster.Status == "UP" {
			if canConnect(cluster.Endpoint, username, password) {
				return cluster, "Ready", nil
			}
		}

		return cluster, "Pending", nil
	}
}

func canConnect(connection, username, password string) bool {
	client, err := kafkaClient(connection, username, password)
	if err != nil {
		// TODO: log.Printf("[ERROR] Could not build client %s", err)
		return false
	}

	err = client.RefreshMetadata()
	if err != nil {
		// TODO: log.Printf("[ERROR] Could not refresh metadata %s", err)
		return false
	}

	// TODO: log.Printf("[INFO] Success! Connected to %s", connection)
	return true
}

func ClusterDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*clientapi.Client)
	accountID := d.Get("environment_id").(string)
	var diags diag.Diagnostics

	if err := c.DeleteCluster(d.Id(), accountID); err != nil {
		return diag.FromErr(err)
	}
	return diags
}

func ClusterRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*clientapi.Client)
	accountID := d.Get("environment_id").(string)

	cluster, err := c.GetCluster(d.Id(), accountID)
	if err == nil {
		err = d.Set("bootstrap_servers", cluster.Endpoint)
	}
	if err == nil {
		err = d.Set("name", cluster.Name)
	}
	if err == nil {
		err = d.Set("region", cluster.Region)
	}
	if err == nil {
		err = d.Set("service_provider", cluster.ServiceProvider)
	}
	if err == nil {
		err = d.Set("availability", cluster.Durability)
	}
	if err == nil {
		// TODO: Change to *schema.Set
		err = d.Set("deployment", map[string]interface{}{"sku": cluster.Deployment.Sku})
	}
	if err == nil {
		err = d.Set("storage", cluster.Storage)
	}
	if err == nil {
		err = d.Set("network_ingress", cluster.NetworkIngress)
	}
	if err == nil {
		err = d.Set("network_egress", cluster.NetworkEgress)
	}
	if err == nil {
		err = d.Set("cku", cluster.Cku)
	}

	return diag.FromErr(err)
}

func kafkaClient(connection, username, password string) (sarama.Client, error) {
	bootstrapServers := strings.Replace(connection, "SASL_SSL://", "", 1)
	// TODO: log.Printf("[INFO] Trying to connect to %s", bootstrapServers)

	cfg := sarama.NewConfig()
	cfg.Net.SASL.Enable = true
	cfg.Net.SASL.User = username
	cfg.Net.SASL.Password = password
	cfg.Net.SASL.Handshake = true
	cfg.Net.TLS.Enable = true

	return sarama.NewClient([]string{bootstrapServers}, cfg)
}
