package commands

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/klubbot/klubbot/internal/scheduler"
)

func RemoveUnsubs(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Aye, aye, Captain! 🚀",
		},
	})

	unsubCleanerJob, err := scheduler.NewUnsubCleanerJob(s)
	if err != nil {
		log.Printf("Error creating unsub cleaner job: %v", err)
		return
	}

	unsubCleanerJob.Run()
}
