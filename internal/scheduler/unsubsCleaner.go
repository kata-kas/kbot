package scheduler

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/go-co-op/gocron"
	"github.com/klubbot/klubbot/internal/kajabi"
	"github.com/klubbot/klubbot/internal/normalization"
)

type UnsubCleanerJob struct {
	*gocron.Job
	s *discordgo.Session
}

func NewUnsubCleanerJob(s *discordgo.Session) (*UnsubCleanerJob, error) {
	return &UnsubCleanerJob{&gocron.Job{}, s}, nil
}

func (j UnsubCleanerJob) Run() ([]discordgo.User, error) {
	start := time.Now()
	kajabiUsers := kajabi.ScrapeUsers(os.Getenv("SUBS_URL"))

	discordUsers, err := j.s.GuildMembers(os.Getenv("KLUB_GUILD_ID"), "", 1000)
	if err != nil {
		fmt.Println("Error loading members:", err)
		return nil, err
	}
	fmt.Printf("Discord users: %v\n", len(discordUsers))

	fmt.Println("Scraping users")

	fmt.Printf("Kajabi users: %v\n", len(kajabiUsers))
	var unsubscribedUsers []discordgo.User
	for _, discordUser := range discordUsers {
		cleanedGlobalName := normalization.NormalizeForComparison(discordUser.User.GlobalName)
		cleanedUsername := normalization.NormalizeForComparison(discordUser.User.Username)

		fmt.Printf("Checking %s -- %s \n", cleanedGlobalName, discordUser.User.Username)

		if discordUser.User.Bot == true ||
			containsRole(discordUser.Roles, os.Getenv("KLUB_ADMIN_ROLE_ID")) ||
			containsRole(discordUser.Roles, os.Getenv("KLUB_VIP_ROLE_ID")) {
			continue
		}

		if !contains(cleanedGlobalName, kajabiUsers) &&
			!contains(cleanedUsername, kajabiUsers) {
			unsubscribedUsers = append(unsubscribedUsers, *discordUser.User)
			continue
		}
	}

	userNames := make([]string, len(unsubscribedUsers))
	userNamesDisplayNamesMap := make([]string, len(unsubscribedUsers))
	for i, user := range unsubscribedUsers {
		userNames[i] = user.Username
		userNamesDisplayNamesMap[i] = fmt.Sprintf("%s -- %s", user.Username, user.String())
	}
	// for i, user := range unsubscribedUsers {
	// 	userNames[i] = user.Username
	// 	userNamesDisplayNamesMap[i] = fmt.Sprintf("%s -- %s", user.Username, user.String())

	// 	DMChan, err := j.s.UserChannelCreate(user.ID)
	// 	if err != nil {
	// 		fmt.Println("Error creating DM channel:", err)
	// 		continue
	// 	}

	// 	_, err = j.s.ChannelMessageSendEmbed(DMChan.ID, messages.UnsubscribedMessage())
	// 	if err != nil {
	// 		fmt.Println("Error sending DM message:", err)
	// 		continue
	// 	}

	// 	err = j.s.GuildMemberDeleteWithReason(os.Getenv("KLUB_GUILD_ID"), user.ID, "Not subbed in Kajabi")
	// 	if err != nil {
	// 		fmt.Println("Error kicking user:", err)
	// 		continue
	// 	}
	// }

	fmt.Printf("Unsubscribed users: %v\n", len(unsubscribedUsers))
	fmt.Printf("Unsubscribed users: %v\n", strings.Join(userNames, "\n"))

	sortedResult := sort.StringSlice(userNamesDisplayNamesMap)
	j.s.ChannelMessageSendEmbed(os.Getenv("KLUB_LOG_CHANNEL_ID"), &discordgo.MessageEmbed{
		Title:       "Finished cleaning unsubscribed users",
		Description: fmt.Sprintf("[TEST] Number of kicked %d \n Unsubscribed users: %s \n", len(unsubscribedUsers), strings.Join(sortedResult, "\n")),
		Color:       0x00ff00,
		Footer: &discordgo.MessageEmbedFooter{
			Text: fmt.Sprintf("Time taken: %v", time.Since(start)),
		},
	})

	return unsubscribedUsers, nil
}

func containsRole(roles []string, targetRole string) bool {
	for _, role := range roles {
		if role == targetRole {
			return true
		}
	}
	return false
}

func contains(user string, users []string) bool {
	for _, u := range users {
		if u == user {
			return true
		}
	}
	return false
}
