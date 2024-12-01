package slack

import (
	"fmt"

	"github.com/slack-io/slacker"
)

func AgeCmd() *slacker.CommandDefinition {
	return &slacker.CommandDefinition{
		Command:     "my dob is <dob>",
		Description: "Calculate age based on date of birth (YYYY-MM-DD)",
		Examples:    []string{"calculate_age 1990-12-15"},
		Handler: func(ctx *slacker.CommandContext) {
			dob := ctx.Request().StringParam("dob", "")

			age, err := calculate(dob)
			if err != nil {
				msg := fmt.Sprintf("Error: %v. Please provide the date in YYYY-MM-DD format.", err)
				ctx.Response().Reply(msg)
				return
			}

			msg := fmt.Sprintf("You are %d years old!", age)
			ctx.Response().Reply(msg)
		},
	}
}

func PongCmd() *slacker.CommandDefinition {
	return &slacker.CommandDefinition{
		Command: "ping",
		Handler: func(ctx *slacker.CommandContext) {
			ctx.Response().Reply("pong")
		},
	}
}

func HelloCmd() *slacker.CommandDefinition {
	return &slacker.CommandDefinition{
		Command: "hello",
		Handler: func(ctx *slacker.CommandContext) {
			ctx.Response().Reply("hi!")
		},
		HideHelp: true,
	}
}
