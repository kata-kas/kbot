package main

import (
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/klubbot/klubbot/bin/env"
	"github.com/klubbot/klubbot/internal/discord"
	"github.com/klubbot/klubbot/internal/scheduler"
)

func init() {
	env.Load(".env")
}

func main() {
	bot, err := discord.InitializeBot()
	if err != nil {
		log.Fatalf("[%s] Error initializing Discord bot: %v", time.Now().Format(time.RFC3339), err)
	}
	defer bot.Close()

	log.Printf("[%s] Bot is now running. Press Ctrl+C to exit.", time.Now().Format(time.RFC3339))

	sch, err := scheduler.NewScheduler()
	if err != nil {
		log.Fatalf("[%s] Error starting scheduler: %v", time.Now().Format(time.RFC3339), err)
	}

	unsubCleanerJob, err := scheduler.NewUnsubCleanerJob(bot)
	if err != nil {
		log.Fatalf("[%s] Error creating unsub cleaner job: %v", time.Now().Format(time.RFC3339), err)
	}

	giveVipRoleJob, err := scheduler.NewGiveVipRoleJob(bot)
	if err != nil {
		log.Fatalf("[%s] Error creating give VIP role job: %v", time.Now().Format(time.RFC3339), err)
	}

	sch.Every(1).Day().At("00:00").Do(unsubCleanerJob.Run)
	sch.Every(1).Day().At("00:00").Do(giveVipRoleJob.Run)

	sch.StartAsync()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Printf("[%s] Press Ctrl+C to exit", time.Now().Format(time.RFC3339))
	<-stop

	log.Printf("[%s] Gracefully shutting down.", time.Now().Format(time.RFC3339))
	sch.Stop()
	bot.Close()
}
