package main

import (
	"fmt"
	"strings"
	"time"
)

func relayEmbed(r Relay) (embed, error) {
	lastSeen, err := time.Parse("2006-01-02 15:04:05", r.LastSeen)
	if err != nil {
		return embed{}, fmt.Errorf("parse last seen time %q: %w", r.LastSeen, err)
	}
	lastRestarted, err := time.Parse("2006-01-02 15:04:05", r.LastRestarted)
	if err != nil {
		return embed{}, fmt.Errorf("parse last restarted time %q: %w", r.LastRestarted, err)
	}

	uptime := time.Since(lastRestarted.UTC()).Truncate(time.Second)
	status, color := "running", colorRunning
	if !r.Running {
		status, color = "down", colorDown
	}

	title := r.Nickname
	if len(r.OrAddresses) > 0 {
		title = fmt.Sprintf("%s (%s)", r.Nickname, strings.Join(r.OrAddresses, ", "))
	}

	return embed{
		Title:     title,
		Color:     color,
		Timestamp: lastSeen.UTC().Format(time.RFC3339),
		Fields: []embedField{
			{Name: "Status", Value: status, Inline: true},
			{Name: "Uptime", Value: formatDuration(uptime), Inline: true},
			{Name: "Bandwidth", Value: formatBytes(r.ObservedBandwidth), Inline: true},
			{Name: "Consensus Weight", Value: fmt.Sprintf("%d", r.ConsensusWeight), Inline: true},
			{Name: "Last Seen", Value: lastSeen.Local().Format(time.RFC1123), Inline: false},
			{Name: "Flags", Value: strings.Join(r.Flags, ", "), Inline: false},
		},
	}, nil
}

func formatBytes(b int) string {
	switch {
	case b >= 1<<30:
		return fmt.Sprintf("%.2f GB", float64(b)/float64(1<<30))
	case b >= 1<<20:
		return fmt.Sprintf("%.2f MB", float64(b)/float64(1<<20))
	case b >= 1<<10:
		return fmt.Sprintf("%.2f KB", float64(b)/float64(1<<10))
	default:
		return fmt.Sprintf("%d B", b)
	}
}

func formatDuration(d time.Duration) string {
	h := int(d.Hours())
	m := int(d.Minutes()) % 60
	s := int(d.Seconds()) % 60
	return fmt.Sprintf("%d hours %d min %d sec", h, m, s)
}
