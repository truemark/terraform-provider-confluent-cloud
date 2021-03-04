terraform {
  required_providers {
    truemark-confluent-cloud = {
      source = "truemark.io/terraform/truemark-confluent-cloud"
      version = "1.0.0"
    }
  }
}

resource "truemark-confluent-cloud_environment" "my_environment" {
   name = "my_environment"
}
