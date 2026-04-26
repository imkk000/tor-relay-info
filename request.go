package main

import (
	"encoding/json"
	"net/http"
	urlpkg "net/url"
	"time"
)

const baseURL = "https://onionoo.torproject.org"

var client = &http.Client{
	Transport: &http.Transport{
		ForceAttemptHTTP2: true,
	},
	Timeout: 1 * time.Minute,
}

func FetchRelays(name string) (*RelayList, error) {
	q := make(urlpkg.Values)
	q.Set("search", name)
	req, err := NewRequest(baseURL, "details", q)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var relays RelayList
	if err := json.NewDecoder(resp.Body).Decode(&relays); err != nil {
		return nil, err
	}
	return &relays, nil
}

func NewRequest(baseURL, path string, q urlpkg.Values) (*http.Request, error) {
	uri, err := urlpkg.JoinPath(baseURL, path)
	if err != nil {
		return nil, err
	}
	url, err := urlpkg.Parse(uri)
	if err != nil {
		return nil, err
	}
	url.RawQuery = q.Encode()
	req, err := http.NewRequest(http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Origin", "https://metrics.torproject.org")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:136.0) Gecko/20100101 Firefox/136.0")

	return req, nil
}
