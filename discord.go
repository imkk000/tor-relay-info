package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	colorRunning = 0x57F287
	colorDown    = 0xED4245
)

type webhookPayload struct {
	Embeds []embed `json:"embeds"`
}

type embed struct {
	Title     string       `json:"title"`
	Color     int          `json:"color"`
	Fields    []embedField `json:"fields"`
	Timestamp string       `json:"timestamp,omitempty"`
}

type embedField struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline"`
}

func sendWebhook(url string, embeds []embed) error {
	if len(embeds) == 0 {
		return nil
	}
	body, err := json.Marshal(webhookPayload{Embeds: embeds})
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer func() { _ = resp.Body.Close() }()
	if resp.StatusCode >= 300 {
		return fmt.Errorf("discord webhook returned status %d", resp.StatusCode)
	}
	return nil
}
