terraform {
  required_providers {
    truemark_confluent-cloud = {
      source = "truemark.io/terraform/truemark-confluent-cloud"
      version = "1.0.0"
    }
  }
}

resource "truemark_confluentcloud_schema_registry" "" {
    environment_id = ""
    region = "eu-west-1"
    service_provider = "aws"
    endpoint = ""
}