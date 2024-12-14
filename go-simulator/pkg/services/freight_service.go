package services

import "math"

type FreightService struct{}

func NewFreightService() *FreightService {
	return &FreightService{}
}

func (f *FreightService) CalculateFreightPrice(distance int) float64 {
	return math.Floor((float64(distance)*0.25+0.47)*100) / 100
}
