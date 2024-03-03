package messages

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

func DeniedMessage() *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Title:       "Witamy w Klubie Seeking Greatness",
		Description: "Hej, niestety system nie wpuścił Cię do społeczności Discord Klubu. Powodem tego może być to, że albo masz nieaktywną miesięczną subskrypcję, albo Twoje dane się nie zgadzają.",
		Color:       0xFF0000, // Red color

		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Aby wpłynąć na subskrypcję:",
				Value:  "Jeżeli Twoja miesięczna subskrypcja jest aktywna, to warunkiem wpuszczenia Cię do społeczności Discord Klubu jest to, aby Twoja nazwa użytkownika zarówno na platformie Kajabi, jak i wyświetlana nazwa użytkownika na koncie Discord zawierała Twoje pełne imię i nazwisko.",
				Inline: false,
			},
			{
				Name:   "Przykład poprawnego wpisania pełnego imienia i nazwiska:",
				Value:  "Na Kajabi: Joanna Mazurek\nNa Discord: Joanna Mazurek",
				Inline: false,
			},
			{
				Name:   "Aby zmienić nazwę użytkownika na Kajabi, kliknij w link poniżej:",
				Value:  "[Zmień nazwę na Kajabi](https://www.seekinggreatness.com/settings/account)",
				Inline: false,
			},
			{
				Name:   "Nie wiesz jak zmienić nazwę użytkownika?",
				Value:  "[Zobacz film instruktażowy](https://youtu.be/J1Q5QW0hqdc)",
				Inline: false,
			},
			{
				Name:   "Problemy nadal istnieją?",
				Value:  "Jeżeli masz aktywną subskrypcję i Twoja nazwa użytkownika zarówno na Kajabi, jak i na Discord są takie same i zawierają Twoje pełne imię i nazwisko, a system nadal nie chce Cię wpuścić do społeczności Klubu, napisz nam wiadomość na mail: klub@seekinggreatness.com, a my się tym zajmiemy.",
				Inline: false,
			},
		},
		Timestamp: time.Now().Format(time.RFC3339),
	}
}
