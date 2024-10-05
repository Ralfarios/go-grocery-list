package repositories

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/Ralfarios/go-grocery-list/internals/domain"
	"github.com/Ralfarios/go-grocery-list/internals/ports"
)

type GroceryRepository struct {
	filepath string
}

var _ ports.GroceryRepository = (*GroceryRepository)(nil)

func NewGroceryRepository(filepath string) *GroceryRepository {
	return &GroceryRepository{filepath: filepath}
}

func (repository *GroceryRepository) AddGrocery(description string, status string) (*domain.Grocery, error) {
	groceries, err := repository.loadDb()
	if err != nil {
		return nil, err
	}

	id := len(groceries) + 1

	newGrocery, err := domain.NewGrocery(id, description, status, nil, nil)
	if err != nil {
		return nil, err
	}

	if newGrocery == nil {
		return nil, errors.New("failed to create new grocery")
	}

	groceries = append(groceries, *newGrocery)

	if err := repository.saveDb(groceries); err != nil {
		return nil, err
	}

	return newGrocery, nil
}

func (repository *GroceryRepository) GetAllGroceries() ([]domain.Grocery, error) {
	groceries, err := repository.loadDb()

	if err != nil {
		return nil, err
	}

	return groceries, nil
}

func (repository *GroceryRepository) loadDb() ([]domain.Grocery, error) {
	file, err := os.Open(repository.filepath)

	if err != nil {
		if os.IsNotExist(err) {
			return []domain.Grocery{}, nil
		}
		return nil, err
	}

	defer file.Close()

	var groceries []domain.Grocery

	if err := json.NewDecoder(file).Decode(&groceries); err != nil {
		return nil, err
	}

	return groceries, nil
}

func (repository *GroceryRepository) saveDb(groceries []domain.Grocery) error {
	file, err := os.OpenFile(repository.filepath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return nil
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(groceries)
}
