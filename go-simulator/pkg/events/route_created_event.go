package events

import "github.com/PauloPHAL/APP-Controle_De_Veiculos/go-simulator/pkg/collection"

type RouteCreatedEvent struct {
	EventName  string                  `json:"event"`
	RouteID    string                  `json:"id"`
	Distance   int                     `json:"distance"`
	Directions []collection.Directions `json:"directions"`
}

func NewRouteCreatedEvent(routeID string, distance int, directions []collection.Directions) *RouteCreatedEvent {
	return &RouteCreatedEvent{
		EventName:  "RouteCreated",
		RouteID:    routeID,
		Distance:   distance,
		Directions: directions,
	}
}
