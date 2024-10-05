package cmd

import (
	handler "github.com/Ralfarios/go-grocery-list/internals/handlers"
	"github.com/Ralfarios/go-grocery-list/internals/ports"
	"github.com/Ralfarios/go-grocery-list/internals/repositories"
	"github.com/Ralfarios/go-grocery-list/internals/services"
	"github.com/spf13/cobra"
)

var groceryService ports.GroceryService

var groceryHandler *handler.GroceryHandler

func init() {
	groceryRepo := repositories.NewGroceryRepository("internals/db/grocery.json")
	groceryService = services.NewGroceryService(groceryRepo)
	groceryHandler = handler.NewGroceryHandler(groceryService)

	rootCmd.AddCommand(AddGroceryCmd)

	AddGroceryCmd.Flags().StringP("description", "d", "", "Describe the item that you want to add to the grocery list")
}

var AddGroceryCmd = &cobra.Command{
	Use:   "add",
	Short: "Adding item to your grocery list",
	Long: `Adding item to your grocery list by the description.
For example:

go-grocery-list add -d "Buy 2 eggs"`,

	Run: func(cmd *cobra.Command, args []string) {
		groceryHandler.AddGrocery(cmd, args)
	},
}
