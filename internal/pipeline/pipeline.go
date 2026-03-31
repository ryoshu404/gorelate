package pipeline

import (
	"log/slog"

	"github.com/ryoshu404/gorelate/internal/fetchers"
)

// Pipeline runs one ingestion cycle:
// fetch → correlate → score → enrich → store → summarize → stubs
type Pipeline struct {
	fetchers []fetchers.Fetcher
}

func New(f []fetchers.Fetcher) *Pipeline {
	return &Pipeline{fetchers: f}
}

// Run executes a single ingestion cycle.
// Individual fetcher failures are logged and skipped —
// they never abort the cycle.
func (p *Pipeline) Run() {
	slog.Info("pipeline cycle starting")

	var all []fetchers.IOCRecord
	for _, f := range p.fetchers {
		records, err := f.Fetch()
		if err != nil {
			slog.Error("fetcher failed", "feed", f.Name(), "error", err)
			continue
		}
		slog.Info("fetcher complete", "feed", f.Name(), "count", len(records))
		all = append(all, records...)
	}

	// TODO: correlate
	// TODO: score
	// TODO: upsert to store
	// TODO: VT enrichment for HIGH confidence
	// TODO: check summary schedule
	// TODO: check stub threshold

	slog.Info("pipeline cycle complete", "total_records", len(all))
}
