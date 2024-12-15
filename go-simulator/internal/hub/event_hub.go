package hub

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/PauloPHAL/APP-Controle_De_Veiculos/go-simulator/internal/handlers"
	"github.com/PauloPHAL/APP-Controle_De_Veiculos/go-simulator/pkg/events"
	"github.com/PauloPHAL/APP-Controle_De_Veiculos/go-simulator/pkg/services"
	"github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/mongo"
)

type EventHub struct {
	routeService       *services.RouteService
	mongo              *mongo.Client
	channelDrivedMoved chan *events.DriverMovedEvent
	freightWriter      *kafka.Writer
	simulatorWriter    *kafka.Writer
}

func NewEventHub(
	routeService *services.RouteService,
	mongo *mongo.Client,
	channelDrivedMoved chan *events.DriverMovedEvent,
	freightWriter *kafka.Writer,
	simulatorWriter *kafka.Writer,
) *EventHub {
	return &EventHub{
		routeService:       routeService,
		mongo:              mongo,
		channelDrivedMoved: channelDrivedMoved,
		freightWriter:      freightWriter,
		simulatorWriter:    simulatorWriter,
	}
}

func (eh *EventHub) HandleEvent(msg []byte) error {
	var baseEvent struct {
		EventName string `json:"event"`
	}

	err := json.Unmarshal(msg, &baseEvent)
	if err != nil {
		return fmt.Errorf("error unmarshalling event: %w", err)
	}

	switch baseEvent.EventName {
	case "RouteCreated":
		var event events.RouteCreatedEvent

		err := json.Unmarshal(msg, &event)
		if err != nil {
			return fmt.Errorf("error unmarshalling event: %w", err)
		}

		return eh.handleRouteCreated(&event)

	case "DeliveryStarted":
		var event events.DeliveryStartedEvent

		err := json.Unmarshal(msg, &event)
		if err != nil {
			return fmt.Errorf("error unmarshalling event: %w", err)
		}

		return eh.handleDeliveryStarted(&event)

	default:
		return errors.New("event not found")
	}
}

func (eh *EventHub) handleRouteCreated(event *events.RouteCreatedEvent) error {
	freightCalculatedEvent, err := handlers.RouteCreated(event, eh.routeService)
	if err != nil {
		return err
	}

	value, err := json.Marshal(freightCalculatedEvent)
	if err != nil {
		return err
	}

	err = eh.freightWriter.WriteMessages(
		context.Background(),
		kafka.Message{
			Key:   []byte(freightCalculatedEvent.RouteID),
			Value: value,
		},
	)
	if err != nil {
		return err
	}

	return nil
}

func (eh *EventHub) handleDeliveryStarted(event *events.DeliveryStartedEvent) error {
	err := handlers.DeliveryStarted(event, eh.routeService, eh.channelDrivedMoved)
	if err != nil {
		return err
	}
	go eh.sendDirections()

	return nil
}

func (eh *EventHub) sendDirections() {
	for {
		select {
		case movedEvent := <-eh.channelDrivedMoved:
			value, err := json.Marshal(movedEvent)
			if err != nil {
				return
			}
			err = eh.simulatorWriter.WriteMessages(
				context.Background(),
				kafka.Message{
					Key:   []byte(movedEvent.RouterID),
					Value: value,
				},
			)
			if err != nil {
				return
			}
		case <-time.After(500 * time.Millisecond):
			return
		}
	}
}
