package config

// Consumer holds the configuration values for the Kafka consumer.
type Consumer struct {
	Brokers []string `env:"KAFKA_BROKERS" default:"[192.168.178.17:9092]"`
	Topic   string   `env:"KAFKA_STORE_MATCH_MATCH" default:"store.match.match"`
	GroupID string   `env:"KAFKA_CONSUMER_MATCH_GROUP_ID" default:"consumer-match-group"`
}
