package confluent_cloud

import "fmt"

func ValidateKafkaClusterRegion(val interface{}, key string) (warns []string, errs []error) {
	// 		// 			// We might need a post checker or something. The key is that some of these are only valid
	// 		// 			// Amazon Web Services (AWS)
	// 		// 			// 	af-south-1 (Cape Town)
	// 		// 			// 	ap-east-1 (AWS Hong Kong)
	// 		// 			// 	ap-northeast-1 (Tokyo)
	// 		// 			// 	ap-northeast-2 (Seoul)
	// 		// 			// 	ap-south-1 (Mumbai)
	// 		// 			// 	ap-southeast-1 (Singapore)
	// 		// 			// 	ap-southeast-2 (Sydney)
	// 		// 			// 	ca-central-1 (Canada Central)
	// 		// 			// 	eu-central-1 (Frankfurt)
	// 		// 			// 	eu-north-1 (Stockholm)
	// 		// 			// 	eu-west-1 (Ireland)
	// 		// 			// 	eu-west-2 (London)
	// 		// 			// 	eu-west-3 (Paris)
	// 		// 			// 	me-south-1 (Bahrain)
	// 		// 			// 	sa-east-1 (São Paulo)
	// 		// 			// 	us-east-1 (N. Virginia)
	// 		// 			// 	us-east-2 (Ohio)
	// 		// 			// 	us-west-1 (N. California) (Only supports Single-Zone (SZ) dedicated clusters with PrivateLink/VPC peering/Transit Gateway.)
	// 		// 			// 	us-west-2 (Oregon)

	// 		// 			// Azure (Microsoft Azure)
	// 		// 			// 	australiaeast (New South Wales)
	// 		// 			// 	canadacentral (Canada)
	// 		// 			// 	centralus (Iowa)
	// 		// 			// 	eastus (Virginia)
	// 		// 			// 	eastus2 (Virginia)
	// 		// 			// 	francecentral (France)
	// 		// 			// 	northeurope (Ireland)
	// 		// 			// 	southeastasia (Singapore)
	// 		// 			// 	uksouth (London)
	// 		// 			// 	westus2 (Washington)
	// 		// 			// 	westeurope (Netherlands)

	// 		// 			// GCP (Google Cloud Platform)
	// 		// 			// 	asia-east1 (Taiwan)
	// 		// 			// 	asia-east2 (Hong Kong)
	// 		// 			// 	asia-northeast1 (Tokyo)
	// 		// 			// 	asia-northeast3 (Seoul)
	// 		// 			// 	asia-south1 (Mumbai)
	// 		// 			// 	asia-southeast1 (Singapore)
	// 		// 			// 	asia-southeast2 (Jakarta)
	// 		// 			// 	australia-southeast1 (Sydney)
	// 		// 			// 	europe-north1 (Finland)
	// 		// 			// 	europe-west1 (Belgium)
	// 		// 			// 	europe-west2 (London)
	// 		// 			// 	europe-west3 (Frankfurt)
	// 		// 			// 	europe-west4 (Netherlands)
	// 		// 			// 	europe-west6 (Zurich)
	// 		// 			// 	northamerica-northeast1 (Montreal)
	// 		// 			// 	southamerica-east1 (São Paulo)
	// 		// 			// 	us-central1 (Iowa)
	// 		// 			// 	us-east1 (S. Carolina)
	// 		// 			// 	us-east4 (N. Virginia)
	// 		// 			// 	us-west1 (Oregon)
	// 		// 			// 	us-west2 (Los Angeles)
	// 		// 			// 	us-west4 (Las Vegas)
	return nil, nil
}

func ValidateAvailability(val interface{}, key string) (warns []string, errs []error) {
	v := val.(string)
	if val != "LOW" && val != "HIGH" {
		errs = append(errs, fmt.Errorf("%q must be `LOW` or `HIGH`, got: %s", key, v))
	}
	return
}
