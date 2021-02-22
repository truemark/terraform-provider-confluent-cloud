terraform {
  required_providers {
    truemark_confluent_cloud  = {
      source  = "truemark/confluent-cloud"
      version = "0.0.001"
    }
  }
}

provider "confluent_cloud" {

}

resource "confluent_cloud_environment" "mything" {
   name = "myenv"
}