package handlers

import (
	"github.com/PauloPHAL/APP-Controle_De_Veiculos/go-simulator/pkg/collection"
	"github.com/PauloPHAL/APP-Controle_De_Veiculos/go-simulator/pkg/events"
	"github.com/PauloPHAL/APP-Controle_De_Veiculos/go-simulator/pkg/services"
)

func RouteCreated(event *events.RouteCreatedEvent, routeService *services.RouteService) (*events.FreightCalculatedEvent, error) {
	route := collection.NewRoute(event.RouteID, event.Distance, event.Directions)

	routeCreated, err := routeService.CreateRoute(route)
	if err != nil {
		return nil, err
	}

	freightCalculatedEvent := events.NewFreightCalculatedEvent(routeCreated.ID, routeCreated.FreightPrice)
	return freightCalculatedEvent, nil
}
