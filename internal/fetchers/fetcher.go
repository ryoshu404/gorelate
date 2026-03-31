package fetchers

import "time"

// IOCRecord is the normalized output of every fetcher.
// All feed-specific data is mapped into this common shape
// before entering the pipeline.
type IOCRecord struct {
	Indicator       string            `json:"indicator"`
	Type            string            `json:"type"`            // ip, domain, url, hash
	Sources         []string          `json:"sources"`
	FirstSeen       time.Time         `json:"first_seen"`
	LastSeen        time.Time         `json:"last_seen"`
	Tags            []string          `json:"tags"`
	ThreatCategory  string            `json:"threat_category"`
	RawFeedData     map[string]any    `json:"raw_feed_data"`
}

// Fetcher is the interface every feed fetcher must implement.
type Fetcher interface {
	Name() string
	Fetch() ([]IOCRecord, error)
}
