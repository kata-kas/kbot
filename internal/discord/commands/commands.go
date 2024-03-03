package commands

import (
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
)

var commandStructure = []*discordgo.ApplicationCommand{
	{
		Name:        "ping",
		Description: "Ping Klub",
	},
	{
		Name:        "show-message",
		Description: "Display Bot Message",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "message",
				Description: "Message to display",
				Required:    true,
				Choices: []*discordgo.ApplicationCommandOptionChoice{
					{
						Name:  "Unsubscribed",
						Value: "unsubscribed-message",
					},
					{
						Name:  "Denied",
						Value: "denied-message",
					},
					{
						Name:  "Welcome",
						Value: "welcome-message",
					},
					{
						Name:  "Launch",
						Value: "launch-message",
					},
				},
			},
		},
	},
	{
		Name:        "remove-unsubs",
		Description: "Remove unsubscribed users",
	},
	{
		Name:        "give-vip-role",
		Description: "Give VIP role to Kajabi Klub Full Subs",
	},
	{
		Name:        "test-on-join",
		Description: "Test on join",
	},
}

var commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
	"ping": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "pong",
			},
		})
	},
	"show-message":  ShowMessage,
	"remove-unsubs": RemoveUnsubs,
	"give-vip-role": UpgradeUsers,
	"test-on-join":  TestOnJoin,
}

func RegisterCommands(bot *discordgo.Session) {
	bot.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})
	for _, v := range commandStructure {
		_, err := bot.ApplicationCommandCreate(bot.State.User.ID, os.Getenv("KLUB_GUILD_ID"), v)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
	}
}
