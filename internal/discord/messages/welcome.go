package messages

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

func WelcomeMessage() *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Title:       "Witamy w Klubie Seeking Greatness!",
		Description: "Hej, witaj w Klubie. Cieszę się, że podjęłaś decyzję, aby dołączyć do naszej społeczności. :blush:",
		Color:       0x00FF00, // Green color

		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Pierwsze kroki:",
				Value:  "Proponuję w pierwszej kolejności przejść do zakładki [Klub Discord](https://discord.com/channels/986025971750666270/1141314274527019118). Tam wszystkiego się dowiesz, od czego możesz zacząć :wink:",
				Inline: false,
			},
			{
				Name:   "Przewodnik dla Ciebie:",
				Value:  "Kolejnym Twoim krokiem będzie przejście do [Przewodnika](https://discord.com/channels/986025971750666270/1141314274527019118), który wprowadzi Cię do struktury całego Klubu i przeprowadzi przez wszystkie najważniejsze początkowe kroki, włącznie z najważniejszymi szkoleniami. \n Pozdrawiam serdecznie i życzę miłego dnia. :wink:",
				Inline: false,
			},
		},
		Timestamp: time.Now().Format(time.RFC3339),
	}
}
