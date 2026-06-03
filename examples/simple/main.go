package main

import (
	"context"
	"fmt"
	"os"

	"github.com/untcha/go-adaptivecards/adaptivecards/card"
	"github.com/untcha/go-adaptivecards/adaptivecards/elements"
	"github.com/untcha/go-adaptivecards/adaptivecards/webhook"
)

func main() {
	// Webhook URL
	url := os.Getenv("TEAMS_TEST_WORKFLOW_URL")

	textblock := elements.NewTextBlock("Hello, World!")

	card, err := card.NewCard().
		AddElement(textblock).
		Build()

	if err != nil {
		panic(err)
	}

	// Validate runs logical checks and validates against the embedded
	// Adaptive Cards JSON Schema before the card is sent or serialized.
	if err := card.Validate(); err != nil {
		panic(err)
	}

	json, _ := card.MarshalJSON()
	fmt.Println(string(json))

	// Post to Teams webhook if URL is provided
	if url != "" {
		if err := webhook.PostToWorkflowRaw(context.Background(), url, card); err != nil {
			panic(err)
		}
	}
}
