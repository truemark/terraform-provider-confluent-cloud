package confluent_cloud

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	clientapi "github.com/truemark/terraform-provider-confluent-cloud/confluent_cloud/client"
)

func kafkaClusterCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("into confluent_kafka_cluster/kafkaClusterCreate()\n")

	// Schema: map[string]*schema.Schema{
	// 	"name"
	// 	"accountId"
	// 	"storage"
	// 	"network_ingress"
	// 	"network_egress"
	// 	"region"
	// 	"serviceProvider"
	// 	"durability"
	// 	"deployment"
	// 		"sku"
	// 		"account_id"
	// 	"cku"
	c := meta.(*clientapi.Client)

	name := d.Get("name").(string)
	accountID := d.Get("account_id").(string)
	storage := d.Get("storage").(int)
	networkIngress := d.Get("network_ingress").(int)
	networkEgress := d.Get("network_egress").(int)
	region := d.Get("region").(string)
	serviceProvider := d.Get("service_provider").(string)
	availability := d.Get("availability").(string)
	cku := d.Get("cku").(int)

	deployment := d.Get("deployment").(interface{})
	log.Printf("dep type: %s\n", reflect.TypeOf(deployment))

	dep := clientapi.ClusterCreateDeploymentConfig{
		AccountID: accountID,
		Sku:       "BASIC",
	}

	// val := deployment["sku"]
	// if val != "" {
	// 	dep.Sku = val.(string)
	// } else {
	// dep.Sku = "BASIC"
	// }
	// dep.Sku = "BASIC"

	// req := clientapi.ClusterCreateConfig{
	// 	Name:            name,
	// 	Region:          region,
	// 	ServiceProvider: serviceProvider,
	// 	Storage:         storage,
	// 	AccountID:       accountID,
	// 	Availability:    availability,
	// 	Deployment:      dep,
	// 	NetworkIngress:  networkIngress,
	// 	NetworkEgress:   networkEgress,
	// 	Cku:             cku,
	// }

	req2 := clientapi.ClusterCreateConfig{
		Name:            name,
		Region:          region,
		ServiceProvider: serviceProvider,
		Storage:         storage,
		AccountID:       accountID,
		Availability:    availability,
		Deployment:      dep,
		NetworkIngress:  networkIngress,
		NetworkEgress:   networkEgress,
		Cku:             cku,
	}
	log.Printf("%s\n", req2)
	cluster, err := c.CreateCluster(req2)
	log.Printf("Cluster-Info: %s\n", cluster)
	if err != nil {
		log.Printf("An error occurred creating the Kafka Cluster. Error was: %s\n", err.Error())
		return diag.FromErr(err)
	}
	d.SetId(cluster.ID)

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
		Description:     "truemark-confluent-cloud cluster connection bootstrap",
	}

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

	_, err = stateConf.WaitForStateContext(ctx)
	if err != nil {
		return diag.FromErr(fmt.Errorf("Error waiting for cluster (%s) to be ready: %s", d.Id(), err))
	}

	err = c.DeleteAPIKey(fmt.Sprintf("%d", key.ID), accountID, logicalClusters)
	if err != nil {
		log.Printf("[ERROR] Unable to delete bootstrap api key %s", err)
	}

	return nil
}

func clusterReady(client *clientapi.Client, clusterID, accountID, username, password string) resource.StateRefreshFunc {
	return func() (result interface{}, s string, err error) {
		cluster, err := client.GetCluster(clusterID, accountID)
		log.Printf("[DEBUG] Waiting for Cluster to be UP: current status %s %s:%s", cluster.Status, username, password)
		log.Printf("[DEBUG] cluster %v", cluster)

		if err != nil {
			return cluster, "UNKNOWN", err
		}

		log.Printf("[DEBUG] Attempting to connect to %s, created %s", cluster.Endpoint, cluster.Deployment.Created)
		if cluster.Status == "UP" {
			if canConnect(cluster.Endpoint, username, password) {
				return cluster, "Ready", nil
			}
		}

		return cluster, "Pending", nil
	}
}

func canConnect(connection, username, password string) bool {
	_, err := kafkaClient(connection, username, password)
	if err != nil {
		log.Printf("[ERROR] Could not build client %s", err)
		return false
	}

	// // err = client.RefreshMetadata()
	// if err != nil {
	// 	log.Printf("[ERROR] Could not refresh metadata %s", err)
	// 	return false
	// }

	log.Printf("[INFO] Success! Connected to %s", connection)
	return true
}

func kafkaClient(connection, username, password string) (*kafka.AdminClient, error) {
	// log.Printf("[INFO] Trying to connect to %s", bootstrapServers)

	// cfg := sarama.NewConfig()
	// cfg.Net.SASL.Enable = true
	// cfg.Net.SASL.User = username
	// cfg.Net.SASL.Password = password
	// cfg.Net.SASL.Handshake = true
	// cfg.Net.TLS.Enable = true
	bootstrapServers := strings.Replace(connection, "SASL_SSL://", "", 1)
	config := &kafka.ConfigMap{
		"bootstrap.servers":       bootstrapServers,
		"broker.version.fallback": "0.10.0.0",
		"api.version.fallback.ms": 0,
		// "sasl.mechanisms":         "PLAIN",
		// "security.protocol":       "SASL_SSL",
		// "sasl.username":           ccloudAPIKey,
		// "sasl.password":           ccloudAPISecret,
	}
	adminClient, err := kafka.NewAdminClient(config)
	if err != nil {
		log.Printf("Error occurred instantiating client. %s\n", err.Error())
		return nil, err
	}

	return adminClient, nil
}

func kafkaClusterRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

func kafkaClusterUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

func kafkaClusterDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}
