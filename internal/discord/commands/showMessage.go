package commands

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/klubbot/klubbot/internal/discord/messages"
)

func ShowMessage(s *discordgo.Session, i *discordgo.InteractionCreate) {
	command := i.ApplicationCommandData().Options[0].StringValue()
	switch command {
	case "unsubscribed-message":
		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{messages.UnsubscribedMessage()},
			},
		})
		if err != nil {
			log.Printf("Error sending unsubscribed message: %v", err)
			return
		}
	case "denied-message":
		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{messages.DeniedMessage()},
			},
		})
		if err != nil {
			log.Printf("Error sending kicked message: %v", err)
			return
		}
	case "welcome-message":
		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{messages.WelcomeMessage()},
			},
		})
		if err != nil {
			log.Printf("Error sending welcome message: %v", err)
			return
		}
	case "launch-message":
		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{messages.LaunchMessage()},
			},
		})
		if err != nil {
			log.Printf("Error sending launch message: %v", err)
			return
		}
	}
}
