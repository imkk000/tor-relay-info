package main

import (
	"errors"
	"os"

	"github.com/robfig/cron/v3"
	"github.com/rs/zerolog/log"
)

var (
	name       string
	webhookURL string
)

func main() {
	if len(os.Args) > 1 {
		name = os.Args[1]
	} else {
		name = os.Getenv("RELAY_NAME")
	}
	if name == "" {
		log.Fatal().Err(errors.New("relay name is empty (pass an argument or set RELAY_NAME)")).Msg("get config")
	}

	webhookURL = os.Getenv("DISCORD_WEBHOOK_URL")
	if webhookURL == "" {
		log.Fatal().Err(errors.New("DISCORD_WEBHOOK_URL is empty")).Msg("get config")
	}

	Run()

	c := cron.New()
	if _, err := c.AddFunc("@every 12h", Run); err != nil {
		log.Fatal().Err(err).Msg("schedule job")
	}
	c.Run()
}

func Run() {
	relays, err := FetchRelays(name)
	if err != nil {
		log.Err(err).Msg("fetch relays")
		return
	}

	embeds := make([]embed, 0, len(relays.Relays))
	for _, r := range relays.Relays {
		e, err := relayEmbed(r)
		if err != nil {
			log.Err(err).Msg("build embed")
			continue
		}
		embeds = append(embeds, e)
	}

	// Discord allows at most 10 embeds per webhook message.
	for i := 0; i < len(embeds); i += 10 {
		end := min(i+10, len(embeds))
		if err := sendWebhook(webhookURL, embeds[i:end]); err != nil {
			log.Err(err).Msg("send webhook")
		}
	}
}
