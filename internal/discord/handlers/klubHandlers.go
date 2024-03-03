package handlers

import (
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/klubbot/klubbot/internal/discord/messages"
	"github.com/klubbot/klubbot/internal/kajabi"
)

func OnKlubJoin(s *discordgo.Session, m *discordgo.GuildMemberAdd) {
	if m.GuildID != os.Getenv("KLUB_GUILD_ID") {
		return
	}

	fmt.Printf("New member joined: %s\n", m.User.Username)

	//var isUserSubbed bool
	// if m.User.GlobalName == "" {
	// 	isUserSubbed = kajabi.IsUserSubbed(m.User.Username)
	// } else {
	// 	isUserSubbed = kajabi.IsUserSubbed(m.User.GlobalName)
	// }
	isUserSubbed := kajabi.IsUserSubbed("Agnieszka+Majchrowska")

	if !isUserSubbed {
		fmt.Printf("User %s is not subbed\n", m.User.Username)

		// s.GuildMemberDeleteWithReason(m.GuildID, m.User.ID, "Not subbed")

		// DMChan, err := s.UserChannelCreate(m.User.ID)
		// if err != nil {
		// 	fmt.Println("Error creating DM channel:", err)
		// 	return
		// }

		// _, err = s.ChannelMessageSendEmbed(DMChan.ID, messages.DeniedMessage())
		// if err != nil {
		// 	fmt.Println("Error sending DM message:", err)
		// }

		_, err := s.ChannelMessageSendEmbed(os.Getenv("KLUB_LOG_CHANNEL_ID"), messages.ReportMessage("USER DENIED", m.User.Username))
		if err != nil {
			fmt.Println("Error sending report message:", err)
		}
		return
	} else {
		DMChan, err := s.UserChannelCreate(m.User.ID)
		if err != nil {
			fmt.Println("Error creating DM channel:", err)
			return
		}
		_, err = s.ChannelMessageSendEmbed(DMChan.ID, messages.WelcomeMessage())
		if err != nil {
			fmt.Println("Error sending welcome message:", err)
		}
		_, err = s.ChannelMessageSendEmbed(os.Getenv("KLUB_LOG_CHANNEL_ID"), messages.ReportMessage("USER JOINED", m.User.Username))
		if err != nil {
			fmt.Println("Error sending report message:", err)
		}
	}

	fmt.Printf("User %s is subbed\n", m.User.Username)
}
