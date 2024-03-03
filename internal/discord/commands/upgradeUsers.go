package commands

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/klubbot/klubbot/internal/scheduler"
)

func UpgradeUsers(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Aye, aye, Captain! 🚀",
		},
	})

	upgradedUsersJob, err := scheduler.NewGiveVipRoleJob(s)
	if err != nil {
		log.Printf("Error running give VIP role job: %v", err)
		return
	}

	upgradedUsers, downgradedUsers, err := upgradedUsersJob.Run(os.Getenv("VIP_SUBS_URL"))
	if err != nil {
		log.Printf("Error running give VIP role job: %v", err)
		return
	}

	log.Printf("Upgraded users: %v", upgradedUsers)
	log.Printf("Downgraded users: %v", downgradedUsers)

	sortedUpgradedUsers := sort.StringSlice(upgradedUsers)
	sortedDowngradedUsers := sort.StringSlice(downgradedUsers)

	s.ChannelMessageSendEmbed(os.Getenv("KLUB_LOG_CHANNEL_ID"), &discordgo.MessageEmbed{
		Title:       "Finished giving VIP roles",
		Description: fmt.Sprintf("[LIVE] Number of new VIPs %d \n  Users: %s \n", len(sortedUpgradedUsers), strings.Join(sortedUpgradedUsers, "\n")),
		Color:       0x00ff00,
	})
	s.ChannelMessageSendEmbed(os.Getenv("KLUB_LOG_CHANNEL_ID"), &discordgo.MessageEmbed{
		Title:       "Finished revoking VIP roles",
		Description: fmt.Sprintf("[LIVE] Number of downgraded VIPs %d \n  Users: %s \n", len(sortedDowngradedUsers), strings.Join(sortedDowngradedUsers, "\n")),
		Color:       0xff0000,
	})
}
