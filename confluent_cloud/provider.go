package confluent_cloud

import (
	"context"
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
//

//// func NewGRPCProviderServer(p *Provider) *GRPCProviderServer
//
//func (*GRPCProviderServer) ApplyResourceChange ¶
//func (s *GRPCProviderServer) ApplyResourceChange(ctx context.Context, req *tfprotov5.ApplyResourceChangeRequest) (*tfprotov5.ApplyResourceChangeResponse, error)
//
//// func (*GRPCProviderServer) ConfigureProvider ¶
//func (s *GRPCProviderServer) ConfigureProvider(ctx context.Context, req *tfprotov5.ConfigureProviderRequest) (*tfprotov5.ConfigureProviderResponse, error)
//
//// func (*GRPCProviderServer) GetProviderSchema ¶
//func (s *GRPCProviderServer) GetProviderSchema(_ context.Context, req *tfprotov5.GetProviderSchemaRequest) (*tfprotov5.GetProviderSchemaResponse, error)
//
//// func (*GRPCProviderServer) ImportResourceState ¶
//func (s *GRPCProviderServer) ImportResourceState(ctx context.Context, req *tfprotov5.ImportResourceStateRequest) (*tfprotov5.ImportResourceStateResponse, error)
//
//// func (*GRPCProviderServer) PlanResourceChange ¶
//func (s *GRPCProviderServer) PlanResourceChange(ctx context.Context, req *tfprotov5.PlanResourceChangeRequest) (*tfprotov5.PlanResourceChangeResponse, error)
//
//// func (*GRPCProviderServer) PrepareProviderConfig ¶
//func (s *GRPCProviderServer) PrepareProviderConfig(_ context.Context, req *tfprotov5.PrepareProviderConfigRequest) (*tfprotov5.PrepareProviderConfigResponse, error)
//
//// func (*GRPCProviderServer) ReadDataSource ¶
//func (s *GRPCProviderServer) ReadDataSource(ctx context.Context, req *tfprotov5.ReadDataSourceRequest) (*tfprotov5.ReadDataSourceResponse, error)
//
//// func (*GRPCProviderServer) ReadResource ¶
//func (s *GRPCProviderServer) ReadResource(ctx context.Context, req *tfprotov5.ReadResourceRequest) (*tfprotov5.ReadResourceResponse, error)
//
//// func (*GRPCProviderServer) StopContext ¶
//func (s *GRPCProviderServer) StopContext(ctx context.Context) context.Context
//
//StopContext derives a new context from the passed in grpc context. It creates a goroutine to wait for the server stop and propagates cancellation to the derived grpc context.
//
//// func (*GRPCProviderServer) StopProvider ¶
//func (s *GRPCProviderServer) StopProvider(_ context.Context, _ *tfprotov5.StopProviderRequest) (*tfprotov5.StopProviderResponse, error)
//
//// func (*GRPCProviderServer) UpgradeResourceState ¶
//func (s *GRPCProviderServer) UpgradeResourceState(ctx context.Context, req *tfprotov5.UpgradeResourceStateRequest) (*tfprotov5.UpgradeResourceStateResponse, error)
//
//// func (*GRPCProviderServer) ValidateDataSourceConfig ¶
//func (s *GRPCProviderServer) ValidateDataSourceConfig(_ context.Context, req *tfprotov5.ValidateDataSourceConfigRequest) (*tfprotov5.ValidateDataSourceConfigResponse, error)
//
//// func (*GRPCProviderServer) ValidateResourceTypeConfig ¶
//func (s *GRPCProviderServer) ValidateResourceTypeConfig(_ context.Context, req *tfprotov5.ValidateResourceTypeConfigRequest) (*tfprotov5.ValidateResourceTypeConfigResponse, error)
//
//
//
//func GRPCProviderFunc() tfprotov5.ProviderServer {
//	return tfprotov5.ProviderServer{
//		schema{},
//	}
//}

func Provider() *schema.Provider {
	// fmt.Println("Into Provider()")

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
			"truemark-confluent-cloud_environment_create": ResourceEnvironment(),
			// 	"confluentcloud_kafka_cluster":   kafkaClusterResource(),
			// 	"confluentcloud_api_key":         apiKeyResource(),
			// 	"confluentcloud_environment":     environmentResource(),
			// 	"confluentcloud_schema_registry": schemaRegistryResource(),
			// 	"confluentcloud_service_account": serviceAccountResource(),
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
