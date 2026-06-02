package main

import (
	"context"
	"fmt"
	"os"

	"github.com/untcha/go-adaptivecards/adaptivecards/card"
	m "github.com/untcha/go-adaptivecards/adaptivecards/core/model"
	"github.com/untcha/go-adaptivecards/adaptivecards/elements"
	"github.com/untcha/go-adaptivecards/adaptivecards/webhook"
)

func main() {
	// Webhook URL
	url := os.Getenv("TEAMS_TEST_WORKFLOW_URL")

	title := elements.NewTextBlock("Full-width card").WithWeight(m.WeightBolder)
	body := elements.NewTextBlock("This card spans the full message column in Teams.").WithWrap(true)

	// SetFullWidth emits the Teams host extension "msteams":{"width":"Full"}.
	// It is ignored by non-Teams renderers and survives the validated
	// webhook.PostToWorkflowRaw path (no post-marshal JSON injection needed).
	card, err := card.NewCard().
		AddElement(title).
		AddElement(body).
		SetFullWidth(true).
		Build()

	if err != nil {
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
