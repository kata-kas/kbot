package messages

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

func UnsubscribedMessage() *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Title:       "Podziękowanie od Akademii Seeking Greatness",
		Description: "Chcielibyśmy Tobie podziękować za czas spędzony razem w Klubie. Mamy nadzieję, że wykorzystasz naszą wiedzę i doświadczenie, aby osiągnąć wszystkie swoje cele i marzenia!",
		Color:       0xFFFF00, // Yellow color

		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Kontakt z nami:",
				Value:  "Jeśli jest coś, co możemy dla Ciebie zrobić, pozostajemy do Twojej dyspozycji pod adresem email: klub@seekinggreatness.com",
				Inline: false,
			},
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Dziękujemy ponownie i do zobaczenia w przyszłości!\nPozdrawiamy,\nSylwia i Tom\nAkademia Seeking Greatness \n",
		},

		Timestamp: time.Now().Format(time.RFC3339),
	}
}
