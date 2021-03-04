terraform {
  required_providers {
    truemark-confluent-cloud = {
      source = "truemark.io/terraform/truemark-confluent-cloud"
      version = "1.0.0"
    }
  }
}

resource "truemark_confluentcloud_environment" "truemark_example_environment" {
   name = "truemark_example_environment"
}
