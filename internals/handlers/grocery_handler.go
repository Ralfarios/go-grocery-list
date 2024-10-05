package handler

import (
	"fmt"

	"github.com/Ralfarios/go-grocery-list/internals/ports"
	"github.com/spf13/cobra"
)

type GroceryHandler struct {
	groceryService ports.GroceryService
}

var _ ports.GroceryHandler = (*GroceryHandler)(nil)

func NewGroceryHandler(service ports.GroceryService) *GroceryHandler {
	return &GroceryHandler{groceryService: service}
}

func (handler *GroceryHandler) AddGrocery(cmd *cobra.Command, args []string) {
	description, _ := cmd.Flags().GetString("description")

	grocery, err := handler.groceryService.AddGrocery(description)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(grocery)
}
