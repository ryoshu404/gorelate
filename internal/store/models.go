package store

import "time"

// IOCRecord mirrors the ioc_records table.
type IOCRecord struct {
	ID             int64     `db:"id"`
	Indicator      string    `db:"indicator"`
	Type           string    `db:"type"`
	Sources        []string  `db:"sources"`        // stored as JSON array
	FirstSeen      time.Time `db:"first_seen"`
	LastSeen       time.Time `db:"last_seen"`
	Confidence     string    `db:"confidence"`     // low, medium, high
	Tags           []string  `db:"tags"`           // stored as JSON array
	ThreatCategory string    `db:"threat_category"`
	RawFeedData    []byte    `db:"raw_feed_data"`  // JSON blob
	VTEnriched     bool      `db:"vt_enriched"`
	DetectionStubs bool      `db:"detection_stubs"`
	Active         bool      `db:"active"`
}

// CycleSummary mirrors the cycle_summaries table.
type CycleSummary struct {
	ID          int64     `db:"id"`
	GeneratedAt time.Time `db:"generated_at"`
	SummaryType string    `db:"summary_type"` // emerging, daily
	Content     string    `db:"content"`
	WindowStart time.Time `db:"window_start"`
	WindowEnd   time.Time `db:"window_end"`
}

// DetectionStub mirrors the detection_stubs table.
type DetectionStub struct {
	ID          int64     `db:"id"`
	IOCRecordID int64     `db:"ioc_record_id"`
	StubType    string    `db:"stub_type"` // yara
	Content     string    `db:"content"`
	GeneratedAt time.Time `db:"generated_at"`
}
