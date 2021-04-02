terraform {
  required_providers {
    truemark_confluent-cloud = {
      source = "truemark.io/terraform/truemark-confluent-cloud"
      version = "1.0.0"
    }
  }
}

# Fields:
# 	cluster_id
# 	logical_clusters 
# 	user_id
# 	environment_id
# 	description
# 	key
# 	secret
resource "truemark_confluentcloud_api_key" "mykey" {
	cluster_id = ""
	environment_id   = "env-n5wk6"

	# logical_clusters = [] 
	# user_id = 124
	# description = ""
	# key = ""
	# secret = ""
}