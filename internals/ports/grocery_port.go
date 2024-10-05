package ports

import (
	"github.com/Ralfarios/go-grocery-list/internals/domain"
	"github.com/spf13/cobra"
)

type GroceryService interface {
	AddGrocery(description string) (*domain.Grocery, error)
}

type GroceryRepository interface {
	AddGrocery(description string, status string) (*domain.Grocery, error)
}

type GroceryHandler interface {
	AddGrocery(cmd *cobra.Command, args []string)
}
