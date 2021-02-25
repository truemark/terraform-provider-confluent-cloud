package confluent_cloud

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/truemark/terraform-provider-confluent-cloud/confluent_cloud/client"
)

// A TerraForm Provider that supportsthe following operations:
//
//    - Environment
//    - Kafka Clusters
//    - Kafka ACLs
//    - Kafka Topics

func Provider() *schema.Provider {
	log.Printf("[INFO] Creating Provider")
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"username": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("CONFLUENT_CLOUD_USERNAME", ""),
			},
			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("CONFLUENT_CLOUD_PASSWORD", ""),
			},
		},
		ConfigureContextFunc: providerConfigure,
		ResourcesMap: map[string]*schema.Resource{
			"truemark-confluent-cloud_environment":   ResourceEnvironment(),
			"truemark-confluent-cloud_kafka_cluster": ResourceKafkaCluster(),
			// "truemark-confluent-cloud_api_key": apiKeyResource(),
			// 	"truemark-confluent-cloud_schema_registry": schemaRegistryResource(),
			// 	"truemark-confluent-cloud_service_account": serviceAccountResource(),
		},
	}
}

func ResourceEnvironment() *schema.Resource {
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

func ResourceKafkaCluster() *schema.Resource {
	return &schema.Resource{
		CreateContext: kafkaClusterCreate,
		ReadContext:   kafkaClusterRead,
		UpdateContext: kafkaClusterUpdate,
		DeleteContext: kafkaClusterDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    false,
				Description: "what am I?",
			},
			"account_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    false,
				Description: "what am I?",
			},
			"storage": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    false,
				Description: "what am I?",
			},
			"network_ingress": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    false,
				Description: "what am I?",
			},
			"network_egress": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    false,
				Description: "what am I?",
			},
			"region": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    false,
				Description: "what am I?",
			},
			"service_provider": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    false,
				Description: "what am I?",
			},
			"availability": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "LOW(single-zone) or HIGH(multi-zone)",
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					v := val.(string)
					if val != "LOW" && val != "HIGH" {
						errs = append(errs, fmt.Errorf("%q must be `LOW` or `HIGH`, got: %s", key, v))
					}
					return
				},
			},
			// "deployment": {
			// 	Type:        schema.TypeList,
			// 	Required:    true,
			// 	ForceNew:    false,
			// 	Description: "what am I?",
			// 	Elem: &schema.Resource{
			// 		Schema: map[string]*schema.Schema{
			// 			"sku": {
			// 				Type:        schema.TypeString,
			// 				Required:    true,
			// 				ForceNew:    false,
			// 				Description: "",
			// 			},
			// 			"account_id": {
			// 				Type:        schema.TypeString,
			// 				Required:    true,
			// 				ForceNew:    false,
			// 				Description: "",
			// 			},
			// 		},
			// 	},
			// },

			"cku": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    false,
				Description: "what am I?",
			},

			"deployment": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Deployment settings.  Currently only `sku` is supported.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sku": {
							Type:     schema.TypeString,
							Required: true,
						},
						"account_id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "what am I?",
						},
					},
				},
			},
		},
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	log.Printf("[INFO] Initializing ConfluentCloud client")

	username := d.Get("username").(string)
	password := d.Get("password").(string)

	log.Printf("Username: %s\n", username)
	log.Printf("Password: %s\n", password)

	var diags diag.Diagnostics
	c := client.NewClient(username, password)
	loginErr := c.Login()
	if loginErr == nil {
		log.Printf("[INFO] Login to Confluent.Cloud Succeeded\n")
		return c, diags
	}

	log.Printf("[INFO] Loging Error Occurred: %s\n", loginErr.Error())
	err := resource.RetryContext(ctx, 30*time.Minute, func() *resource.RetryError {
		err := c.Login()
		if strings.Contains(err.Error(), "Exceeded rate limit") {
			log.Printf("[INFO] ConfluentCloud API rate limit exceeded, retrying.")
			return resource.RetryableError(err)
		} else {
			log.Printf("[INFO] rate limit is still okay...\n")
		}

		return resource.NonRetryableError(err)
	})

	return c, diag.FromErr(err)
}
