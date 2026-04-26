package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
)

var name string

func main() {
	if len(os.Args) <= 1 {
		log.Fatal().Err(errors.New("name is empty")).Msg("get arguments")
	}
	name = os.Args[1]
	Run()

	ticker := time.NewTicker(10 * time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		fmt.Print("\033[2J\033[H")
		Run()
	}
}

func Run() {
	relays, err := FetchRelays(name)
	if err != nil {
		log.Fatal().Err(err).Msg("fetch relays")
	}

	for _, r := range relays.Relays {
		lastSeen, err := time.Parse("2006-01-02 15:04:05", r.LastSeen)
		if err != nil {
			log.Err(err).Msgf("parse last seen time: %s", r.LastSeen)
			continue
		}
		lastRestarted, err := time.Parse("2006-01-02 15:04:05", r.LastRestarted)
		if err != nil {
			log.Err(err).Msgf("parse last restarted time: %s", r.LastRestarted)
			continue
		}
		uptime := time.Since(lastRestarted.UTC()).Truncate(time.Second)
		status := "running"
		if !r.Running {
			status = "down"
		}

		fmt.Printf("%-18s [%s]\n%s\n%s\n%s\n%s - %d\n\n",
			strings.Join(r.OrAddresses, ", "), status,
			formatDuration(uptime),
			lastSeen.Local().Format(time.RFC1123),
			strings.Join(r.Flags, ", "),
			formatBytes(r.ObservedBandwidth), r.ConsensusWeight,
		)
	}
}
