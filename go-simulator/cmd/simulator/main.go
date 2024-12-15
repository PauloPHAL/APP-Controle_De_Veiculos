package main

import (
	"context"
	"fmt"
	"log"

	"github.com/PauloPHAL/APP-Controle_De_Veiculos/go-simulator/internal/config"
	"github.com/PauloPHAL/APP-Controle_De_Veiculos/go-simulator/internal/hub"
	"github.com/PauloPHAL/APP-Controle_De_Veiculos/go-simulator/pkg/events"
	"github.com/PauloPHAL/APP-Controle_De_Veiculos/go-simulator/pkg/services"
	"github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	config := config.GetConfig()

	mongoConection, err := mongo.Connect(context.Background(), options.Client().ApplyURI(config.DatabaseHost()))
	if err != nil {
		panic(err)
	}

	routeService := services.NewRouteService(mongoConection, services.NewFreightService())

	channelDriverMoved := make(chan *events.DriverMovedEvent)

	freightWriter := &kafka.Writer{
		Addr:     kafka.TCP(config.KafkaBroker()),
		Topic:    config.KafkaTopic__WRITER__Freight(),
		Balancer: &kafka.LeastBytes{},
	}

	simulatorWriter := &kafka.Writer{
		Addr:     kafka.TCP(config.KafkaBroker()),
		Topic:    config.KafkaTopic__WRITER__Simulator(),
		Balancer: &kafka.LeastBytes{},
	}

	routeReader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{config.KafkaBroker()},
		Topic:   config.KafkaTopic__READER__Route(),
		GroupID: config.KafkaGroupID(),
	})

	eventHub := hub.NewEventHub(routeService, mongoConection, channelDriverMoved, freightWriter, simulatorWriter)

	fmt.Println("Simulator is running")
	for {
		m, err := routeReader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("error: %v", err)
			continue
		}

		go func(msg []byte) {
			err = eventHub.HandleEvent(m.Value)
			if err != nil {
				log.Printf("error: %v", err)
			}
		}(m.Value)
	}
}
