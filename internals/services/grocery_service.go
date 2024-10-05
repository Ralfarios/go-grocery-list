package services

import (
	"github.com/Ralfarios/go-grocery-list/internals/domain"
	"github.com/Ralfarios/go-grocery-list/internals/ports"
)

type GroceryService struct {
	groceryRepository ports.GroceryRepository
}

var _ ports.GroceryService = (*GroceryService)(nil)

func NewGroceryService(repo ports.GroceryRepository) *GroceryService {
	return &GroceryService{groceryRepository: repo}
}

func (service *GroceryService) AddGrocery(description string) (*domain.Grocery, error) {
	grocery, err := service.groceryRepository.AddGrocery(description, "todo")

	if err != nil {
		return nil, err
	}

	return grocery, nil

}
