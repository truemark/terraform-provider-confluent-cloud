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

resource "truemark-confluent-cloud_environment_create" "mything" {
   name = "myenv"
}