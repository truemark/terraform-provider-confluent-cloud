terraform {
  required_providers {
    truemark-confluent-cloud = {
      source = "truemark.io/terraform/truemark-confluent-cloud"
      version = "1.0.0"
    }
  }
}

provider "truemark-confluent-cloud" {
  username = "briancabbott@gmail.com"
  password = "Blu3Bl00p8480!"
}

resource "truemark-confluent-cloud_environment" "mything" {
   name = "myenv190"
}

resource "truemark-confluent-cloud_kafka_cluster" "test" {
  name              = "kafka-cluster-test"
  account_id        = truemark-confluent-cloud_environment.mything.id
  storage           = 5000
  network_ingress   = 100
  network_egress    = 100
  region            = "eu-west-1"
  service_provider  = "aws"
  availability      = "LOW"
  deployment {
    account_id = truemark-confluent-cloud_environment.mything.id
    sku = "BASIC"
  }
  cku               = 4
}