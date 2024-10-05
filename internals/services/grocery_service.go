package services

import (
	"errors"

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

func (service *GroceryService) GetAllGroceries() ([]domain.Grocery, error) {
	groceries, err := service.groceryRepository.GetAllGroceries()

	if err != nil {
		return nil, err
	}

	return groceries, nil
}

func (service *GroceryService) DeleteGrocery(id int) error {
	err := service.groceryRepository.DeleteGrocery(id)

	if id < 1 {
		return errors.New("id is invalid")
	}

	if err != nil {
		return err
	}

	return nil
}
