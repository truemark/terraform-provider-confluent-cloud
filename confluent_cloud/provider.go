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
			"create": {

			}, 
			"delete": {

			},
			"list": {

			},
			"update": {

			},
			"use": {

			},
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
			"create": {

			},
			"delete": {

			},
			"describe": {

			},
			"list": {

			},
			"update": {

			},
			"use": {

			},
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
			"create": {

			},
			"delete": {

			},
			"list": {

			},
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
			"consume": {

			},
			"create": {

			},
			"delete": {

			},
			"describe": {

			},
			"list": {

			},
			"produce": {

			},
			"update": {

			},
		},
	}
	return p
}