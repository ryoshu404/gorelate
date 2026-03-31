package summarizer

import "github.com/ryoshu404/gorelate/internal/store"

// SummaryType defines the two summary cycles.
type SummaryType string

const (
	Emerging SummaryType = "emerging" // 8-hour shift summary
	Daily    SummaryType = "daily"    // 12:00 UTC daily briefing
)

// Summarizer handles LLM triage summary generation.
type Summarizer struct {
	apiKey string
}

func New(apiKey string) *Summarizer {
	return &Summarizer{apiKey: apiKey}
}

// Generate sends a filtered snapshot of HIGH confidence IOCs
// to the LLM and returns an analyst-facing summary paragraph.
// Returns an empty string and no error if there are no HIGH
// confidence IOCs in the window — caller should skip and log.
func (s *Summarizer) Generate(iocs []store.IOCRecord, t SummaryType) (string, error) {
	if len(iocs) == 0 {
		return "", nil
	}
	// TODO: implement Anthropic API call
	// Model: claude-sonnet-4-20250514
	// Input: filtered IOC snapshot (indicator, type, confidence, sources, tags, threat_category, vt_enriched)
	// Output: short analyst-facing paragraph
	return "", nil
}
