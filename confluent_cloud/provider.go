package confluent_cloud

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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
	fmt.Println("Into Provider()")

	p := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"environment": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "Confluent.io Environment Configuration.",
				Elem:        resourceEnvironment(),
			},
			"kafka-cluster": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "Confluent.io Kafka-Clusters Configuration.",
				Elem:        resourceKafkaCluster(),
			},
			"kafka-acl": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "Confluent.io Kafka-ACLs Configuration.",
				Elem:        resourceKafkaACL(),
			},
			"kafka-topic": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "Confluent.io Kafa-Topics Configuration.",
				Elem:        resourceKafkaTopic(),
			},
		},
	}

	fmt.Println("Returning from Provider()")

	return p
}

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
func resourceEnvironment() *schema.Provider {
	fmt.Println("Into resourceEnvironment()")

	p := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"create": {},
			"delete": {},
			"list":   {},
			"update": {},
			"use":    {},
		},
	}
	return p
}

// //
// Manage Kafka clusters.
//
// Usage:
//   ccloud kafka cluster [command]
//
// Available Commands:
//   create      Create a Kafka cluster.
//   delete      Delete a Kafka cluster.
//   describe    Describe a Kafka cluster.
//   list        List Kafka clusters.
//   update      Update a Kafka cluster.
//   use         Make the Kafka cluster active for use in other commands.
//
// Global Flags:
//   -h, --help            Show help for this command.
//   -v, --verbose count   Increase verbosity (-v for warn, -vv for info, -vvv for debug, -vvvv for trace).
//
// Use "ccloud kafka cluster [command] --help" for more information about a command.
func resourceKafkaCluster() *schema.Provider {
	fmt.Println("Into resourceKafkaCluster()")

	p := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"create":   {},
			"delete":   {},
			"describe": {},
			"list":     {},
			"update":   {},
			"use":      {},
		},
	}
	return p
}

////
// Manage Kafka ACLs.
//
// Usage:
//   ccloud kafka acl [command]
//
// Available Commands:
//   create      Create a Kafka ACL.
//   delete      Delete a Kafka ACL.
//   list        List Kafka ACLs for a resource.
//
// Global Flags:
//   -h, --help            Show help for this command.
//   -v, --verbose count   Increase verbosity (-v for warn, -vv for info, -vvv for debug, -vvvv for trace).
//
// Use "ccloud kafka acl [command] --help" for more information about a command.
func resourceKafkaACL() *schema.Provider {
	fmt.Println("Into resourceKafkaAcl()")

	p := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"create": {},
			"delete": {},
			"list":   {},
		},
	}
	return p
}

///
// Manage Kafka topics.
//
// Usage:
//   ccloud kafka topic [command]
//
// Available Commands:
//   consume     Consume messages from a Kafka topic.
//   create      Create a Kafka topic.
//   delete      Delete a Kafka topic.
//   describe    Describe a Kafka topic.
//   list        List Kafka topics.
//   produce     Produce messages to a Kafka topic.
//   update      Update a Kafka topic.
//
// Global Flags:
//   -h, --help            Show help for this command.
//   -v, --verbose count   Increase verbosity (-v for warn, -vv for info, -vvv for debug, -vvvv for trace).
//
// Use "ccloud kafka topic [command] --help" for more information about a command.
func resourceKafkaTopic() *schema.Provider {
	fmt.Println("Into resourceKafkaTopic()")

	p := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"consume":  {},
			"create":   {},
			"delete":   {},
			"describe": {},
			"list":     {},
			"produce":  {},
			"update":   {},
		},
	}
	return p
}
