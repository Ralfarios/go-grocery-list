package handler

import (
	"fmt"
	"strconv"
	"time"

	"github.com/Ralfarios/go-grocery-list/internals/ports"
	"github.com/alexeyco/simpletable"
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

	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "ID"},
			{Align: simpletable.AlignCenter, Text: "DESCRIPTION"},
			{Align: simpletable.AlignCenter, Text: "STATUS"},
			{Align: simpletable.AlignCenter, Text: "CREATED AT"},
			{Align: simpletable.AlignCenter, Text: "UPDATED AT"},
		},
	}

	r := []*simpletable.Cell{
		{Align: simpletable.AlignCenter, Text: fmt.Sprintf("%d", grocery.Id)},
		{Align: simpletable.AlignCenter, Text: grocery.Description},
		{Align: simpletable.AlignCenter, Text: grocery.Status},
		{Align: simpletable.AlignCenter, Text: grocery.CreatedAt.Format(time.RFC850)},
		{Align: simpletable.AlignCenter, Text: grocery.UpdatedAt.Format(time.RFC850)},
	}

	table.Body.Cells = append(table.Body.Cells, r)

	table.SetStyle(simpletable.StyleCompactLite)

	fmt.Println(table)
}

func (handler *GroceryHandler) GetAllGroceries(cmd *cobra.Command, args []string) {
	groceries, err := handler.groceryService.GetAllGroceries()
	if err != nil {
		fmt.Println(err)
		return
	}

	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "ID"},
			{Align: simpletable.AlignCenter, Text: "DESCRIPTION"},
			{Align: simpletable.AlignCenter, Text: "STATUS"},
			{Align: simpletable.AlignCenter, Text: "CREATED AT"},
			{Align: simpletable.AlignCenter, Text: "UPDATED AT"},
		},
	}

	for _, row := range groceries {
		r := []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: fmt.Sprintf("%d", row.Id)},
			{Align: simpletable.AlignCenter, Text: row.Description},
			{Align: simpletable.AlignCenter, Text: row.Status},
			{Align: simpletable.AlignCenter, Text: row.CreatedAt.Format(time.RFC850)},
			{Align: simpletable.AlignCenter, Text: row.UpdatedAt.Format(time.RFC850)},
		}

		table.Body.Cells = append(table.Body.Cells, r)
	}

	table.SetStyle(simpletable.StyleCompactLite)

	fmt.Println(table.String())
}

func (handler *GroceryHandler) DeleteGrocery(cmd *cobra.Command, args string) {
	id, err := strconv.Atoi(args)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = handler.groceryService.DeleteGrocery(id)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Grocery item with id %X has been deleted", id)
}

func (handler *GroceryHandler) UpdateGrocery(cmd *cobra.Command, argId string, argDescription string) {
	id, err := strconv.Atoi(argId)
	if err != nil {
		fmt.Println(err)
		return
	}

	grocery, err := handler.groceryService.UpdateGrocery(id, argDescription)

	if err != nil {
		fmt.Println(err)
		return
	}

	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "ID"},
			{Align: simpletable.AlignCenter, Text: "DESCRIPTION"},
			{Align: simpletable.AlignCenter, Text: "STATUS"},
			{Align: simpletable.AlignCenter, Text: "CREATED AT"},
			{Align: simpletable.AlignCenter, Text: "UPDATED AT"},
		},
	}

	r := []*simpletable.Cell{
		{Align: simpletable.AlignCenter, Text: fmt.Sprintf("%d", grocery.Id)},
		{Align: simpletable.AlignCenter, Text: grocery.Description},
		{Align: simpletable.AlignCenter, Text: grocery.Status},
		{Align: simpletable.AlignCenter, Text: grocery.CreatedAt.Format(time.RFC850)},
		{Align: simpletable.AlignCenter, Text: grocery.UpdatedAt.Format(time.RFC850)},
	}

	table.Body.Cells = append(table.Body.Cells, r)

	table.SetStyle(simpletable.StyleCompactLite)

	fmt.Println(table)
}
