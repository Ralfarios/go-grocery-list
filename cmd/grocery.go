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

	rootCmd.AddCommand(AddGroceryCmd, GetAllGroceriesCmd, DeleteGroceryCmd, EditGroceryCmd)
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

var GetAllGroceriesCmd = &cobra.Command{
	Use:   "list",
	Short: "See the grocery list items",
	Long:  `See the all the grocery list items.`,
	Run: func(cmd *cobra.Command, args []string) {
		groceryHandler.GetAllGroceries(cmd, args)
	},
}

var DeleteGroceryCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete item to your grocery list",
	Long: `Delete item to your grocery list by its id.
For example:
	
go-grocery-list delete 1`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]
		groceryHandler.DeleteGrocery(cmd, id)
	},
}

var EditGroceryCmd = &cobra.Command{
	Use:   "update",
	Short: "Update the description of grocery list item.",
	Long: `Update the description of grocery list item based on its ID.
For example:

go-grocery-list update 1 "indomie soto 2 pkgs"
	`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]
		description := args[1]

		groceryHandler.UpdateGrocery(cmd, id, description)
	},
}
