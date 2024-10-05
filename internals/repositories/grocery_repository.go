package repositories

import (
	"encoding/json"
	"errors"
	"os"
	"strings"
	"time"

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

	if len(strings.TrimSpace(description)) < 1 {
		return nil, errors.New("please insert the description")
	}

	var id int

	if len(groceries) < 1 {
		id = len(groceries) + 1
	} else {
		id = groceries[len(groceries)-1].Id + 1
	}

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

func (repository *GroceryRepository) DeleteGrocery(id int) error {
	groceries, err := repository.loadDb()

	if err != nil {
		return err
	}

	newGroceries, isExist := filter(groceries, func(item domain.Grocery) bool {
		return item.Id != id
	})

	if !isExist {
		return errors.New("item not found")
	}

	if err := repository.saveDb(newGroceries); err != nil {
		return err
	}

	return nil
}

func (repository *GroceryRepository) UpdateGrocery(id int, description string) (*domain.Grocery, error) {
	groceries, err := repository.loadDb()

	if err != nil {
		return nil, err
	}

	if len(strings.TrimSpace(description)) < 1 {
		return nil, errors.New("please insert the description")
	}

	currGrocery := find(groceries, func(item domain.Grocery, _ int) bool {
		return item.Id == id
	})

	if currGrocery == nil {
		return nil, errors.New("item not found")
	}

	currGrocery.Description = description
	currGrocery.UpdatedAt = time.Now()

	for idx := range groceries {
		if groceries[idx].Id == id {
			groceries[idx] = *currGrocery
		}
	}

	err = repository.saveDb(groceries)

	if err != nil {
		return nil, err
	}

	return currGrocery, nil
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

func find[T any](values []T, cb func(item T, index int) bool) *T {
	for idx, item := range values {
		if cb(item, idx) {
			return &item
		}

	}
	return nil
}

func filter[T any](values []T, cb func(T) bool) (output []T, isExist bool) {
	isExist = false
	output = []T{}

	for _, item := range values {
		if cb(item) {
			output = append(output, item)
		} else {
			isExist = true
		}
	}

	return
}
