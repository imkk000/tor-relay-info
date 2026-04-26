package main

type RelayList struct {
	Version         string  `json:"version"`
	BuildRevision   string  `json:"build_revision"`
	RelaysPublished string  `json:"relays_published"`
	Relays          []Relay `json:"relays"`
}

type Relay struct {
	Nickname                 string   `json:"nickname"`
	Fingerprint              string   `json:"fingerprint"`
	OrAddresses              []string `json:"or_addresses"`
	LastSeen                 string   `json:"last_seen"`
	LastChangedAddressOrPort string   `json:"last_changed_address_or_port"`
	FirstSeen                string   `json:"first_seen"`
	Running                  bool     `json:"running"`
	Flags                    []string `json:"flags"`
	Country                  string   `json:"country"`
	CountryName              string   `json:"country_name"`
	As                       string   `json:"as"`
	AsName                   string   `json:"as_name"`
	ConsensusWeight          int      `json:"consensus_weight"`
	LastRestarted            string   `json:"last_restarted"`
	BandwidthRate            int      `json:"bandwidth_rate"`
	BandwidthBurst           int      `json:"bandwidth_burst"`
	ObservedBandwidth        int      `json:"observed_bandwidth"`
	AdvertisedBandwidth      int      `json:"advertised_bandwidth"`
	ExitPolicy               []string `json:"exit_policy"`
	ExitPolicySummary        struct {
		Reject []string `json:"reject"`
	} `json:"exit_policy_summary"`
	Contact                 string   `json:"contact"`
	Platform                string   `json:"platform"`
	Version                 string   `json:"version"`
	VersionStatus           string   `json:"version_status"`
	AllegedFamily           []string `json:"alleged_family,omitempty"`
	EffectiveFamily         []string `json:"effective_family"`
	ConsensusWeightFraction float64  `json:"consensus_weight_fraction"`
	GuardProbability        float64  `json:"guard_probability"`
	MiddleProbability       float64  `json:"middle_probability"`
	ExitProbability         float64  `json:"exit_probability"`
	RecommendedVersion      bool     `json:"recommended_version"`
	Measured                bool     `json:"measured"`
}
