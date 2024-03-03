package messages

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

func ReportMessage(action string, username string) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{

		Title:       "Report",
		Description: action,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Username",
				Value:  username,
				Inline: false,
			},
		},
		Color:     0x3498db, // Blue color
		Timestamp: time.Now().Format(time.RFC3339),
	}
}
