module github.com/truemark/terraform-provider-confluent-cloud

go 1.15

replace (
	github.com/truemark/terraform-provider-confluent-cloud/confluent_cloud => ./confluent_cloud
	github.com/truemark/terraform-provider-confluent-cloud/confluent_cloud/client => ./confluent_cloud/client
)

require (
	// github.com/confluentinc/confluent-kafka-go v1.5.2 // indirect
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.4.3
	github.com/truemark/terraform-provider-confluent-cloud/confluent_cloud v1.0.0
)
