package collection

type Route struct {
	ID           string       `bson:"_id" json:"id"`
	Distance     int          `bson:"distance" json:"distance"`
	Directions   []Directions `bson:"directions" json:"directions"`
	FreightPrice float64      `bson:"freight_price" json:"freight_price"`
}

func NewRoute(id string, distance int, directions []Directions) *Route {
	return &Route{
		ID:         id,
		Distance:   distance,
		Directions: directions,
	}
}
