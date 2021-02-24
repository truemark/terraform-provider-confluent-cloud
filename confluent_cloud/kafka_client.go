package confluent_cloud

type KafkaConfig struct{}

func createKafkaClient(config *KafkaConfig) {
	// Create a new AdminClient.
	// AdminClient can also be instantiated using an existing
	// Producer or Consumer instance, see NewAdminClientFromProducer and
	// NewAdminClientFromConsumer.
	//broker := ""
	//a, err := kafka.NewAdminClient(&kafka.ConfigMap{"bootstrap.servers": broker})
	//if err != nil {
	//	fmt.Printf("Failed to create Admin client: %s\n", err)
	//	os.Exit(1)
	//}
}
