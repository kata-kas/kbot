package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/klubbot/klubbot/internal/discord/handlers"
)

func TestOnJoin(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Aye, aye, Captain! 🚀",
		},
	})

	user := i.Member.User
	handlers.OnKlubJoin(s, &discordgo.GuildMemberAdd{
		Member: &discordgo.Member{
			GuildID: i.GuildID,
			User:    user,
		},
	})
}
