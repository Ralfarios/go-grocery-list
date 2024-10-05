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

	rootCmd.AddCommand(AddGroceryCmd, GetAllGroceriesCmd, DeleteGroceryCmd, EditGroceryCmd, MarkGroceryCmd)
	GetAllGroceriesCmd.Flags().StringP("status", "s", "", "Determine based on what status, default value is empty")
}

var AddGroceryCmd = &cobra.Command{
	Use:   "add",
	Short: "Adding item to your grocery list",
	Long: `Adding item to your grocery list by the description.
For example:

go-grocery-list add "Buy 2 eggs"`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		description := args[0]

		groceryHandler.AddGrocery(cmd, description)
	},
}

var GetAllGroceriesCmd = &cobra.Command{
	Use:   "list",
	Short: "See the grocery list items",
	Long: `See the all the grocery list items. You can filter based on its status.
For example:

go-grocery-list list -s "todo"

There are todo | in-progress | done, if empty, it will show all status.
`,
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

var MarkGroceryCmd = &cobra.Command{
	Use:   "mark",
	Short: "Change the grocery list item status from todo to in-progress, or done, or vice-versa",
	Long: `Change the grocery list item status based on its ID from todo to other status like in-progress, or done, or back to todo.
For example:

go-grocery-list mark in-progress 5`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		status := args[0]
		id := args[1]

		groceryHandler.MarkGrocery(cmd, id, status)
	},
}
