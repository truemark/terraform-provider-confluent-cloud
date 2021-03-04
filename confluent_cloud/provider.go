package confluent_cloud

import (
	"context"
	"strings"
	"time"

	"github.com/cgroschupp/go-client-confluent-cloud/confluentcloud"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// A TerraForm Provider that supportsthe following operations:
//
//    - Environment
//    - Kafka Clusters
//    - Kafka ACLs
//    - Kafka Topics

func Provider() *schema.Provider {
	// TODO: log.Printf("[INFO] Creating Provider")
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"username": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("TRUEMARK_CONFLUENTCLOUD_USERNAME", ""),
			},
			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("TRUEMARK_CONFLUENTCLOUD_PASSWORD", ""),
			},
		},
		ConfigureContextFunc: providerConfigure,
		ResourcesMap: map[string]*schema.Resource{
			"truemark-confluent-cloud_environment":     resourceEnvironment(),
			"truemark-confluent-cloud_kafka_cluster":   resourceKafkaCluster(),
			"truemark-confluent-cloud_api_key":         resourceAPIKey(),
			"truemark-confluent-cloud_schema_registry": resourceSchemaRegistry(),
			"truemark-confluent-cloud_service_account": resourceServiceAccount(),
		},
	}
}

func resourceEnvironment() *schema.Resource {
	return &schema.Resource{
		CreateContext: EnvironmentCreate,
		ReadContext:   EnvironmentRead,
		UpdateContext: EnvironmentUpdate,
		DeleteContext: EnvironmentDelete,
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

func resourceKafkaCluster() *schema.Resource {
	return &schema.Resource{
		CreateContext: ClusterCreate,
		ReadContext:   ClusterRead,
		DeleteContext: ClusterDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the cluster",
			},
			"environment_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Environment ID",
			},
			"bootstrap_servers": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"service_provider": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "AWS / GCP",
			},
			"region": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				Description:  "where",
				ValidateFunc: ValidateKafkaClusterRegion,
			},
			"availability": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				Description:  "LOW(single-zone) or HIGH(multi-zone)",
				ValidateFunc: ValidateAvailability,
			},
			"storage": {
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
				Description: "Storage limit(GB)",
			},
			"network_ingress": {
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
				Description: "Network ingress limit(MBps)",
			},
			"network_egress": {
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
				Description: "Network egress limit(MBps)",
			},
			"deployment": {
				Type:        schema.TypeSet,
				Optional:    true,
				ForceNew:    true,
				Description: "Deployment settings.  Currently only `sku` is supported.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sku": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"cku": {
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
				Description: "cku",
			},
		},
	}
}

func resourceAPIKey() *schema.Resource {
	return &schema.Resource{
		CreateContext: APIKeyCreate,
		ReadContext:   APIKeyRead,
		// UpdateContext: APIKeyUpdate,
		DeleteContext: APIKeyDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"cluster_id": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "",
			},
			"logical_clusters": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional:    true,
				ForceNew:    true,
				Description: "Logical Cluster ID List to create API Key",
			},
			"user_id": {
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
				Description: "User ID",
			},
			"environment_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Environment ID",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Description",
			},
			"key": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"secret": {
				Type:      schema.TypeString,
				Computed:  true,
				Sensitive: true,
			},
		},
	}
}

func resourceSchemaRegistry() *schema.Resource {
	return &schema.Resource{
		CreateContext: SchemaRegistryCreate,
		ReadContext:   SchemaRegistryRead,
		DeleteContext: SchemaRegistryDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"environment_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Environment ID",
			},
			"region": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "where",
			},
			"service_provider": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Cloud provider",
			},
			"endpoint": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

// No Read or Update for Service-Account
func resourceServiceAccount() *schema.Resource {
	return &schema.Resource{
		CreateContext: ServiceAccountCreate,
		ReadContext:   ServiceAccountRead,
		// UpdateContext: ServiceAccountUpdate,
		DeleteContext: ServiceAccountDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "",
			},
			"description": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Service Account Description",
			},
		},
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	// TODO: log.Printf("[INFO] Initializing ConfluentCloud client")
	username := d.Get("username").(string)
	password := d.Get("password").(string)

	var diags diag.Diagnostics
	c := confluentcloud.NewClient(username, password)
	loginErr := c.Login()
	if loginErr == nil {
		return c, diags
	}
	err := resource.RetryContext(ctx, 30*time.Minute, func() *resource.RetryError {
		err := c.Login()
		if strings.Contains(err.Error(), "Exceeded rate limit") {
			// TODO: log.Printf("[INFO] ConfluentCloud API rate limit exceeded, retrying.")
			return resource.RetryableError(err)
		}
		return resource.NonRetryableError(err)
	})
	return c, diag.FromErr(err)
}
