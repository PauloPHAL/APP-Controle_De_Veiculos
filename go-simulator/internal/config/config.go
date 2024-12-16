package config

import (
	"os"
)

var cfg *config = nil

type database struct {
	host string
}

type kafka struct {
	kafkaBroker          string
	kafkaRouteTopic      string
	kafkaFreightTopic    string
	kafkaSimulationTopic string
	kafkaGroupID         string
}

type config struct {
	database *database
	kafka    *kafka
}

func newConfig() error {
	cfg = &config{
		database: &database{
			host: getEnv("DATABASE_HOST", "mongodb://admin:admin@mongo:27017/routes?authSource=admin"),
		},
		kafka: &kafka{
			kafkaBroker:          getEnv("KAFKA_BROKER", "kafka:9092"),
			kafkaRouteTopic:      getEnv("KAFKA_ROUTE_TOPIC", "route"),
			kafkaFreightTopic:    getEnv("KAFKA_FREIGHT_TOPIC", "freight"),
			kafkaSimulationTopic: getEnv("KAFKA_SIMULATION_TOPIC", "simulation"),
			kafkaGroupID:         getEnv("KAFKA_GROUP_ID", "route-group"),
		},
	}
	return nil
}

func GetConfig() *config {
	if cfg == nil {
		if err := newConfig(); err != nil {
			panic(err)
		}
	}
	return cfg
}

func getEnv(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}

//-------------------------------------------------------------

func (c *config) DatabaseHost() string {
	return c.database.host
}

func (c *config) KafkaBroker() string {
	return c.kafka.kafkaBroker
}

func (c *config) KafkaTopic__READER__Route() string {
	return c.kafka.kafkaRouteTopic
}

func (c *config) KafkaTopic__WRITER__Freight() string {
	return c.kafka.kafkaFreightTopic
}

func (c *config) KafkaTopic__WRITER__Simulator() string {
	return c.kafka.kafkaSimulationTopic
}

func (c *config) KafkaGroupID() string {
	return c.kafka.kafkaGroupID
}
