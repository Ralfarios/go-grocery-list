package ports

import (
	"github.com/Ralfarios/go-grocery-list/internals/domain"
	"github.com/spf13/cobra"
)

type GroceryService interface {
	AddGrocery(description string) (*domain.Grocery, error)
	GetAllGroceries(status string) ([]domain.Grocery, error)
	DeleteGrocery(id int) error
	UpdateGrocery(id int, description string) (*domain.Grocery, error)
	MarkGrocery(id int, status string) (*domain.Grocery, error)
}

type GroceryRepository interface {
	AddGrocery(description string, status string) (*domain.Grocery, error)
	GetAllGroceries(status string) ([]domain.Grocery, error)
	DeleteGrocery(id int) error
	UpdateGrocery(id int, description string) (*domain.Grocery, error)
	MarkGrocery(id int, description string) (*domain.Grocery, error)
}

type GroceryHandler interface {
	AddGrocery(cmd *cobra.Command, argDescription string)
	GetAllGroceries(cmd *cobra.Command, args []string)
	DeleteGrocery(cmd *cobra.Command, args string)
	UpdateGrocery(cmd *cobra.Command, argId string, argDescription string)
	MarkGrocery(cmd *cobra.Command, argId string, argStatus string)
}
