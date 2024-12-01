package slack

import (
	"chronos/config"

	"github.com/slack-io/slacker"
)

func SlackClient() *slacker.Slacker {
	cfg := config.LoadConfig()
	return slacker.NewClient(cfg.SlackBotToken, cfg.SlackAppToken)
}

func Commands(bot *slacker.Slacker) {
	bot.AddCommand(AgeCmd())   // @BotName my dob is 2000-03-11
	bot.AddCommand(PongCmd())  // @BotName ping
	bot.AddCommand(HelloCmd()) // @BotName hello
}
