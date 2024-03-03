package scheduler

import (
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/go-co-op/gocron"
	"github.com/klubbot/klubbot/internal/kajabi"
	"github.com/klubbot/klubbot/internal/normalization"
)

type GiveVipRoleJob struct {
	*gocron.Job
	s *discordgo.Session
}

func NewGiveVipRoleJob(s *discordgo.Session) (*GiveVipRoleJob, error) {
	return &GiveVipRoleJob{&gocron.Job{}, s}, nil
}

func (j GiveVipRoleJob) Run(url string) ([]string, []string, error) {
	kajabiUsers := kajabi.ScrapeUsers(url)

	discordUsers, err := j.s.GuildMembers(os.Getenv("KLUB_GUILD_ID"), "", 1000)
	if err != nil {
		fmt.Println("Error loading members:", err)
		return nil, nil, err
	}

	kajabiUserSet := make(map[string]bool)
	for _, kajabiUser := range kajabiUsers {
		kajabiUserSet[kajabiUser] = true
	}

	upgradedUsers := make([]string, 0)
	downgradedUsers := make([]string, 0)
	for _, discordUser := range discordUsers {
		if discordUser.User.Bot == true ||
			containsRole(discordUser.Roles, os.Getenv("KLUB_VIP_ROLE_ID")) ||
			containsRole(discordUser.Roles, os.Getenv("KLUB_ADMIN_ROLE_ID")) {
			continue
		}

		cleanedGlobalName := normalization.NormalizeForComparison(discordUser.User.GlobalName)
		cleanedUsername := normalization.NormalizeForComparison(discordUser.User.Username)

		if contains(cleanedGlobalName, kajabiUsers) ||
			contains(cleanedUsername, kajabiUsers) {
			upgradedUsers = append(upgradedUsers, discordUser.User.Username)
			err := j.s.GuildMemberRoleAdd(os.Getenv("KLUB_GUILD_ID"), discordUser.User.ID, os.Getenv("KLUB_VIP_ROLE_ID"))
			if err != nil {
				fmt.Println("Error adding VIP role:", err)
				continue
			}
		} else {
			downgradedUsers = append(downgradedUsers, discordUser.User.Username)
			err := j.s.GuildMemberRoleRemove(os.Getenv("KLUB_GUILD_ID"), discordUser.User.ID, os.Getenv("KLUB_VIP_ROLE_ID"))
			if err != nil {
				fmt.Println("Error removing VIP role:", err)
				continue
			}
		}
	}
	return upgradedUsers, downgradedUsers, nil
}
