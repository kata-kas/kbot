package discord

import (
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/klubbot/klubbot/internal/discord/commands"
	"github.com/klubbot/klubbot/internal/discord/handlers"
)

func InitializeBot() (*discordgo.Session, error) {
	token := os.Getenv("DISCORD_TOKEN")
	bot, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}

	bot.AddHandler(func(s *discordgo.Session, _ *discordgo.Ready) {
		log.Printf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
	})

	bot.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsAll)

	err = bot.Open()
	if err != nil {
		return nil, err
	}

	commands.RegisterCommands(bot)
	handlers.RegisterHandlers(bot)

	return bot, nil
}
