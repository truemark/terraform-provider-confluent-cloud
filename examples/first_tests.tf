
### TrueMark.io Current TerraForm Examples
# https://registry.terraform.io/modules/truemark/provisioner/aws/latest
# https://registry.terraform.io/modules/truemark/newrelic/aws/latest
# https://registry.terraform.io/modules/truemark/s3-iam/aws/latest
# https://registry.terraform.io/modules/truemark/route53-gmail/aws/latest
# https://registry.terraform.io/modules/truemark/certificate-route53/aws/latest
#
terraform {
  required_providers {
    terraform-provider-truemark-confluent-cloud = {
      source  = "truemark/truemark-confluent-cloud/
      terraform-provider-truemark-confluent-cloud
"
      version = "1.0.0"
    }
  }
}


provider "truemark_confluent_cloud " {
    # Our Configuration Settings will go in here... Erik - Grappling a bit here - U/P? 
    #   - Probably not eh? What mechanisms/methods do we want to expose for PM mgmt?
    #   - Vault? :) #HashiFan 
}


resource "kafka_cluster_one" "our_first_cluster" {

}
