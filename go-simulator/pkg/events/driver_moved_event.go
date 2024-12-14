package events

type DriverMovedEvent struct {
	EventName string  `json:"event"`
	RouterID  string  `json:"route_id"`
	Lat       float64 `json:"lat"`
	Lng       float64 `json:"lng"`
}

func NewDriverMovedEvent(routeID string, lat float64, lng float64) *DriverMovedEvent {
	return &DriverMovedEvent{
		EventName: "DriverMoved",
		RouterID:  routeID,
		Lat:       lat,
		Lng:       lng,
	}
}
