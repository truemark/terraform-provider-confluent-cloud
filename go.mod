module github.com/truemark/terraform-provider-confluent-cloud

go 1.15

replace github.com/truemark/terraform-provider-confluent-cloud => ./confluent_cloud

require (
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.4.3
	golang.org/x/net v0.0.0-20201224014010-6772e930b67b // indirect
)
