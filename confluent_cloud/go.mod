module github.com/truemark/terraform-provider-confluent-cloud/confluent_cloud

go 1.15

replace (
	github.com/truemark/terraform-provider-confluent-cloud/confluent_cloud/client => ./confluent_cloud/client
)


require (
	github.com/go-resty/resty/v2 v2.5.0
	github.com/truemark/terraform-provider-confluent-cloud/confluent_cloud/client v0.0.0-20210223090147-2c6c285bfc7d 
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.4.2
	golang.org/x/net v0.0.0-20201224014010-6772e930b67b
)
