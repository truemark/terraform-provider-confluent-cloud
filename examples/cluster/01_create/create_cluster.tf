terraform {
  required_providers {
    truemark-confluent-cloud = {
      source = "truemark.io/terraform/truemark-confluent-cloud"
      version = "1.0.0"
    }
  }
}

resource "truemark-confluent-cloud_kafka_cluster" "test-terraform2" {
  name             = "test-terraform2"
  service_provider = "aws"
  region           = "eu-west-1" # ADD A FRIGGEN VALIDATOR HERE
  availability     = "LOW"
  environment_id   = "env-n5wk6"
  deployment {
    sku = "BASIC"
  }
  network_egress  = 100
  network_ingress = 100
  storage         = 5000
}
