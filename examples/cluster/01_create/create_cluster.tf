terraform {
  required_providers {
    truemark_confluent-cloud = {
      source = "truemark.io/terraform/truemark-confluent-cloud"
      version = "1.0.0"
    }
  }
}

truemark.io/terraform/truemark-confluent-cloud/1.0.0/darwin_amd64/terraform-provider-truemark-confluent-cloud 

resource "truemark_confluentcloud_kafka_cluster" "test-terraform" {
  name             = "test-terraform"
  service_provider = "aws"
  region           = "eu-west-1" # ADD A FRIGGEN VALIDATOR HERE
  availability     = "LOW"
  environment_id   = "env-n5wmz"
  deployment {
    sku = "BASIC"
  }
  network_egress  = 100
  network_ingress = 100
  storage         = 5000
}
