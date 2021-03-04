module github.com/truemark/terraform-provider-confluent-cloud

go 1.15

replace (
	github.com/truemark/terraform-provider-confluent-cloud/confluent_cloud => ./confluent_cloud
)

require (
	github.com/Shopify/sarama v1.28.0 // indirect
	github.com/cgroschupp/go-client-confluent-cloud v0.0.0-20201105075001-2e15b5846d7e // indirect\
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.4.3
	github.com/truemark/terraform-provider-confluent-cloud/confluent_cloud v1.0.0
)
