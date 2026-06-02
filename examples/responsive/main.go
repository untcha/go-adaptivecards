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

	// targetWidth tailors which elements render at each card width. It is a
	// documented Adaptive Cards host feature that is NOT part of the published
	// JSON schema, so the library validates it logically and strips it before
	// strict schema validation.
	//
	// Two things to know when testing this:
	//   - It filters on the CARD's width, not the device. A normal message card
	//     is "standard" width, so "atLeast:wide" never shows. This example calls
	//     SetFullWidth(true) below so the card can reach the "wide" bucket — the
	//     two features are companions.
	//   - Support varies by host: Teams Desktop honors targetWidth; some clients
	//     (e.g. Teams mobile) ignore it and show every element regardless. That's
	//     the expected graceful-degradation fallback for an out-of-schema feature.
	always := elements.NewTextBlock("Always visible").WithWrap(true)
	wideOnly := elements.NewTextBlock("Shown only on wide cards (desktop)").
		WithWrap(true).
		WithTargetWidth(m.TargetWidthAtLeast(m.TargetWidthWide))
	narrowOnly := elements.NewTextBlock("Shown only on narrow cards (mobile)").
		WithWrap(true).
		WithTargetWidth(m.TargetWidthAtMost(m.TargetWidthNarrow))

	c, err := card.NewCard().
		AddElement(always).
		AddElement(wideOnly).
		AddElement(narrowOnly).
		SetFullWidth(true).
		Build()

	if err != nil {
		panic(err)
	}

	json, _ := c.MarshalJSON()
	fmt.Println(string(json))

	// Post to Teams webhook if URL is provided
	if url != "" {
		if err := webhook.PostToWorkflowRaw(context.Background(), url, c); err != nil {
			panic(err)
		}
	}
}
