package ports

import (
	"github.com/Ralfarios/go-grocery-list/internals/domain"
	"github.com/spf13/cobra"
)

type GroceryService interface {
	AddGrocery(description string) (*domain.Grocery, error)
	// GetAllGroceries() ([]domain.Grocery, error)
	// UpdateGrocery(id int, description string) (domain.Grocery, error)
	// DeleteGrocery(id int) error
}

type GroceryRepository interface {
	AddGrocery(description string, status string) (*domain.Grocery, error)
	// GetAllGroceries() ([]domain.Grocery, error)
	// UpdateGrocery(id int, description string, status string, created_at time.Time) (domain.Grocery, error)
	// DeleteGrocery(id int) error
}

type GroceryHandler interface {
	AddGrocery(cmd *cobra.Command, args []string)
}
