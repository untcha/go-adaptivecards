package main

import (
	"context"
	"fmt"
	"os"

	"github.com/untcha/go-adaptivecards/adaptivecards/actions"
	"github.com/untcha/go-adaptivecards/adaptivecards/card"
	"github.com/untcha/go-adaptivecards/adaptivecards/containers"
	"github.com/untcha/go-adaptivecards/adaptivecards/core/model"
	"github.com/untcha/go-adaptivecards/adaptivecards/elements"
	"github.com/untcha/go-adaptivecards/adaptivecards/webhook"
)

func main() {
	// Webhook URL
	url := os.Getenv("TEAMS_TEST_WORKFLOW_URL")

	// Create a simple table with employee data
	table := containers.NewTable().
		AddColumnWithWidth(2).                      // Name column - relative width 2
		AddColumnWithWidth(1).                      // Department - relative width 1
		AddColumnWithWidth("80px").                 // Status - fixed 80px width
		AddTextRow("Name", "Department", "Status"). // Header row
		AddTextRow("Alice Johnson", "Engineering", "Active").
		AddTextRow("Bob Smith", "Marketing", "Active").
		AddTextRow("Carol Davis", "Sales", "On Leave").
		WithFirstRowAsHeader(true).
		WithShowGridLines(true).
		GridStyleEmphasis().
		CellsAlignLeft().
		WithID("employee-table")

	// Create action for the card
	action, err := actions.NewActionOpenURL("View Full Directory", "https://company.com/directory")
	if err != nil {
		panic(err)
	}

	// Create the card with the table
	card := card.NewCard().
		AddTextBlock(elements.NewTextBlock("Employee Directory").
			WithSize(model.SizeExtraLarge).
			WithWeight(model.WeightBolder).
			WithColor(model.ColorAccent)).
		AddTextBlock(elements.NewTextBlock("Current team members and their status").
			WithColor(model.ColorDark).
			WithSpacing(model.SpacingSmall)).
		AddTable(table).
		AddAction(action).
		SetVerticalContentAlignment(model.VContentAlignTop)

	finalCard, err := card.Build()
	if err != nil {
		fmt.Printf("Card build error: %v\n", err)
		return
	}

	// Validate the card
	if err := finalCard.Validate(); err != nil {
		fmt.Printf("Card validation error: %v\n", err)
		return
	}

	// Marshal to JSON
	json, err := finalCard.MarshalJSON()
	if err != nil {
		fmt.Printf("JSON marshal error: %v\n", err)
		return
	} else {
		fmt.Println("Generated Adaptive Card JSON:")
		fmt.Println()
		fmt.Println(string(json))
	}

	// Post to Teams webhook if URL is provided
	if url != "" {
		fmt.Println("\nPosting to Teams webhook...")
		if err := webhook.PostToWorkflowRaw(context.Background(), url, finalCard); err != nil {
			fmt.Printf("Webhook error: %v\n", err)
		} else {
			fmt.Println("Successfully posted to Teams!")
		}
	}
}
