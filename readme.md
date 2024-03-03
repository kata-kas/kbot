``` mermaid
graph TD
  subgraph clusterKajabi
    kajabiScraper((Kajabi Scraper))
  end

  subgraph clusterDiscordBot
    discordBot((Discord Bot))
  end

  subgraph clusterScheduler
    scheduler((Scheduler))
  end

  subgraph clusterServer
    modChannel((Mod Logs))
  end

  kajabiScraper -->|Scrape all users or search for a user| discordBot
  discordBot -->|Check if user is subscribed| modChannel
  discordBot -->|Kick user and send DM| modChannel
  discordBot -->|Log actions| modChannel
  scheduler -->|Daily scrape| kajabiScraper
  scheduler -->|Kick unsubscribed users| discordBot
```