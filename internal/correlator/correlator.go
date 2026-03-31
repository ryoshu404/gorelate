package correlator

import "github.com/ryoshu404/gorelate/internal/fetchers"

// Result holds a deduplicated IOC with merged sources and tags
// from all feeds that reported it this cycle.
type Result struct {
	Indicator      string
	Type           string
	Sources        []string
	Tags           []string
	ThreatCategory string
	RawFeedData    map[string]any
}

// Correlate merges IOC records from all fetchers,
// deduplicating by indicator and merging sources/tags.
func Correlate(records []fetchers.IOCRecord) []Result {
	// TODO: implement dedup and source merge
	return nil
}
