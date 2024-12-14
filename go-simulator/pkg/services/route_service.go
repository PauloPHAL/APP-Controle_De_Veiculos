package services

import (
	"github.com/PauloPHAL/APP-Controle_De_Veiculos/go-simulator/pkg/collection"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type RouteService struct {
	mongo          *mongo.Client
	freightService *FreightService
}

func NewRouteService(mongo *mongo.Client, freightService *FreightService) *RouteService {
	return &RouteService{
		mongo:          mongo,
		freightService: freightService,
	}
}

func (rs *RouteService) CreateRoute(route *collection.Route) (*collection.Route, error) {
	// Calcula o preço do frete com base na distância da rota
	route.FreightPrice = rs.freightService.CalculateFreightPrice(route.Distance)

	// bson.M é um tipo do pacote BSON (Binary JSON) usado para representar documentos BSON em Go.
	// Ele é um mapa de strings para interfaces vazias, permitindo a criação de documentos BSON dinâmicos.
	update := bson.M{
		"$set": bson.M{
			"distance":      route.Distance,     // Atualiza o campo "distance" com o valor de route.Distance
			"directions":    route.Directions,   // Atualiza o campo "directions" com o valor de route.Directions
			"freight_price": route.FreightPrice, // Atualiza o campo "freightPrice" com o valor de route.FreightPrice
		},
	}

	filter := bson.M{"_id": route.ID}           // Filtro para encontrar o documento com o ID especificado em route.ID
	options := options.Update().SetUpsert(true) // Opção para criar um novo documento se não existir (upsert)

	_, err := rs.mongo.Database("routes").Collection("routes").UpdateOne(nil, filter, update, options)
	if err != nil {
		return nil, err // Retorna um erro se a operação de atualização falhar
	}

	// Se a operação de atualização for bem-sucedida, retorna um objeto Route vazio (pode ser modificado para retornar o objeto atualizado)
	return route, nil
}

func (rs *RouteService) GetRoute(id string) (collection.Route, error) {
	var route collection.Route

	filter := bson.M{"_id": id} // Filtro para encontrar o documento com o ID especificado em id

	err := rs.mongo.Database("routes").Collection("routes").FindOne(nil, filter).Decode(&route)
	if err != nil {
		return collection.Route{}, err // Retorna um erro se a operação de busca falhar
	}

	return route, nil // Retorna o objeto Route encontrado
}
