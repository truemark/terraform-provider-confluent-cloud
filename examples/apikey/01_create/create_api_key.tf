terraform {
  required_providers {
    truemark_confluent-cloud = {
      source = "truemark.io/terraform/truemark-confluent-cloud"
      version = "1.0.0"
    }
  }
}

resource "truemark_confluentcloud_api_key" {
	cluster_id = ""
	logical_clusters = [] 
	user_id = 124
	environment_id = ""
	description = ""
	key = ""
	secret = ""
}