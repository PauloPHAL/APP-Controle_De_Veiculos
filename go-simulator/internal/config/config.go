package config

import "os"

var cfg *config = nil

type database struct {
	host string
}

type kafka struct {
	broker         string
	topicFreight   string
	topicSimulator string
	topicRoute     string
}

type config struct {
	database *database
	kafka    *kafka
}

func newConfig() error {
	cfg = &config{
		database: &database{
			host: getEnv("DATABASE_HOST", "mongodb://admin:admin@localhost:27017/routes?authSource=admin"),
		},
		kafka: &kafka{
			broker:         getEnv("KAFKA_BROKER", "localhost:9092"),
			topicFreight:   getEnv("TOPIC_FREIGHT_WRITER", "freight"),
			topicSimulator: getEnv("TOPIC_SIMULATOR_WRITER", "simulator"),
			topicRoute:     getEnv("TOPIC_ROUTE_READER", "route"),
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
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

//-------------------------------------------------------------

func (c *config) DatabaseHost() string {
	return c.database.host
}

func (c *config) KafkaBroker() string {
	return c.kafka.broker
}

func (c *config) KafkaTopic__WRITER__Freight() string {
	return c.kafka.topicFreight
}

func (c *config) KafkaTopic__WRITER__Simulator() string {
	return c.kafka.topicSimulator
}

func (c *config) KafkaTopic__READER__Route() string {
	return c.kafka.topicRoute
}
