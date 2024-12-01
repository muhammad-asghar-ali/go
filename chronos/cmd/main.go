package main

import (
	"context"
	"log"

	"chronos/slack"
)

func main() {
	bot := slack.SlackClient()

	slack.Commands(bot)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	log.Println("Slack Age Calculator Bot is running...")
	if err := bot.Listen(ctx); err != nil {
		log.Fatalf("Error starting bot: %v", err)
	}
}
