package handlers

import (
	"time"

	"github.com/PauloPHAL/APP-Controle_De_Veiculos/go-simulator/pkg/events"
	"github.com/PauloPHAL/APP-Controle_De_Veiculos/go-simulator/pkg/services"
)

func DeliveryStarted(event *events.DeliveryStartedEvent, routeService *services.RouteService, channel chan *events.DriverMovedEvent) error {
	route, err := routeService.GetRoute(event.RouteID)
	if err != nil {
		return err
	}

	driverMovedEvent := events.NewDriverMovedEvent(route.ID, 0, 0)
	go func() {
		for _, direction := range route.Directions {
			driverMovedEvent = events.NewDriverMovedEvent(route.ID, direction.Lat, direction.Lng)
			time.Sleep(time.Second)
			channel <- driverMovedEvent
		}
	}()

	return nil
}
